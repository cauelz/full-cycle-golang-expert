# Exercício 07: Prepared Statements

## Objetivo
Aprender a usar prepared statements para inserção e consulta em Go.

## Tarefas
1. Escreva um programa Go que:
   - Use prepared statements para inserir um novo usuário.
   - Use prepared statements para buscar usuários pelo nome.
   - Exiba os resultados.

## Dicas
- Use `db.Prepare` para criar o statement.
- Use `stmt.Exec` e `stmt.Query` para executar.

---

**Exemplo de saída esperada:**

```
Usuário inserido com sucesso!
ID: 3, Nome: João
``` 