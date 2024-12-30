Projeto final do curso de golang

Sistema de Gerenciamento de Usuários

Descrição:
	Crie um sistema simples de gerenciamento de usuários que permita adicionar, 
	listar, atualizar e deletar usuários de um banco de dados MySQL. O sistema deve
	ser implementado em Go e deve seguir os princípios da Clean Architecture. Além 
	disso, você deverá usar as seguintes ferramentas:

		- GORM para manipulação de dados do banco.
		- Zaplogger para criação de logs.
		- Gin para roteamento de chamadas.
		- Viper para inicialização de arquivos de configuração. (Não usar yaml e sim env vars)

Requisitos:

	Banco de Dados:

		Crie uma tabela chamada users no MySQL com as seguintes colunas:
			id (INT, AUTO_INCREMENT, PRIMARY KEY)
			name (VARCHAR(100))
			email (VARCHAR(100), UNIQUE)
			age (INT)

	Operações CRUD:

		Create (Criar):
			- Adicione um novo usuário ao banco de dados.
			- O email deve ser único e não pode ser nulo.
			- A idade deve ser um número inteiro positivo.

		Read (Ler):

			- Liste todos os usuários com suporte a paginação (ex: listar usuários de 
			10 em 10).
			- Obtenha um usuário específico pelo ID.

		Update (Atualizar):

			- Atualize as informações de um usuário existente.
			- O email não pode ser alterado se já existir outro usuário com o mesmo 
			  email.

		Delete (Deletar):

			- Remova um usuário do banco de dados pelo ID.

Estrutura do Projeto:

	O projeto deve ser organizado em pacotes Go seguindo os princípios da Clean 
	Architecture. A estrutura de pastas deve ser algo como:

		│
		├── /adapter
		│      ├── /api
		│      │	└── ...       
		│      └──....
		│
		├── /core
		│   ├── /usecase         
		│   ├── /repository     
		│   └── /domain          
		│
		├── /Infrastrcuture                 
		│		└── ...
		└── main.go

Requisitos Técnicos:
	
	GORM:
		- Use o GORM para interagir com o banco de dados MySQL.
		- Configure o GORM para usar o MySQL e defina os modelos de dados.

	Zaplogger:

		- Configure o Zaplogger para registrar logs em um arquivo no diretório /logs.
		- Registre logs para todas as operações CRUD (ex: "Usuário criado com sucesso", "Erro ao atualizar usuário").

	Gin:

		Configure o Gin para criar uma API RESTful com os seguintes endpoints:

			POST /users → Criar um usuário.
			GET /users → Listar todos os usuários (com paginação).
			GET /users/:id → Obter um usuário pelo ID.
			PUT /users/:id → Atualizar um usuário pelo ID.
			DELETE /users/:id → Deletar um usuário pelo ID.

	Viper:

		Use o Viper para carregar configurações.

	Clean Architecture:

		O projeto deve seguir os princípios da Clean Architecture, com camadas bem 
		definidas (controller, use case, repository).

	Tratamento de Erros:

		Todas as operações devem tratar erros de forma adequada.
		Retorne respostas HTTP apropriadas (ex: 404 para usuário não encontrado, 400 
		para dados inválidos).

	Logs:
		O sistema deve registrar logs para todas as operações CRUD.


Validações:

	Adicione validações para garantir que os dados inseridos sejam válidos 
	(ex: email único, idade positiva).

	Use o pacote validator para validar os dados recebidos nas requisições.

Docker:

	Crie um Dockerfile e um docker-compose.yml para rodar a aplicação e o banco de 
	dados SQL em contêineres.

Documentação:

	Adicione no repositório do código uma documentação básica sobre como rodar o 
	projeto.