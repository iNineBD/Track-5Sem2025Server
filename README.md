# Track-5Sem2025Server

This project is an API developed in Go that provides analytical data about projects registered in the Taiga system. The API is designed to facilitate the analysis of data such as statistics of tags, users, and card statuses.

## Installation and Execution Guide for Track-5Sem2025Server

This guide will teach you how to configure and run the project locally.

### Prerequisites

Before starting, make sure that:

- Go is installed on your machine.
  - If not, follow the instructions at [Go Documentation](https://go.dev/doc/install).
- You have an active internet connection to download the necessary dependencies.
- A PostgreSQL database is configured and running.

### Steps to Configure and Run

1. **Clone the Repository**  
   Open the terminal and run the command below to clone the repository:

   ```bash
   git clone https://github.com/ininetrack/Track-5Sem2025Server.git
   cd Track-5Sem2025Server/src
   ```

2. **Configure the Database**  
   Ensure that PostgreSQL is configured with the required credentials. Update the `.env` file with the database information:

   ```env
   DB_USER=user-database
   DB_PASSWORD=password-database
   DB_HOST=host-database
   DB_PORT=port-database
   DB_NAME=name-database
   DB_SCHEMA=schema-database
   EMAIL_HOST=email-host
   EMAIL_PORT=email-port
   EMAIL_HOST_USERNAME=email-host-user
   EMAIL_HOST_PASSWORD=email-host-password
   EMAIL_HOST_FROM=email-host-from
   ```

3. **Install Dependencies**  
   Run the command below to install the project dependencies:

   ```bash
   go mod tidy
   ```

4. **Run the Application**  
   Start the application with the command:

   ```bash
   go run main.go
   ```

5. **Access the Application in the Browser**  
   Once the application is running, open your preferred browser and enter the following address:

   ```
   http://localhost:8080/swagger/index.html
   ```

   Done! Now you are ready to explore the analytical data of Track-5Sem2025Server and use it in your application or for visualization.

### Authorization in Swagger

To authorize in Swagger, after copying the token, you need to prepend the word "Bearer" followed by a space before the token. For example:

```
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuZHJlbHVpejEwODhAZ21haWwuY29tIiwiZXhwIjoxNzQ2NTc2NDE2LCJyb2xlIjo5OTY1NjEyLCJ1c2VyX2lkIjoyfQ.AJIwXmhVofrykeamLzUQQxu7WkvZvfQc6cOzDt5-P7w
```

### Main Features

- **Project Query**: Displays all projects registered in the Taiga system.
- **Project Statistics**: Provides analytical data about cards by tags, cards by users, and cards by status of a project.

### Technologies Used

- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Documentation**: Swagger
- **CI/CD**: GitHub Actions

### For tests, it's necessary be inside the ./src directory

Check the files with the linter

```bash
golangci-lint run ./...
```

To correct the .go files following the GO convention, using the linter

```bash
go fmt ./...
```

Run all tests and validate the coverage

```bash
go test -v -coverprofile=coverage_report/coverage.out -covermode=atomic ./...
```

After execution, convert the generated coverage.out file to .html for better visualization:

```bash
go tool cover -html=coverage_report/coverage.out -o coverage_report/coverage.html
```

   - Open the generated HTML report in the browser (on Unix systems such as Linux or MacOS)

         ```bash
         xdg-open coverage_report/coverage.html
         ```

   - To MacOS, you can use:

      ```bash
      open coverage_report/coverage.html
      ```

   - On Windows, you can simply open the file directly or use:

      ```bash
      start coverage_report/coverage.html
      ```
---
