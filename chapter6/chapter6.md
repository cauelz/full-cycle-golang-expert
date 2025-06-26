# Trabalhando com Banco de Dados em Go

A linguagem Go oferece um pacote padrão chamado `database/sql` para trabalhar com bancos de dados relacionais de forma simples e eficiente. Esse pacote fornece uma interface genérica para interação com diversos bancos de dados, como MySQL, PostgreSQL, SQLite, entre outros, por meio de drivers específicos.

## O que é o pacote `database/sql`?

O pacote `database/sql` não implementa a comunicação direta com o banco de dados, mas define uma interface comum. Para se conectar a um banco específico, é necessário instalar o driver correspondente. Por exemplo, para usar o MySQL, você pode utilizar o driver `github.com/go-sql-driver/mysql`.

## Como funciona?

O fluxo básico para usar o `database/sql` é:
1. Importar o pacote `database/sql` e o driver do banco desejado.
2. Abrir uma conexão com o banco de dados usando `sql.Open`.
3. Executar comandos SQL (consultas, inserções, atualizações, etc.).
4. Fechar a conexão ao final do uso.

## Exemplo básico de uso

Abaixo está um exemplo simples de como conectar-se a um banco de dados MySQL, inserir e consultar dados:

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Abrindo conexão (user:senha@/nome_do_banco)
    db, err := sql.Open("mysql", "root:senha@/exemplo")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Inserindo um dado
    _, err = db.Exec("INSERT INTO usuarios(nome) VALUES(?)", "João")
    if err != nil {
        panic(err)
    }

    // Consultando dados
    rows, err := db.Query("SELECT id, nome FROM usuarios")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var nome string
        err = rows.Scan(&id, &nome)
        if err != nil {
            panic(err)
        }
        fmt.Println(id, nome)
    }
}
```

> **Dica:** Sempre confira a documentação do driver escolhido para detalhes de configuração e uso.

Com o `database/sql`, é possível realizar operações como transações, consultas preparadas, tratamento de erros e muito mais, tornando o acesso a bancos de dados em Go robusto e flexível.

