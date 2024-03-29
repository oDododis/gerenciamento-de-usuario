# Gerenciador de Usuário Simples.
## Item necessario para roda-lo:
- A linguagem utilizada é: [Go](https://golang.org/dl/)
- Para enviar as requisições localmente: [Postman Desktop](https://www.postman.com/downloads/ )

## Instalação:
Para instalar é apenas clonar neste repositório:
````
git clone https://github.com/oDododis/gerenciamento-de-usuario/
cd gerenciamento-de-usuario/
````
## Inicialização:
Antes de iniciar o programa precisamos criar as tablas do banco vazias, com este comando:

````
go run migration/main.go
````

Para iniciar o gerenciador execute este comando:
````
go run main.go
````

Apos aparecer isso no console:
````
[GIN-debug] Listening and serving HTTP on :8080
````
o servidor estará aberto.

## Usando o gerenciador:
Temos as seguintes Endpoints no servidor:

### 1. **POST /createUser**
   
 Este Endpoint cria o usuário. Para criar o usuário precisamos ter as segunites informações, com os seguintes requisitos:
  - Nome completo, com um mínimo de 3 e um máximo de 150 caracteres.
  - Email, com o formato de email.
  - Username, com um mínimo de 3 e um máximo de 150 caracteres.
  - Senha, com um mínimo de 8 caracteres e precisa conter algum dos seguintes caracteres especiais: "!", "@", "#", "$", "%", "&", "*", "(", ")", "_" e "+".
 
 No Postman a requisição dos dados devem ser feita no Body e raw na formatação JSON:
 ![img.png](img/img.png)
   
 #### Exemplos:
 Usuáro 1:

        {
        "fullname": "Douglas Barbosa",
        "email": "douglas@barbosa.com",
        "username": "oDododis",
        "password": "1234567*"
        }
 Usuário 2:
        
        {
        "fullname": "Segundo Usuário",
        "email": "segundo@usuario.com",
        "username": "segUsuario",
        "password": "2222@@@@"
        }

Como resposta para casos de sucesso, retornará as seguintes informações:
#### Exemplo para `Status 201 Created`:

Informaçoes colocadas corretamente:
![img_1.png](img/img_1.png)

- Responderá com as informações de ID, Nome completo, Email e Username.

Como resposta para casos de fracasso, retornará informações sobre o erro:

#### Exemplos para `Status 400 Bad Request`: 
   
O nome completo não segue as restrições mínimas:
![img_2.png](img/img_2.png)

O email já existe no Banco de Dados:
![img_3.png](img/img_3.png)

- Responderá com as informações com tipo de erro.

### 2. **POST /login**
Este Endpoint faz o login do usuário. Para fazer o login precisamos ter as informações parecida com a do Create User, com os seguintes requisitos:
   - Email, com o formato de email.
   - Senha, com um mínimo de 8 caracteres e precisa conter algum dos seguires caracteres especiais: "!", "@", "#", "$", "%", "&", "*", "(", ")", "_" e "+".

O email e a senha devem ser o mesmo de algum usuario ques esteje no banco de dados e a requisição deve ser feita a mesma maneira do Create User, no Body, raw e em formatação JSON:
![img_4.png](img/img_4.png)
#### Exemplo:
        {
        "email": "douglas@barbosa.com",
        "password": "1234567*"
        }

Como resposta para casos de sucesso, retornará as seguintes informações:
#### Exemplo para `Status 202 Accepted`:
Informaçoes colocadas corretamente:
![img_5.png](img/img_5.png)
- Responderá com as informações do Token que será utilizado para as requisições restantes.

Como resposta para casos de fracasso, retornará informações sobre o erro:
#### Exemplo para `Status 400 Bad Request`:
Email invalido:
![img_6.png](img/img_6.png)
#### Exemplo para `Status 403 Forbidden`:
Senha incorreta:
![img_7.png](img/img_7.png)
#### Exemplo para `Status 404 Not Found`:
Email não existe no Banco de Dados
![img_8.png](img/img_8.png)
- Responderá com as informações com tipo de erro.

### 3. **PUT /updateUser/{ID do usuário}**
Este Endpoint atualiza o usuário. Para este Endpoint funcionar precisamos do Token criado no login. No postman devemos colocar no em Authorization com o tipo Bearer token:
![img_9.png](img/img_9.png)
Na url deve estar o ID do usuário:
#### Exemplo:
Com o ID do usuario 2:
`http://localhost:8080/updateUser/2`

E para a requisição dos dados temos que colocar os mesmos tipos de dados colocados no Create User, no Body, raw e com formatação em JSON:
 - Nome completo, com um mínimo de 3 e um máximo de 150 caracteres.
 - Email, com o formato de email.
 - Username, com um mínimo de 3 e um máximo de 150 caracteres.
 - Senha, com um mínimo de 8 caracteres e precisa conter algum dos seguintes caracteres especiais: "!", "@", "#", "$", "%", "&", "*", "(", ")", "_" e "+".

![img_10.png](img/img_10.png)
#### Exemplo:
    {
    "fullname": "Douglas Spigolon Barbosa",
    "email": "douglas@barbosa.com",
    "username": "dododis",
    "password": "0987654+"
    }
Como resposta para casos de sucesso, retornará as seguintes informações:
#### Exemplo para `Status 202 Accepted`:
Informações colocadas corretamente:
![img_11.png](img/img_11.png)
- Responderá com as informações de ID, Nome completo, Email e Username, atualizados.

Como resposta para casos de fracasso, retornará informações sobre o erro:
#### Exemplos para `Status 400 Bad Request`:
O email já existe no Banco de Dados:
![img_3.png](img/img_3.png)
ID invalido:
![img_14.png](img/img_14.png)
#### Exemplo para `Status 401 Unauthorized`:
Token invalido:
![img_12.png](img/img_12.png)
#### Exemplo para `Status 404 Not Found`:
ID não encontrado:
![img_15.png](img/img_15.png)
- Responderá com as informações com tipo de erro.

### 4. **GET /getUserID/{ID do usuário}**
Este Endpoint procura o usuário pelo ID. Para este Endpoint funcionar precisamos do Token criado no login. No postman devemos colocar no em Authorization com o tipo Bearer token:
![img_9.png](img/img_9.png)
Na url deve estar o ID do usuário:
#### Exemplo:
Com o ID do usuário 2:
`http://localhost:8080/getUserID/2`

Como resposta para casos de sucesso, retornará as seguintes informações:
#### Exemplo para `Status 202 Accepted`:
ID existe:
![img_13.png](img/img_13.png)
- Responderá com as informações de ID, Nome completo, Email e Username.

Como resposta para casos de fracasso, retornará informações sobre o erro:
#### Exemplo para `Status 400 Bad Request`:
ID invalido:
![img_14.png](img/img_14.png)
#### Exemplo para `Status 401 Unauthorized`:
Token invalido:
![img_12.png](img/img_12.png)
#### Exemplo para `Status 404 Not Found`:
ID não encontrado:
![img_15.png](img/img_15.png)
- Responderá com as informações com tipo de erro.


### 5. **GET /getUserEmail/{Email do usuário}**
Este Endpoint procura o usuário pelo Email. Para este Endpoint funcionar precisamos do Token criado no login. No postman devemos colocar no em Authorization com o tipo Bearer token:
![img_9.png](img/img_9.png)
Na url deve estar o Email do usuário:
#### Exemplo:
Com o email do usuario 2:
`http://localhost:8080/getUserID/douglas@barbosa.com`

Como resposta para casos de sucesso, retornará as seguintes informações:
#### Exemplo para `Status 202 Accepted`:
Email existe:
![img_13.png](img/img_13.png)
- Responderá com as informações de ID, Nome completo, Email e Username.

Como resposta para casos de fracasso, retornará informações sobre o erro:
#### Exemplo para `Status 401 Unauthorized`:
Token invalido:
![img_12.png](img/img_12.png)
#### Exemplo para `Status 404 Not Found`:
ID não encontrado:
![img_15.png](img/img_15.png)
- Responderá com as informações com tipo de erro.
### 6. **GET /getUserList**
Este Endpoint lista os usuários presentes no banco. Para este Endpoint funcionar precisamos do Token criado no login. No postman devemos colocar no em Authorization com o tipo Bearer token:
![img_9.png](img/img_9.png)

Como resposta para casos de sucesso, retornará as seguintes informações:
#### Exemplo para `Status 202 Accepted`:
Lista completa dos usuários no Banco de Dados:
![img_16.png](img/img_16.png)
- Responderá com as informações de ID, Nome completo, Email e Username, de todos do Banco.

Como resposta para casos de fracasso, retornará informações sobre o erro:
#### Exemplo para `Status 404 Not Found`:
Não tem usuários no Banco de Dados:
![img_17.png](img/img_17.png)
- Responderá com as informações com tipo de erro.

### 7. **DELETE /deleteUser/{ID do usuário}**
Este Endpoint exclue o usuário pelo ID. Para este Endpoint funcionar precisamos do Token criado no login. No postman devemos colocar no em Authorization com o tipo Bearer token:
![img_9.png](img/img_9.png)
Na url deve estar o ID do usuário:
#### Exemplo:
Com o ID do usuario 2:
`http://localhost:8080/deleteUser/2`

Como resposta para casos de sucesso, retornará as seguintes informações:
#### Exemplo para `Status 202 Accepted`:
ID existente:
![img_18.png](img/img_18.png)
- Responderá com a informação que o usuário do ID foi excluido.

Como resposta para casos de fracasso, retornará informações sobre o erro:
#### Exemplo para `Status 400 Bad Request`:
ID invalido:
![img_14.png](img/img_14.png)
#### Exemplo para `Status 401 Unauthorized`:
Token invalido:
![img_12.png](img/img_12.png)
#### Exemplo para `Status 404 Not Found`:
ID não encontrado:
![img_15.png](img/img_15.png)
- Responderá com as informações com tipo de erro.
