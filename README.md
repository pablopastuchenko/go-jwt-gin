# go-jwt-gin

Este projeto é uma API desenvolvida em Go utilizando o framework Gin. Ele implementa autenticação JWT (JSON Web Token) para proteger rotas e gerenciar sessões de usuários.

---

## 🔧 Configuração do Ambiente

1. **Clone o Repositório**
   ```bash
   git clone https://github.com/pablopastuchenko/go-jwt-gin.git
   cd go-jwt-gin
   ```

2. **Configure as Dependências**
   Certifique-se de ter o Go instalado em sua máquina. Para baixar as dependências do projeto, execute:
   ```bash
   go mod tidy
   ```

3. **Configure as Variáveis de Ambiente**
   Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```env
   PORT=8080
   JWT_SECRET=your_jwt_secret
   ```
   - **PORT**: Porta onde a aplicação será executada.
   - **JWT_SECRET**: Chave secreta usada para assinar os tokens JWT.

4. **Inicie o Servidor**
   Execute o seguinte comando para iniciar o servidor:
   ```bash
   go run main.go
   ```

---

## 🚀 Endpoints

### **Rotas Públicas**

1. **Cadastrar Usuário**
   - **Rota**: `/register`
   - **Método**: `POST`
   - **Descrição**: Registra um novo usuário no sistema.
   - **Body**:
     ```json
     {
       "first_name": "John",
       "last_name": "Doe",
       "email": "john.doe@example.com",
       "password": "123456"
     }
     ```

2. **Login**
   - **Rota**: `/login`
   - **Método**: `POST`
   - **Descrição**: Gera um token JWT para autenticação do usuário.
   - **Body**:
     ```json
     {
       "email": "john.doe@example.com",
       "password": "123456"
     }
     ```

   - **Resposta**:
     ```json
     {
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
     }
     ```

### **Rotas Protegidas**

> Todas as rotas protegidas exigem um cabeçalho `Authorization` com o token JWT:
> ```
> Authorization: Bearer <token>
> ```

1. **Obter Informações do Usuário**
   - **Rota**: `/user`
   - **Método**: `GET`
   - **Descrição**: Retorna as informações do usuário autenticado.

2. **Atualizar Informações do Usuário**
   - **Rota**: `/user`
   - **Método**: `PUT`
   - **Descrição**: Atualiza os dados do usuário.
   - **Body**:
     ```json
     {
       "first_name": "Jane",
       "last_name": "Smith"
     }
     ```

3. **Deletar Conta**
   - **Rota**: `/user`
   - **Método**: `DELETE`
   - **Descrição**: Remove a conta do usuário.

---

## 🛠️ Arquitetura

- **Gin Framework**: Utilizado para criar a API e gerenciar rotas.
- **JWT**: Implementado para autenticação e autorização de usuários.
- **Middleware**: Middleware personalizado para validar tokens JWT e proteger rotas sensíveis.

---

## 🧪 Testando a API

1. **Utilize o Postman ou Insomnia**
   - Configure as requisições de acordo com as rotas descritas acima.

2. **Testando Rotas Protegidas**
   - Faça login para obter o token JWT.
   - Adicione o token no cabeçalho `Authorization` como `Bearer <token>` para acessar rotas protegidas.

---

## 📂 Estrutura do Projeto

```plaintext
go-jwt-gin/
├── helpers/
│   └── token_helper.go     # Funções para gerar e validar tokens JWT
├── middleware/
│   └── auth_middleware.go  # Middleware de autenticação
├── routes/
│   └── user_routes.go      # Definição das rotas de usuário
├── .env                    # Configurações de ambiente
├── go.mod                  # Gerenciador de dependências
├── main.go                 # Ponto de entrada da aplicação
```

---

## 🤝 Contribuições

1. Faça um fork do repositório.
2. Crie uma branch para sua feature: `git checkout -b minha-feature`.
3. Faça commit das alterações: `git commit -m 'Adiciona nova feature'`.
4. Envie para o repositório remoto: `git push origin minha-feature`.
5. Abra um Pull Request.

---

## 📝 Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---

**Desenvolvido por [Pablo Pastuchenko](https://github.com/pablopastuchenko)**

