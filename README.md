# Gerenciador de Usuario Simples.


## Item necessario para rodar-lo:
- A linguagem utilizada é: [Go](https://golang.org/dl/)
- Para enviar as requisições: [Postman](https://www.postman.com/downloads/ )

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

Para iniciar o gerenciador execute este commando:
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
   
 Este Endpoint cria o usuário. Para criar o usuario precisamos ter as seguntes informações, com os seguntes requisitos:
  - Nome completo, com um mínimo de 3 e um máximo de 150 caracteres.
  - Email, com o formato de email.
  - Username, com um mínimo de 3 e um máximo de 150 caracteres.
  - Senha, com um mínimo de 8 caracteres e precisa conter algum dos seguites caracteres especiais: "!", "@", "#", "$", "%", "&", "*", "(", ")", "_" e "+".
 
 No Postman a requisição dos dados devem ser feita no Body e raw na formatação JSON:
 ![img.png](img.png)
   
 #### Exemplo:
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

 Como resposta para casos de susseso, retornará as seguintes informações:
#### Exemplo para `Status 201 Created`:

   Informaçoes colocadas corretamente:
   ![img_1.png](img_1.png)

   Como resposta para casos de fracasso, retornará informações sobre o erro:

#### Exemplos para `Status 400 Bad Request`: 
   
O nome completo não segue as restrioções mínimas:
![img_2.png](img_2.png)

O email já existe no Banco de Dados:
![img_3.png](img_3.png)

### 2. **POST /login**
   - Este Endpoint faz o login do usuário.

3. **PUT /updateUser/{ID do usuário}**
   - asd

4. **GET /getUserID/{ID do usuário}**
   - asd

5. **GET /getUserEmail/{Email do usuário}**
   - asd

6. **GET /getUserList**
   - asd

7. **DELETE /deleteUser/{ID do usuário}**
   - asd