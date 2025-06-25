# Exercício 10 — Busca de padrão com contexto de cancelamento

## Conceitos
- Busca concorrente em arquivos grandes
- Cancelamento imediato de goroutines com context
- Sincronização de resultados via channels

## Enunciado
Implemente um programa que busque por uma palavra/padrão em um arquivo grande, usando múltiplas goroutines para processar diferentes partes do arquivo. O programa deve parar imediatamente todas as buscas assim que encontrar a primeira ocorrência, usando context para cancelar as goroutines restantes.

> _Objetivos: Praticar busca concorrente em arquivos, uso de context para cancelamento imediato e sincronização de goroutines com channels._ 