# app-redesocial 

# Visão Geral - app-redesocial 
--------------------------------------------
1 - oque é
2 - Componentes
--------------------------------------------

# Oque é:
Será uma rede social onde possivelmente podera criar publicações que contenham apenas texto.
As entidades que terá na aplicação são:
1 - Usuários
    > CRUD
    > Seguir outro usuário
    > Parar de seguir outro usuário
    > Buscar todos os usuários que segue 
    > Buscar todos os usuários que são seguidos
    > Atualizar Senha 

    Duas tabelas no Banco de dados 
        > Usuários
        > Seguidores

2 - Publicações
    > CRUD
    > Buscar publicações de acordo com os usuários que segue 
    > Curtir

    Apenas uma tabela no Banco de dados
        > Publicações


# Componentes:
Teremos dois componentes para o desenvolvimente da aplicação que são:
1 - API (Back-End)
2 - Web App (Front-End)


----------------------------------------------------------------------------------
# Estrutura da API 
1 - Estrutura da aplicação.
2 - Pacotes

A API será o meio de comunicação entre o aplicação web e o banco de dados


2 - Pacotes:
Os pacotes da aplicação pode ser divididos em dois tipos:
    > Pacotes Principais
    > Pacotes Auxiliares 

Pacotes Principais:
    > Main
    > Router
    > Controllers
    > Modelos
    > Repositórios

Pacotes Auxiliares:
    > Config
    > Banco
    > Autenticação
    > Middleware
    > Segurança 
    > Respostas
    