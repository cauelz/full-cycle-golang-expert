package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

// Product representa um produto da API
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Client é um cliente HTTP personalizado com recursos avançados
type Client struct {
	baseURL    string
	httpClient *http.Client
	bufferPool *sync.Pool
}

// NewClient cria um novo cliente HTTP otimizado
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				MaxConnsPerHost:     100,
				IdleConnTimeout:     90 * time.Second,
				TLSHandshakeTimeout: 10 * time.Second,
			},
		},
		bufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

// GetProducts busca todos os produtos da API
func (c *Client) GetProducts(ctx context.Context) ([]Product, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/products", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return products, nil
}

// CreateProduct cria um novo produto via API
func (c *Client) CreateProduct(ctx context.Context, product Product) error {
	// Obter buffer do pool
	buf := c.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer c.bufferPool.Put(buf)

	// Codificar produto para JSON usando o buffer do pool
	if err := json.NewEncoder(buf).Encode(product); err != nil {
		return fmt.Errorf("error encoding product: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/products", buf)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// UploadFile faz upload de um arquivo para a API
func (c *Client) UploadFile(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Criar form-data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}
	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/upload", body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// SearchProducts busca produtos com parâmetros de query
func (c *Client) SearchProducts(ctx context.Context, query url.Values) ([]Product, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/products", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return products, nil
}

func main() {
	// Criar cliente
	client := NewClient("http://localhost:8080")

	// Criar contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Buscar produtos
	products, err := client.GetProducts(ctx)
	if err != nil {
		log.Fatalf("Error getting products: %v", err)
	}
	fmt.Printf("Found %d products\n", len(products))

	// Criar novo produto
	newProduct := Product{
		Name:        "New Product",
		Description: "A new product",
		Price:       99.99,
	}
	if err := client.CreateProduct(ctx, newProduct); err != nil {
		log.Fatalf("Error creating product: %v", err)
	}
	fmt.Println("Product created successfully")

	// Buscar produtos com filtros
	query := url.Values{}
	query.Set("minPrice", "50")
	query.Set("maxPrice", "100")
	filteredProducts, err := client.SearchProducts(ctx, query)
	if err != nil {
		log.Fatalf("Error searching products: %v", err)
	}
	fmt.Printf("Found %d filtered products\n", len(filteredProducts))

	// Upload de arquivo
	if err := client.UploadFile(ctx, "example.txt"); err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}
	fmt.Println("File uploaded successfully")
} 