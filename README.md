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

### Main Features

- **Project Query**: Displays all projects registered in the Taiga system.
- **Project Statistics**: Provides analytical data about cards by tags, cards by users, and cards by status of a project.

### Technologies Used

- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Documentation**: Swagger
- **CI/CD**: GitHub Actions
