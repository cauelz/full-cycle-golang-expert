# Pacotes em Go  

O conceito de pacotes (`packages`) em Go é bem direto e simples: trata-se de uma coleção de arquivos `.go` que estão no mesmo diretório. Uma boa prática é agrupar arquivos `.go` em um pacote que compartilhe a mesma finalidade ou contexto. Por exemplo: um pacote chamado "operações matemáticas" pode conter funções como soma, subtração, divisão etc.  

Todo arquivo que pertence a um pacote começa declarando o nome desse pacote com `package <nome>`.  

---

## Pacote `main`  

O pacote `main` é especial, pois indica o **ponto de entrada** de um programa em Go. Quando compilamos e executamos um projeto, o Go procura pelo pacote `main` para iniciar a execução.  

Quando o pacote `main` é encontrado, é obrigatório ter uma função que marque o início da execução do programa: a `func main()`. Essa função não recebe argumentos e não retorna valores.  

### Exemplo de estrutura:  
```go
package main  

func main() {  
    // Código de inicialização aqui  
}