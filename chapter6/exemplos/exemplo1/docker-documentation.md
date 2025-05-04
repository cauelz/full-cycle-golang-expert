# Docker Compose - MySQL Setup

Este documento explica detalhadamente cada componente do arquivo `docker-compose.yml` que configura um container MySQL.

## Estrutura do Docker Compose

```yaml
version: '3'

services:
  mysql:
    image: mysql:5.7
    container_name: mysql
    restart: always
    platform: linux/amd64
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: goexpert
      MYSQL_PASSWORD: root
    ports:
    - 3306:3306
```

## Explicação Linha por Linha

### Versão do Docker Compose
```yaml
version: '3'
```
- Define a versão do formato do arquivo Docker Compose que está sendo utilizada
- A versão 3 é uma das mais recentes e estáveis, oferecendo suporte a diversos recursos modernos do Docker

### Definição dos Serviços
```yaml
services:
```
- Seção onde todos os serviços (containers) são definidos
- Cada serviço representa um container que será executado

### Configuração do MySQL
```yaml
mysql:
```
- Nome do serviço que está sendo definido
- Este nome será usado como hostname na rede Docker

### Imagem do Container
```yaml
image: mysql:5.7
```
- Especifica qual imagem Docker será utilizada
- `mysql:5.7` indica:
  - `mysql`: o nome da imagem oficial do MySQL
  - `5.7`: a versão específica do MySQL que será utilizada

### Nome do Container
```yaml
container_name: mysql
```
- Define um nome personalizado para o container
- Este nome será usado para referenciar o container no Docker
- Facilita a identificação do container em comandos Docker

### Política de Reinicialização
```yaml
restart: always
```
- Define a política de reinicialização do container
- `always`: o container sempre será reiniciado automaticamente
  - Em caso de falha
  - Quando o Docker daemon reiniciar
  - Quando o host reiniciar

### Plataforma
```yaml
platform: linux/amd64
```
- Especifica a arquitetura do sistema operacional para o container
- Importante especialmente quando se trabalha com diferentes arquiteturas (ex: Mac M1)
- `linux/amd64` é a arquitetura x86_64 padrão

### Variáveis de Ambiente
```yaml
environment:
  MYSQL_ROOT_PASSWORD: root
  MYSQL_DATABASE: goexpert
  MYSQL_PASSWORD: root
```
- Define as variáveis de ambiente necessárias para configurar o MySQL:
  - `MYSQL_ROOT_PASSWORD`: define a senha do usuário root
  - `MYSQL_DATABASE`: cria um banco de dados com este nome durante a inicialização
  - `MYSQL_PASSWORD`: define a senha para acesso ao banco de dados

### Mapeamento de Portas
```yaml
ports:
- 3306:3306
```
- Configura o mapeamento de portas entre o host e o container
- `3306:3306` significa:
  - Primeira porta (3306): porta no host (sua máquina)
  - Segunda porta (3306): porta no container
  - Permite acessar o MySQL através da porta 3306 do seu computador

## Como Usar

1. Certifique-se de ter o Docker e Docker Compose instalados
2. No diretório do arquivo `docker-compose.yml`, execute:
   ```bash
   docker-compose up -d
   ```
3. Para parar o container:
   ```bash
   docker-compose down
   ```

## Conexão ao MySQL

- Host: localhost
- Porta: 3306
- Usuário: root
- Senha: root
- Banco de Dados: goexpert

## Observações de Segurança

⚠️ **Importante**: As senhas definidas neste arquivo são para ambiente de desenvolvimento. Em produção:
- Nunca use senhas simples como "root"
- Considere usar variáveis de ambiente ou secrets do Docker
- Implemente políticas de segurança mais robustas 