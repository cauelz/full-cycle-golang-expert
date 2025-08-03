# Exercício 01: Conexão Simples com Banco de Dados

## Objetivo
Aprender a conectar um programa Go a um banco de dados relacional (MySQL ou PostgreSQL) usando a biblioteca `database/sql`.

## Tarefas
1. Instale o driver do banco de dados de sua escolha (ex: `github.com/go-sql-driver/mysql` ou `github.com/lib/pq`).
2. Crie um programa Go que:
   - Estabeleça uma conexão com o banco de dados.
   - Teste a conexão usando `db.Ping()`.
   - Exiba uma mensagem de sucesso ou erro.

## Dicas
- Use variáveis de ambiente para armazenar as credenciais de conexão.
- Consulte a documentação do driver escolhido para montar a string de conexão.

---

**Exemplo de saída esperada:**

```
Conexão bem-sucedida!
```
Ou, em caso de erro:
```
Erro ao conectar: <mensagem de erro>
``` 