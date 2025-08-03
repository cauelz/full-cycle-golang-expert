# Exercício 10: Uso de Contexto

## Objetivo
Aprender a usar `context.Context` para controlar timeout/cancelamento de queries no Go.

## Tarefas
1. Escreva um programa Go que:
   - Use `context.WithTimeout` para definir um tempo limite para uma consulta.
   - Execute uma consulta usando o contexto.
   - Exiba o resultado ou mensagem de timeout.

## Dicas
- Use `db.QueryContext` ou `db.ExecContext`.
- Importe o pacote `context`.

---

**Exemplo de saída esperada:**

```
Consulta realizada com sucesso!
```
Ou, em caso de timeout:
```
Timeout atingido!
``` 