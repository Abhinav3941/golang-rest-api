A RESTful API built with Go, utilizing the Gin framework, SQLite for data storage, and Docker for containerization. This project also integrates GitHub Actions for Continuous Integration (CI).

Features
RESTful Endpoints: Provides CRUD operations for managing resources.
SQLite Database: Lightweight and serverless database solution.
Dockerized Application: Simplifies deployment and environment consistency.
Continuous Integration: Automated testing and build processes via GitHub Actions.




Tech Stack
Language: Go
Framework: Gin
Database: SQLite
Containerization: Docker
CI/CD: GitHub Actions


Prerequisites
Go 1.21 or later
Docker (optional, for containerization)


Clone the repository:
git clone https://github.com/Abhinav3941/golang-rest-api.git
cd golang-rest-api



Install dependencies:
go mod tidy


Run the application:
go run main.go


Using Docker

Build the Docker image:
docker build -t golang-rest-api .

Run the Docker container:
docker run -p 8080:8080 golang-rest-api


The API will be accessible at http://localhost:8080.

