# Exemplo CRUD em Go com MySQL

## Visão Geral

Este exemplo demonstra como realizar operações básicas de CRUD (Create, Read, Update, Delete) em um banco de dados MySQL usando a linguagem Go. O código manipula uma entidade chamada `Product` (Produto), mostrando como inserir, atualizar, buscar e deletar registros em uma tabela do banco de dados.

---

## Estrutura do Projeto

O código está todo no arquivo `main.go` e utiliza as seguintes bibliotecas externas:

- **github.com/go-sql-driver/mysql**: Driver para conectar o Go ao MySQL.
- **github.com/google/uuid**: Gera identificadores únicos universais (UUID) para os produtos.

---

## Explicação Detalhada

### 1. Definição da Estrutura Product

```go
type Product struct {
    ID    string
    Name  string
    Price float64
}
```

- **ID**: Identificador único do produto (string, gerado como UUID).
- **Name**: Nome do produto.
- **Price**: Preço do produto (float64).

### 2. Função de Criação de Produto

```go
func NewProduct(name string, price float64) *Product {
    return &Product{
        ID:    uuid.New().String(),
        Name:  name,
        Price: price,
    }
}
```

- Cria um novo produto com um UUID único.
- Retorna um ponteiro para a estrutura `Product`.

### 3. Função Principal (`main`)

#### a) Conexão com o Banco de Dados

```go
db, error := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
```
- Abre uma conexão com o banco MySQL usando usuário `root`, senha `root`, host `localhost`, porta `3306` e banco `goexpert`.
- O objeto `db` é do tipo `*sql.DB`, que gerencia a conexão.

#### b) Tratamento de Erros

```go
if error != nil {
    panic(error)
}
```
- Se ocorrer erro na conexão, o programa é interrompido com `panic`.

#### c) Fechamento da Conexão

```go
defer db.Close()
```
- Garante que a conexão será fechada ao final da execução da função `main`.

#### d) Operações CRUD

- **Inserção**: Cria um produto e insere no banco.
- **Atualização**: Altera o nome do produto e atualiza no banco.
- **Seleção**: (Comentado) Busca um produto pelo ID.
- **Remoção**: Deleta o produto pelo ID.
- **Listagem**: Busca todos os produtos e imprime no console.

---

## Funções CRUD

### 1. Inserir Produto

```go
func insertProduct(db *sql.DB, product *Product) error {
    stmt, error := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
    // ... tratamento de erro ...
    _, error = stmt.Exec(product.ID, product.Name, product.Price)
    // ... tratamento de erro ...
    return nil
}
```
- Prepara um comando SQL para inserir um produto.
- Usa placeholders (`?`) para evitar SQL Injection.
- Executa o comando passando os valores do produto.

### 2. Atualizar Produto

```go
func updateProduct(db *sql.DB, product *Product) error {
    stmt, error := db.Prepare("update products set name = ?, price = ? where id = ?")
    // ... tratamento de erro ...
    _, error = stmt.Exec(product.Name, product.Price, product.ID)
    // ... tratamento de erro ...
    return nil
}
```
- Atualiza o nome e preço do produto com base no ID.

### 3. Selecionar Produto por ID

```go
func selectProduct(db *sql.DB, id string) (*Product, error) {
    stmt, error := db.Prepare("select id, name, price from products where id = ?")
    // ... tratamento de erro ...
    var product Product
    error = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
    // ... tratamento de erro ...
    return &product, nil
}
```
- Busca um produto pelo ID.
- Usa `QueryRow` para retornar um único resultado.
- Preenche a estrutura `Product` com os dados do banco.

### 4. Selecionar Todos os Produtos

```go
func selectAllProducts(db *sql.DB) ([]*Product, error) {
    rows, error := db.Query("select id, name, price from products")
    // ... tratamento de erro ...
    var products []*Product
    for rows.Next() {
        var product Product
        error = rows.Scan(&product.ID, &product.Name, &product.Price)
        // ... tratamento de erro ...
        products = append(products, &product)
    }
    return products, nil
}
```
- Executa uma consulta para buscar todos os produtos.
- Itera sobre os resultados, preenchendo uma lista de ponteiros para `Product`.

### 5. Deletar Produto

```go
func deleteProduct(db *sql.DB, id string) error {
    stmt, error := db.Prepare("delete from products where id = ?")
    // ... tratamento de erro ...
    _, error = stmt.Exec(id)
    // ... tratamento de erro ...
    return nil
}
```
- Deleta um produto com base no ID.

---

## Conceitos Importantes

- **Ponteiros**: Usados para evitar cópias desnecessárias de estruturas e permitir alteração dos dados originais.
- **Tratamento de Erros**: Sempre verifique e trate erros após operações de banco de dados.
- **Defer**: Garante o fechamento de recursos (como conexões e statements) ao final da função.
- **SQL Injection**: O uso de `?` nos comandos SQL previne ataques de injeção de SQL.
- **UUID**: Garante que cada produto tenha um identificador único, evitando colisões.

---

## Boas Práticas

- Sempre feche conexões e statements com `defer`.
- Nunca ignore erros retornados por funções.
- Use variáveis de ambiente para credenciais sensíveis (não hardcode em produção).
- Separe a lógica de acesso a dados em pacotes diferentes em projetos maiores.

---

## Resumo

Este exemplo cobre o ciclo completo de manipulação de dados em um banco MySQL usando Go, com foco em clareza, segurança e boas práticas. Ideal para quem está começando a trabalhar com banco de dados em Go. 