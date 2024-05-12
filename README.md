# Gin and Gorm Starter Project

This is a starter project for building web applications using the Gin framework for routing and the Gorm library for database operations in Go. It follows the Model-View-Controller (MVC) architectural pattern with a structured internal directory containing configs, controllers, models, repositories, services, and utils.

## Features

- **Gin Framework**: Gin is a lightweight and fast web framework for Go. It provides a flexible and intuitive way to define routes and handle incoming requests.
- **Gorm Library**: Gorm is an Object-Relational Mapping (ORM) library for Go. It simplifies database operations by providing a convenient API for interacting with the database.
- **Structured Directory**: The project follows a structured directory structure, making it easy to navigate and maintain. This includes separate directories for configs, controllers, models, repositories, services, and utils.
- **MVC Pattern**: The project follows the Model-View-Controller (MVC) architectural pattern, separating concerns into distinct components.

## Getting Started

To get started with the Gin and Gorm starter project:

1. Clone the repository: `git clone https://github.com/virak1510/gin-gorm-starter.git`
2. Install the dependencies: `go mod download`
3. Configure the project by modifying the configuration files in the `internal/configs` directory.
4. Start the application: `go run ./cmd` or `air` if you install air

## Directory Structure

The project follows a structured directory structure:

- `internal`: This directory contains the main application code.
  - `configs`: Stores configuration files.
  - `controllers`: Houses the route handlers and business logic.
  - `models`: Holds the data models and database operations.
  - `repositories`: Handles data access and persistence.
  - `services`: Contains the application's business logic.
  - `utils`: Holds utility functions and helper methods.

## Contributing

If you'd like to contribute to the project, please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Make your changes
4. Commit your changes (`git commit -am 'Add new feature'`)
5. Push to the branch (`git push origin feature/your-feature`)
6. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Acknowledgements

- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/go-gorm/gorm)

Feel free to customize this Markdown file to add more details specific to your project! 
