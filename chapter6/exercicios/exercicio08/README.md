# Exercício 08: Transações

## Objetivo
Aprender a realizar múltiplas operações dentro de uma transação em Go.

## Tarefas
1. Escreva um programa Go que:
   - Inicie uma transação.
   - Insira dois usuários diferentes dentro da mesma transação.
   - Faça commit ou rollback conforme o sucesso das operações.

## Dicas
- Use `db.Begin` para iniciar a transação.
- Use `tx.Commit` e `tx.Rollback`.

---

**Exemplo de saída esperada:**

```
Transação realizada com sucesso!
``` 