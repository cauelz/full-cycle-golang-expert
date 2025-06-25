# Exercício 11 — Pipeline de processamento com context

## Conceitos
- Construção de pipelines com channels
- Cancelamento de pipelines usando context
- Manipulação eficiente de arquivos grandes

## Enunciado
Monte um pipeline de processamento de linhas de um arquivo:
1. Uma goroutine lê as linhas do arquivo e envia para um channel.
2. Uma ou mais goroutines processam as linhas (ex: transformam o texto, filtram, etc) e enviam para outro channel.
3. Uma goroutine final coleta e salva o resultado.
Implemente o cancelamento do pipeline usando context.

> _Objetivos: Praticar construção de pipelines com channels, uso de context para cancelar todo o pipeline e manipulação eficiente de arquivos grandes._ 