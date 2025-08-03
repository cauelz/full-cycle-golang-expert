# Exercício 03: Consulta Simples

## Objetivo
Aprender a consultar registros de uma tabela usando Go.

## Tarefas
1. Certifique-se de que a tabela `usuarios` possui alguns registros.
2. Escreva um programa Go que:
   - Conecte ao banco de dados.
   - Busque todos os registros da tabela `usuarios`.
   - Exiba os resultados no terminal.

## Dicas
- Use `db.Query` para executar SELECTs.
- Use um loop para iterar sobre os resultados.

---

**Exemplo de saída esperada:**

```
ID: 1, Nome: João
ID: 2, Nome: Maria
...
``` 