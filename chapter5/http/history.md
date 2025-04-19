# Padrão de Servidores HTTP: Origem e Evolução

## 1. Origem: HTTP e os Primeiros Servidores

### 1.1 Proposta inicial e HTTP/0.9  
- **Contexto (1989–1991):** Tim Berners‑Lee, trabalhando no CERN, propôs um “sistema de hipertexto” para compartilhar documentos científicos. Seu documento de 1989, “HyperText and CERN”, definiu conceitos-chave como URLs, HTML e um protocolo simples de requisição-resposta sobre TCP, batizado inicialmente de HTTP 0.9.  
- **Características do HTTP 0.9:**  
  - Apenas o método `GET`.  
  - Respostas consistiam unicamente no corpo do recurso solicitado (HTML cru), sem headers.  
  - Uso direto de TCP: o cliente abria uma conexão, enviava `GET /arquivo.html\n`, recebia o conteúdo e fechava a conexão.  
- **Por que HTTP 0.9?**  
  - Simplicidade para rápido desenvolvimento do protótipo no NeXT do Berners‑Lee.  
  - Atendeu aos casos de uso iniciais de documentos estáticos e hiperlinks entre eles. :contentReference[oaicite:0]{index=0}

### 1.2 CERN httpd: o primeiro servidor Web  
- **Desenvolvimento:** Entre 1990 e 1991, Berners‑Lee, Ari Luotonen e Henrik Frystyk Nielsen implementaram o primeiro daemon HTTP em C, chamado **CERN httpd** (posteriormente “W3C httpd”) :contentReference[oaicite:1]{index=1}.  
- **Funcionamento básico:**  
  1. **Aceitação de conexões TCP** em uma porta (por padrão, 80).  
  2. **Leitura mínima** da linha de requisição (ex.: `GET /index.html HTTP/0.9`).  
  3. **Mapeamento** do caminho solicitado para um arquivo no disco (`/index.html`).  
  4. **Envio do conteúdo** do arquivo e fechamento da conexão.  
- **Primeira página publicada:** Em dezembro de 1990, a URL `http://info.cern.ch/` já servia uma página explicando o que era a Web e como usá-la :contentReference[oaicite:2]{index=2}.  

### 1.3 Transição para HTTP/1.0  
- **Limitações do 0.9:** sem headers, sem status codes, sem suporte a múltiplos métodos ou upload de dados.  
- **Rascunhos HTTP/1.0 (1992–1996):** apareceram esboços que introduziram:  
  - **Headers de requisição/resposta** (Content-Type, Content-Length, etc.).  
  - **Status codes** (200, 404, 500…).  
  - **Suporte a diferentes métodos** (`POST`, `HEAD`).  
- **Publicação oficial:** Em maio de 1996, HTTP/1.0 foi formalizado como RFC 1945, consolidando práticas de implementações anteriores :contentReference[oaicite:3]{index=3}.  

### 1.4 CGI e a geração de conteúdo dinâmico  
- **O que era CGI (1993):** “Common Gateway Interface” padronizou a interação entre servidores HTTP e programas externos (Perl, C, etc.) para gerar respostas dinâmicas :contentReference[oaicite:4]{index=4}.  
- **Fluxo CGI:**  
  1. Servidor detecta um caminho sob `/cgi-bin/`.  
  2. Invoca o programa correspondente em um novo processo.  
  3. Passa variáveis de ambiente (QUERY_STRING, REQUEST_METHOD, etc.).  
  4. Lê a saída do programa (com headers próprios) e repassa ao cliente.  
- **Desvantagens:**  
  - **Alto custo** de criar um processo por requisição.  
  - **Escalabilidade limitada** em cenários de alto volume ou alta concorrência.  

---

## 2. Evolução do Modelo “Thread‑per‑Request” para Thread Pools em Servidores Java

### 2.1 Modelo “Thread‑per‑Request” Original
- **Como funcionava:** a cada requisição HTTP recebida, o contêiner (por exemplo, Apache Tomcat) **criava** uma nova `Thread` Java, chamava o método `service()`/`doGet()`/`doPost()`, e, ao final do processamento, **encerrava** essa `Thread`.  
- **Desvantagens principais:**  
  - **Overhead de Criação/Destruição** – criar e destruir objetos `Thread` repetidamente consome CPU, memória e gera pressão no garbage collector.  
  - **Fragmentação de Recursos** – picos de tráfego podem levar a milhares de threads simultâneas, sobrecarregando o sistema operacional. citeturn0search1  

### 2.2 Introdução dos Thread Pools
- **Motivação:** reduzir o custo de continuamente instanciar e destruir threads, além de controlar melhor o número máximo de threads ativas.  
- **Implementação em contêineres Java:**  
  - Contêineres como **Tomcat** e **Jetty** passaram a **pré‑criar** um conjunto de threads no startup do servidor, mantendo‑as **ociosas** até que novas requisições chegassem citeturn0search8.  
  - Ao receber uma requisição, o servidor **reusa** imediatamente uma thread livre do pool, em vez de gerar uma nova. Após concluir o `service()`, a thread **volta ao pool** para o próximo trabalho.

### 2.3 Parâmetros de Configuração
- **Tomcat (Connector Attributes):**  
  - `maxThreads` – número máximo de threads que o pool pode ter.  
  - `minSpareThreads` – número mínimo de threads ociosas mantidas prontas.  
  - `acceptCount` – tamanho da fila de requisições enfileiradas quando não há threads livres.  
- **Como funciona na prática:**  
  1. Chega uma requisição HTTP.  
  2. Se existe thread livre, ela é alocada imediatamente.  
  3. Se todas as threads estiverem ocupadas, a requisição entra em **fila** até que alguma thread libere.  
  4. Se a fila atingir o limite (`acceptCount`), o servidor recusa a requisição (geralmente com HTTP 503). citeturn0search5  

### 2.4 Benefícios do Pooling
1. **Menor Latência** – threads pré‑alocadas eliminam o custo de criação em tempo de requisição.  
2. **Uso Controlado de Recursos** – definindo limites no pool, evita‑se explosões de uso de CPU/memória.  
3. **Estabilidade sob Carga** – enfileiramento previsível quando o tráfego excede a capacidade configurada.  
4. **Facilidade de Tuning** – você ajusta `maxThreads` e `acceptCount` conforme o perfil de carga da sua aplicação. citeturn0search3  

--- 

## 3. Servidores Embutidos nas Linguagens

### 3.1. Java Servlet API (1997)
- **Container:** Tomcat, Jetty, etc.  
- **Threads ou conexões agrupadas** para evitar fork/exec por requisição  
- **Mapeamento de rotas:**  
  - `web.xml` (XML)  
  - Depois, _annotations_ em classes `HttpServlet`  

### 3.2. Python BaseHTTPServer / WSGI (2000–2003)
- **BaseHTTPServer:** `HTTPServer` + `BaseHTTPRequestHandler`  
- **WSGI:** interface unificada _server ←→_ framework  
  ```python
  def app(environ, start_response):
      status = '200 OK'
      headers = [('Content-type', 'text/plain')]
      start_response(status, headers)
      return [b"Olá, mundo!"]
  ```

### 3.3. Ruby Rack e Sinatra (2007)
- **Rack:** padroniza interface entre servidor e framework  
- **Sinatra:** DSL minimalista para rotas  
  ```ruby
  require 'sinatra'
  get '/hello' do
    "Olá, mundo!"
  end
  ```

---

## 4. A Era do Event‑Loop: Node.js

- **Lançamento (2009):** V8 + event loop  
- **Exemplo mínimo:**
  ```js
  const http = require('http');
  http.createServer((req, res) => {
    if (req.url === '/ping') {
      res.end('pong');
    }
  }).listen(3000);
  ```
- **Express.js (Connect → Express):**
  ```js
  const express = require('express');
  const app = express();
  app.get('/users', handler);
  app.post('/login', handler);
  app.listen(3000);
  ```

---

## 5. Go e o Pacote `net/http`

- **Conceito de servidor e roteador integrados**  
- **Interface simples:**
  ```go
  mux := http.NewServeMux()
  mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintln(w, "Olá, mundo!")
  })
  http.ListenAndServe(":8080", mux)
  ```
- **`Handler` minimalista:** qualquer tipo que implemente `ServeHTTP(ResponseWriter, *Request)`  
- **Bibliotecas adicionais:** Gorilla Mux, Echo, Fiber, etc.

---

## 5. Por que Esse Padrão Persiste Hoje?

1. **Clareza de responsabilidades**  
   - Servidor: aceita conexões  
   - Roteador/handlers: processam cada rota  

2. **Modularidade e middleware**  
   - Autenticação, logging, CORS aplicáveis globalmente ou por rota  

3. **Desempenho e escalabilidade**  
   - Thread‑pool ou event‑loop vs fork/exec  

4. **Consistência entre ecossistemas**  
   - Padrões como Servlet, WSGI, Rack, Middleware/Router facilitam adoção e migração de conhecimento  

---

## 6. 10 Prompts para Assuntos Complementares

1. **Como funciona internamente o modelo _thread‑per‑request_ em servidores Java?**  
2. **O que são e como implementar middlewares HTTP em Go usando `http.Handler`?**  
3. **Quais as diferenças de performance entre servidores HTTP baseados em _event‑loop_ e em _thread‑pool_?**  
4. **Como o HTTP/2 altera o modelo de roteamento e multiplexação de requisições?**  
5. **Passo a passo para criar um roteador customizado em Python sem usar frameworks WSGI.**  
6. **Arquitetura interna do Express.js: como os _middlewares_ são encadeados?**  
7. **Implementando HTTP/3 e QUIC: desafios e vantagens em servidores modernos.**  
8. **Comparativo entre Gorilla Mux, Echo e Fiber no ecossistema Go.**  
9. **Como integrar balanceamento de carga (NGINX, HAProxy) com servidores HTTP customizados?**  
10. **Estratégias de cache (CDN, cabeçalhos HTTP) e seu impacto em arquiteturas de microsserviços.**


### Referências
- **StackOverflow:** explicação do modelo thread‑per‑request puro e seu ciclo de vida — citeturn0search1  
- **Tutorialspoint:** definição e ideia geral de thread pools em Java — citeturn0search8  
- **Baeldung:** como configurar pools de threads em servidores Java (Tomcat, GlassFish etc.) — citeturn0search5  
- **Medium (Oskar uit de Bos):** descrição de como Tomcat usa thread pool para cada requisição — citeturn0search3 