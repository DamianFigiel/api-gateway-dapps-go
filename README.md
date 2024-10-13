# API Gateway for Decentralized Applications

## Overview
This project is an API Gateway built in Go that sits in front of decentralized applications (dApps) and manages requests to various blockchain networks, currently supporting Ethereum and Polygon. It includes features like JWT-based authentication, rate limiting, and the ability to fetch the latest blockchain block numbers.

### Features
- **JWT Authentication**: Secures the API with JSON Web Tokens (JWT).
- **Rate Limiting**: Limits the number of requests a client can make to prevent abuse.
- **Blockchain Network Integration**: Routes requests to Ethereum and Polygon networks to fetch the latest block numbers.
- **Modular Middleware Design**: Easily extendable for additional middlewares or blockchain networks.

## Getting Started

### Prerequisites
- Go 1.16+ installed on your machine.
- API keys or RPC URLs for Ethereum (e.g., via Infura) and Polygon RPC endpoints.
- Environment variables configured (details below).

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/api-gateway-dapp.git
   cd api-gateway-dapp
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```
3. Set up the environment variables. Create a .env file in the root of the project:
    ```env
    JWT_SECRET=your_secret_key
    RATE_LIMIT_PER_MINUTE=60
    ETHEREUM_ENDPOINT=https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID
    POLYGON_ENDPOINT=https://polygon-mainnet.infura.io/v3/your-infura-api-key
### Usage
1. Build and run module:
    ```bash
    go build
    go run .
    ```
2. Send Requests:
- Use a tool like curl or Postman to send requests to the gateway.
- Add an Authorization header with a valid JWT token for access.
    ```bash
    curl -H "Authorization: Bearer <your_jwt_token>" -X POST http://localhost:8080/api/v1/blockchain -d '{"network": "ethereum"}'
    ```
3. Sample Response:
    ```json
    {
        "block_number": "14330000"
    }
    ```
### Environment Variables
- `JWT_SECRET`: Secret key for encoding and decoding JWT tokens.
- `RATE_LIMIT_PER_MINUTE`: Number of requests allowed per minute.
- `ETHEREUM_ENDPOINT`: RPC endpoint for Ethereum network.
- `POLYGON_ENDPOINT`: RPC endpoint for Polygon network.

## Project Structure
```bash
api-gateway-dapp/
├── main.go                       # Entry point of the application
├── auth_middleware.go            # JWT Authentication middleware
├── rate_limit_middleware.go      # Rate limiting middleware
├── handlers.go                   # Request handlers for blockchain interactions
├── auth_middleware_test.go       # Unit tests for JWT middleware
├── rate_limit_middleware_test.go # Unit tests for rate limiting middleware
├── handlers_test.go              # Unit tests for blockchain handlers
└── .env                          # Environment variables (not included in version control)
```
## Testing
To run the tests for this project, use:
```bash
go test -v ./...
```
## Extending the Project
1. Adding New Blockchain Networks.
2. Adding More Data to Fetch.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

