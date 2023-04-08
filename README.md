# gin-api-rest
API REST usando o GIN framework para Golang

Para rodar a aplicação é necessário ter o Docker instalado. Abra um terminal e execute o comando 'docker-compose build' para criar ou atualizar a imagem do container e 
'docker-compose up' para iniciar os serviços de banco de dados. 

Em seguida utilizar 'GO RUN main.go' para inicializar a aplicação. 

Acessar no navegador o endereço http://localhost:8080/index para visualizar a API. As rotas podem ser testadas com uma aplicação como o postman. Elas estão listadas
no arquivo 'routes.go'.

Caso seja solicitado um caminho que não existe a API vai devolver uma página de erro 404.

:D

Versão do documento 0.1. Vai melhorar no futuro.
