# Introdução à Web e HTTP

## O que é a Web?

A World Wide Web (WWW) é um sistema de documentos interligados que são acessados pela internet. Desde sua criação no final dos anos 1980 por Tim Berners-Lee, a web evoluiu de um conjunto simples de páginas estáticas para um ecossistema dinâmico e interativo.

## Protocolo HTTP

O HTTP (Hypertext Transfer Protocol) é o protocolo que permite a comunicação entre clientes (como navegadores) e servidores web. É um protocolo:

- **Stateless**: Cada requisição é independente
- **Cliente-Servidor**: Separação clara de responsabilidades
- **Baseado em texto**: Mensagens são legíveis por humanos
- **Extensível**: Através de headers, métodos e status codes

### Versões do HTTP

1. **HTTP/0.9** (1991)
   - Apenas método GET
   - Sem headers
   - Apenas documentos HTML

2. **HTTP/1.0** (1996)
   - Headers
   - Mais métodos (POST, HEAD)
   - Metadados

3. **HTTP/1.1** (1997)
   - Conexões persistentes (Keep-Alive)
   - Host header obrigatório
   - Chunked transfers

4. **HTTP/2** (2015)
   - Multiplexação
   - Server Push
   - Compressão de headers
   - Binário em vez de texto

5. **HTTP/3** (2022)
   - Baseado em QUIC (UDP)
   - Melhor performance
   - Menos latência

## Anatomia de uma Requisição HTTP

```http
GET /api/users HTTP/1.1
Host: api.exemplo.com
Accept: application/json
Authorization: Bearer token123
```

### Componentes principais:
1. **Método** (GET, POST, PUT, DELETE, etc.)
2. **Path** (caminho do recurso)
3. **Versão do protocolo**
4. **Headers** (metadados)
5. **Body** (opcional, dados da requisição)

## Anatomia de uma Resposta HTTP

```http
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 82

{
    "id": 1,
    "name": "João",
    "email": "joao@exemplo.com"
}
```

### Componentes principais:
1. **Status Line** (protocolo, código, mensagem)
2. **Headers** (metadados)
3. **Body** (opcional, dados da resposta)

## Status Codes

Os códigos de status HTTP são agrupados em 5 classes:

- **1xx**: Informacional
- **2xx**: Sucesso
- **3xx**: Redirecionamento
- **4xx**: Erro do Cliente
- **5xx**: Erro do Servidor

## Métodos HTTP

| Método  | Descrição                               | Idempotente? | Seguro? |
|---------|----------------------------------------|--------------|---------|
| GET     | Recupera um recurso                     | Sim         | Sim     |
| POST    | Cria um novo recurso                    | Não         | Não     |
| PUT     | Atualiza um recurso existente          | Sim         | Não     |
| DELETE  | Remove um recurso                       | Sim         | Não     |
| PATCH   | Atualiza parcialmente um recurso        | Não         | Não     |
| HEAD    | Como GET, mas sem body na resposta      | Sim         | Sim     |
| OPTIONS | Retorna métodos suportados pelo recurso | Sim         | Sim     |

## Headers Comuns

### Request Headers
- `Accept`: Tipos de conteúdo aceitos
- `Authorization`: Credenciais de autenticação
- `User-Agent`: Identificação do cliente
- `Content-Type`: Tipo do conteúdo enviado

### Response Headers
- `Content-Type`: Tipo do conteúdo retornado
- `Content-Length`: Tamanho do body
- `Cache-Control`: Diretivas de cache
- `Set-Cookie`: Define cookies no cliente

## Próximos Passos

- [História dos Servidores Web](02-historia.md)
- [O Pacote net/http](03-net-http.md) 