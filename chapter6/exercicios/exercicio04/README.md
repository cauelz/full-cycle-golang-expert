# Exercício 04: Consulta com Parâmetros

## Objetivo
Aprender a buscar registros usando parâmetros (WHERE) em Go.

## Tarefas
1. Escreva um programa Go que:
   - Solicite ao usuário um nome para buscar.
   - Busque na tabela `usuarios` apenas os registros que correspondam ao nome informado.
   - Exiba os resultados.

## Dicas
- Use `db.Query` ou `db.QueryRow` com parâmetros.
- Use `?` para parâmetros no MySQL e `$1` no PostgreSQL.

---

**Exemplo de saída esperada:**

```
Digite o nome para buscar: João
ID: 1, Nome: João
``` 