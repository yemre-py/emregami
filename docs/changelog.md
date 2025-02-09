# Changelog

## [1.0.0] - 2024

### Added

- Authentication system with JWT token support
  - Access token and refresh token generation
  - Token validation functionality
  - Configurable token expiry times
- User registration and login functionality
  - Secure password hashing using bcrypt
  - Email and username uniqueness validation
  - UUID generation for user IDs
- Input validation system
  - Username validation rules:
    - 3-20 characters length
    - Must start with a letter
    - Can contain letters, numbers, and underscores
    - No consecutive underscores
    - Must end with letter or number
  - Password validation rules:
    - 8-256 characters length
    - Must contain letters, numbers, and special characters
  - Email validation with regex pattern
- Environment configuration system
  - .env file support using Viper
  - Type-safe configuration getters
- Database integration
  - GORM support for database operations
  - Repository pattern implementation
  - Unique constraints on email and username

### Dependencies

- go v1.23.3
- github.com/go-playground/validator v9.31.0
- github.com/golang-jwt/jwt/v4 v4.5.1
- github.com/google/uuid v1.6.0
- github.com/spf13/viper v1.19.0
- golang.org/x/crypto v0.33.0
- gorm.io/gorm v1.25.12

### Security

- Password hashing using bcrypt
- JWT token-based authentication
- Input validation for all user inputs
- Environment variable based configuration
