package json

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
}

func example1() {
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
