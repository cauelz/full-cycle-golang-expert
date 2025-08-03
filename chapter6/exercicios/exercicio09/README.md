# Exercício 09: Migrations Simples

## Objetivo
Aprender a criar e rodar migrations simples via código Go.

## Tarefas
1. Escreva um programa Go que:
   - Crie uma tabela chamada `produtos` com campos `id` e `nome` se ela não existir.
   - Exiba uma mensagem de sucesso ou erro.

## Dicas
- Use `CREATE TABLE IF NOT EXISTS`.
- Execute o comando via `db.Exec`.

---

**Exemplo de saída esperada:**

```
Tabela criada ou já existente.
``` 