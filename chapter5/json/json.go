package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Pessoa struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
}

	
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func Example1() {
    // 1. Criar uma instância de Pessoa
    p := Pessoa{
        Nome:  "Maria",
        Idade: 30,
    }

    // 2. Converter (Marshal) para JSON
    pJSON, err := json.Marshal(p)
    if err != nil {
        fmt.Println("Erro ao converter para JSON:", err)
        return
    }

    fmt.Println("JSON gerado:", string(pJSON))

    // 3. Converter (Unmarshal) de JSON para struct
    var p2 Pessoa
    err = json.Unmarshal(pJSON, &p2)
    if err != nil {
        fmt.Println("Erro ao ler JSON:", err)
        return
    }

    fmt.Printf("Struct lida: %+v\n", p2)
}

func Example2() {
    // Exemplo 2 - Ler um CEP pelo terminal e realizar uma requisição HTTP para a API
    // https://viacep.com.br

    // 1. Ler o CEP do terminal com o os.Args.
    // Explicação: os.Args é um slice de strings que contém os argumentos passados para o programa.
    // O primeiro elemento (os.Args[0]) é o nome do programa, e os demais são os argumentos passados.
    for _, cep := range os.Args[2:] {

        // 2. Fazer uma requisição HTTP GET para a URL informada.
        req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

        if err != nil {
            // Explicação Fprintf : escreve a string formatada para o Writer (neste caso, os.Stderr).
            fmt.Fprintf(os.Stderr, "Erro ao fazer requisição:", err)
            // Explicação: os.Exit(1) encerra o programa com o código de saída 1, indicando erro.
            os.Exit(1)
        }

        // 3. Criar um defer para fechar a resposta.
        defer req.Body.Close()

        // 4. Ler o corpo da resposta.
        // Explicação: io.ReadAll lê todos os dados do corpo da resposta e retorna um slice de bytes.
        res, err := io.ReadAll(req.Body)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Erro ao ler resposta:", err)
            os.Exit(1)
        }

        // 5. Criar uma variável do tipo ViaCEP.
        var viaCep ViaCEP
        // 6. Fazer o Unmarshal do JSON lido para a variável viaCep.
        err = json.Unmarshal(res, &viaCep)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Erro ao fazer Unmarshal:", err)
            os.Exit(1)
        }

        // 7. Imprimir os dados do CEP.
        fmt.Println("Resposta ViaCEP:", viaCep)

        // 8. Escrever um arquivo com os dados do CEP.

        file, err := os.Create("dados_cep.txt")

        if err != nil {
            fmt.Fprintf(os.Stderr, "Erro ao criar arquivo:", err)
            os.Exit(1)
        }

        defer file.Close()

        _, err = file.WriteString(fmt.Sprintf(
            "CEP: %s\nLocalidade: %s\nBairro: %s\nLogradouro: %s\nComplemento: %s\nEstado: %s\nUnidade: %s\nRegiao: %s\nIbge: %s\nGia: %s\nDdd: %s\nSiafi: %s\n",
            viaCep.Cep, viaCep.Localidade, viaCep.Bairro, viaCep.Logradouro, viaCep.Complemento, viaCep.Estado, viaCep.Unidade, viaCep.Regiao, viaCep.Ibge, viaCep.Gia, viaCep.Ddd, viaCep.Siafi))

        if err != nil {
            fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo:", err)
            os.Exit(1)
        }
    }
}