
# Learn SQL with Go

This project is a basic practice application for interacting with databases using Go. It includes essential components to establish a connection, perform CRUD (Create, Read, Update, Delete) operations, and manage configuration files.

## Features

- **Database Connection**: Establish a secure and reliable connection to a database.
- **CRUD Operations**: Basic functionality to create, read, update, and delete records.
- **Configuration Management**: Read and parse settings from a JSON file.
- **Go Module Support**: Dependency management using Go modules.

## Project Structure

```
learn_sql/
├── config.go           // Loads configuration settings
├── config.json         // JSON file with database configuration
├── connection_check.go // Script to test the database connection
├── crud.go             // Handles CRUD operations
├── db.go               // Core database connection logic
├── go.mod              // Go module configuration
├── go.sum              // Go module dependencies
├── main.go             // Entry point of the application
```

## Requirements

- Go 1.18 or later
- A running database instance (e.g., MySQL, PostgreSQL)
- Properly configured `config.json` file

## Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd learn_sql
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure your `config.json` file with the database credentials:
   ```json
   {
       "host": "localhost",
       "port": 3306,
       "user": "your-username",
       "password": "your-password",
       "dbname": "your-database-name"
   }
   ```

4. Test the database connection:
   ```bash
   go run connection_check.go
   ```

5. Run the application:
   ```bash
   go run main.go
   ```

## Usage

- Use `crud.go` to interact with the database.
- Modify `config.json` for different environments.

## License

This project is licensed under the MIT License.
