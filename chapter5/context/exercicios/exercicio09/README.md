# Exercício 9 — Contagem de palavras com timeout

## Conceitos
- Uso de context com timeout para limitar operações
- Processamento paralelo de arquivos grandes
- Coleta de resultados parciais via channels

## Enunciado
Crie um programa que conte o número de palavras em um arquivo grande, processando o arquivo em paralelo (divida o arquivo em partes e processe cada parte em uma goroutine). O programa deve aceitar um timeout via context: se o tempo acabar, o programa deve cancelar todas as goroutines e retornar o resultado parcial.

> _Objetivos: Praticar uso de context com timeout, divisão de tarefas entre goroutines e coleta de resultados via channel._ 