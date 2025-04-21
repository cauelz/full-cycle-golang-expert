package main

import (
	"encoding/json"
	"io"
	"net/http"
)

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

// Este primeiro exemplo mostra como fazemos um servirdor HTTP básico
// realizando uma requisição para o buscaCEP (https://viacep.com.br)
func RunServer() {

	http.HandleFunc("/", BuscaCepHandler)

	http.ListenAndServe(":8080", nil)

}

// w representa o writer, ou seja, a resposta que será enviada ao cliente
// r representa a requisição que o cliente fez
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Página não encontrada"))
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP não informado"))
		return
	}

	cep, error := BuscaCep(cepParam)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar o CEP"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Opção 1: Usando unmarshal
	// O método Marshal converte a struct em JSON
	// O método Marshal retorna um slice de bytes e um erro
	// json, error := json.Marshal(cep)
	// if error != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Erro ao converter o CEP em JSON"))
	// 	return
	// }
	// w.Write(json)

	// Opção 2: 
	// O método Encode do pacote json converte a struct em JSON e escreve no ResponseWriter
	// O método Encode retorna um erro caso ocorra algum problema
	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {

	// Fazendo uma requisição GET para a API do ViaCEP
	// O endpoint da API é https://viacep.com.br/ws/{cep}/json/
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

	if error != nil {
		return nil, error
	}

	// Executamos o defer para garantir que o corpo da resposta será fechado
	// isto é importante para evitar vazamentos de memória
	defer resp.Body.Close()

	// Utilizamos o pacote io para ler o corpo da resposta
	// O método ReadAll lê todo o corpo da resposta e retorna um slice de bytes
	// O método ReadAll retorna um erro caso ocorra algum problema
	body, error := io.ReadAll(resp.Body)

	if error != nil {
		return nil, error
	}

	var c ViaCEP
	// Fazemos o Unmarshal do corpo da resposta para a struct ViaCEP
	// O método Unmarshal converte o JSON em uma struct

	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}

	return &c, nil
}