# Segurança em Go - Exemplo Prático

Este exemplo demonstra várias práticas de segurança em uma aplicação web Go, incluindo autenticação, CSRF, rate limiting, e outras medidas de proteção.

## Funcionalidades

1. **Autenticação**
   - JWT (JSON Web Tokens)
   - Middleware de autenticação
   - Rotas protegidas

2. **Proteção CSRF**
   - Geração de tokens
   - Validação de tokens
   - Cookie seguro

3. **Headers de Segurança**
   - X-Frame-Options
   - X-XSS-Protection
   - X-Content-Type-Options
   - Content-Security-Policy
   - HSTS

4. **Rate Limiting**
   - Limitação por IP
   - Configuração de burst
   - Proteção contra DDoS

5. **Outras Medidas**
   - HTTPS/TLS
   - Sanitização de entrada
   - Validação de dados
   - Graceful shutdown

## Pré-requisitos

1. Go 1.22 ou superior
2. Certificados TLS (para desenvolvimento)

## Configuração

1. Gerar certificados para desenvolvimento:
   ```bash
   # Gerar chave privada
   openssl genrsa -out key.pem 2048

   # Gerar certificado auto-assinado
   openssl req -new -x509 -sha256 -key key.pem -out cert.pem -days 365
   ```

2. Instalar dependências:
   ```bash
   go mod download
   ```

## Como Executar

1. Clone o repositório e entre no diretório:
   ```bash
   cd exemplos/05-seguranca
   ```

2. Execute o programa:
   ```bash
   go run main.go
   ```

## Testando a API

1. **Criar Usuário**
   ```bash
   curl -k -X POST https://localhost:8443/users \
     -H "Content-Type: application/json" \
     -d '{"username": "john", "password": "password123"}'
   ```

2. **Login**
   ```bash
   curl -k -X POST https://localhost:8443/login \
     -H "Content-Type: application/json" \
     -d '{"username": "john", "password": "password123"}'
   ```

3. **Acessar Rota Protegida**
   ```bash
   # Primeiro, fazer GET para obter token CSRF
   curl -k -c cookies.txt https://localhost:8443/protected

   # Depois, usar token no header
   curl -k -b cookies.txt \
     -H "Authorization: Bearer SEU_JWT_TOKEN" \
     -H "X-CSRF-Token: TOKEN_DO_COOKIE" \
     https://localhost:8443/protected
   ```

## Estrutura do Código

1. **Autenticação**
   - `createToken`: cria JWTs
   - `validateToken`: valida JWTs
   - `authMiddleware`: protege rotas

2. **CSRF**
   - `generateCSRFToken`: gera tokens
   - `validateCSRFToken`: valida tokens
   - `csrfMiddleware`: aplica proteção

3. **Rate Limiting**
   - `IPRateLimiter`: controla requisições por IP
   - Configuração de limites por segundo
   - Burst para picos de tráfego

4. **Segurança**
   - `securityHeadersMiddleware`: adiciona headers
   - `sanitizeInput`: limpa entrada do usuário
   - Configuração TLS

## Boas Práticas Implementadas

1. **Autenticação**
   - Tokens JWT com expiração
   - Secrets gerados aleatoriamente
   - Validação de assinatura

2. **Proteção de Dados**
   - Sanitização de entrada
   - Validação de dados
   - Headers de segurança

3. **Performance e Disponibilidade**
   - Rate limiting
   - Timeouts apropriados
   - Graceful shutdown

4. **HTTPS/TLS**
   - TLS 1.2+
   - HSTS
   - Certificados seguros

## Próximos Passos

1. **Melhorias de Segurança**
   - Implementar hash de senhas (bcrypt)
   - Adicionar autenticação em dois fatores
   - Implementar logout e revogação de tokens
   - Adicionar logging seguro

2. **Funcionalidades**
   - Refresh tokens
   - Roles e permissões
   - Auditoria de acessos
   - Backup e recuperação de dados

3. **Infraestrutura**
   - Configuração via ambiente
   - Métricas e monitoramento
   - Testes de segurança
   - CI/CD seguro 