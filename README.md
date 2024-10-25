# app-redesocial

## Visão Geral

### 1 - O que é?
O **app-redesocial** será uma plataforma de rede social onde os usuários poderão criar publicações que contenham apenas texto. Ele incluirá funcionalidades básicas de interação entre usuários, como seguir e deixar de seguir outros usuários, além de curtidas em publicações. A aplicação terá as seguintes entidades principais:

#### 1.1 - Usuários
- Operações CRUD (Criar, Ler, Atualizar, Deletar)
- Seguir outro usuário
- Parar de seguir outro usuário
- Buscar todos os usuários que segue
- Buscar todos os usuários que são seguidos
- Atualizar senha

O banco de dados terá duas tabelas relacionadas a usuários:
- **Usuários**: para armazenar os dados básicos dos usuários.
- **Seguidores**: para manter o relacionamento de quem segue quem.

#### 1.2 - Publicações
- Operações CRUD
- Buscar publicações de acordo com os usuários que segue
- Curtir publicações

O banco de dados terá uma tabela relacionada a publicações:
- **Publicações**: para armazenar os posts de texto.

### 2 - Componentes
A aplicação será dividida em dois componentes principais:

#### 2.1 - API (Back-End)
A API será o ponto de comunicação entre o front-end (Web App) e o banco de dados, manipulando todas as requisições e respostas.

#### 2.2 - Web App (Front-End)
O Web App será a interface de usuário, permitindo que os usuários interajam com o sistema, façam publicações, sigam outros usuários e curtam postagens.

---

## Estrutura da API

### 1 - Estrutura da Aplicação
A API será responsável por processar as requisições recebidas do front-end, acessando e manipulando os dados no banco de dados. A comunicação será feita através de endpoints definidos, seguindo um modelo RESTful.

### 2 - Pacotes
A API será dividida em pacotes principais e auxiliares para modularizar o código e manter a separação de responsabilidades.

#### 2.1 - Pacotes Principais
- **Main**: Ponto de entrada da aplicação.
- **Router**: Definição das rotas e endpoints da API.
- **Controllers**: Lógica de controle, manipulando as requisições HTTP.
- **Modelos**: Definição das estruturas de dados (modelos) que representam as tabelas do banco de dados.
- **Repositórios**: Acesso direto ao banco de dados para as operações CRUD.

#### 2.2 - Pacotes Auxiliares
- **Config**: Arquivos de configuração da aplicação (banco de dados, servidor, etc).
- **Banco**: Conexão e gerenciamento do banco de dados.
- **Autenticação**: Lógica de autenticação e controle de usuários.
- **Middleware**: Funções intermediárias que processam as requisições antes de chegarem aos controllers.
- **Segurança**: Gerenciamento de segurança e proteção da API (ex: proteção contra ataques CSRF, injeção de SQL, etc).
- **Respostas**: Manipulação de respostas HTTP padronizadas para o front-end.

---

Esse README serve como um guia inicial para entender a estrutura e os principais componentes do **app-redesocial**.

---
## Execução de comandos dentro da api/
- **go mod init api**: vai criar o modulo para a api 
- **go get github.com/gorilla/mux**: uso para configurações de routers