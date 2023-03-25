# Template for Log In, Sign Up, and Logout with JSON Web Tokens and Bootstrap in Golang
This is a template for a login, signup, and logout system built with JSON Web Tokens and Bootstrap in Golang. Before you can use this template, you'll need to set up a few things:
## Prerequisites
- Go installed on your machine
- A database set up (MySQL)
## Environment Variables
The template requires an `.env` file to be present in the root directory. The `.env` file should contain the following variables:
- **PORT**: The port number that the application will run on. For example, 8080.
- **DB_HOST**: The host of the database. This can be localhost or an IP address or domain name.
- **DB_USER**: The username used to connect to the database.
- **DB_PASS**: The password used to connect to the database.
- **DB_NAME**: The name of the database that the application will use.
- **JWT_SECRET_KEY**: A random string used to sign and verify JSON Web Tokens.
### Example
```env
PORT=8080
DB_HOST=127.0.0.1:3306
DB_USER=root
DB_PASS=root
DB_NAME=myapplicationdb
JWT_SECRET_KEY=abc123
```
## Installation
1. Clone this repository:
```go
go get https://github.com/dennisw05/Golang-Log-In-Sign-Up-and-Logout-template-with-JSON-Web-Tokens-and-Boostrap
```
2. Navigate to the project directory:
```go
cd Golang-Log-In-Sign-Up-and-Logout-template-with-JSON-Web-Tokens-and-Boostrap
```
3. Install dependencies:
```go
go mod download
```
## Usage
1. Start the server:
```go
go run main.go
```
2. Open your browser and go to `http://localhost:<PORT>`.
3. You should see the login page with a link to the signup page.
## Contributing
Contributions are welcome! If you want to contribute to this project, please fork this repository and create a pull request.
## License
This project is licensed under the MIT License - see the [`LICENSE`](https://github.com/dennisw05/Golang-Log-In-Sign-Up-and-Logout-template-with-JSON-Web-Tokens-and-Boostrap/blob/main/LICENSE) file for details.
## Acknowledgments
[Bootstrap](https://getbootstrap.com/)

[JSON Web Tokens](https://jwt.io/)
