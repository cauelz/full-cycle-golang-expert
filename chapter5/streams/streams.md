# Streams no Processamento de Dados

## O que são Streams?
**Streams** são um padrão de processamento de dados que permite manipular informações de forma **contínua e incremental**, sem carregar todo o conjunto de dados na memória de uma só vez. Elas funcionam como um fluxo de dados dividido em partes menores (*chunks*), ideal para grandes volumes de dados ou operações em tempo real.

---

## Origem do Conceito
O conceito surgiu de sistemas operacionais e programação de baixo nível, inspirado no **paradigma de pipes do Unix** (ex: `cat arquivo.txt | grep "palavra"`), onde a saída de um comando é redirecionada para a entrada de outro. Em linguagens de programação, o conceito foi adaptado para lidar com operações de I/O de forma eficiente, especialmente em cenários com dados maiores que a memória disponível.

---

## Tipos de Streams
1. **Readable Streams**: Leitura de dados (ex: ler um arquivo).
2. **Writable Streams**: Escrita de dados (ex: salvar um arquivo).
3. **Duplex Streams**: Leitura e escrita simultânea (ex: sockets de rede).
4. **Transform Streams**: Transformam dados durante o fluxo (ex: compressão).

---

## Para Que Servem?
### 1. Eficiência de Memória
- Evitam carregar dados massivos na RAM.
- Exemplo: Processar um arquivo de 10 GB sem estourar a memória.

### 2. Processamento em Tempo Real
- Dados são processados assim que disponíveis.
- Exemplo: Transcodificar vídeo durante o download.

### 3. Composição de Operações
- Encadeamento de streams via `pipe` (Node.js) ou `io.Copy` (Go).
- Exemplo: Ler → Comprimir → Criptografar → Enviar.

---

## Funcionamento em Node.js e Go

### Node.js
```javascript
const fs = require('fs');
const readStream = fs.createReadStream('arquivo.txt');

readStream
  .on('data', (chunk) => console.log(chunk))
  .on('end', () => console.log('Fim do fluxo'));
```
- **APIs**: `fs.createReadStream`, `fs.createWriteStream`.
- **Eventos**: `data`, `end`, `error`.
- **Métodos**: `.pipe()` para encadear streams.

### Go
```go
package main

import (
    "io"
    "os"
)

func main() {
    src, _ := os.Open("input.txt")
    defer src.Close()

    dst, _ := os.Create("output.txt")
    defer dst.Close()

    io.Copy(dst, src) // Copia dados em chunks
}
```
- **Interfaces**: `io.Reader` (`Read()`) e `io.Writer` (`Write()`).
- **Funções**: `io.Copy`, `bufio.NewReader`.

---

## Vantagens
- ✅ **Escalabilidade**: Lida com milhões de conexões (ex: servidores web).
- ✅ **Performance**: Reduz latência (processa dados incrementalmente).
- ✅ **Flexibilidade**: Combina com transformações (ex: filtros, compressão).
- ✅ **Resiliência**: Erros são tratados por chunk.

---

## Casos de Uso
1. **Streaming de Vídeo**: Servir partes de um vídeo conforme o usuário assiste.
2. **ETL (Extract, Transform, Load)**: Processar datasets de 100 GB+.
3. **IoT**: Receber dados de sensores em tempo real.
4. **Logs**: Analisar logs de servidores enquanto são gerados.

---

## Desafios
- **Backpressure**: Controle de fluxo quando o produtor é mais rápido que o consumidor.
- **Complexidade**: Gerenciar erros, concorrência e finalização de streams.

---

## Conclusão
Streams são fundamentais para sistemas que exigem **eficiência de memória** e **processamento contínuo**. Em Node.js, são baseadas em eventos; em Go, em interfaces como `io.Reader` e `io.Writer`. Dominar esse conceito é essencial para construir aplicações robustas com grandes volumes de dados.