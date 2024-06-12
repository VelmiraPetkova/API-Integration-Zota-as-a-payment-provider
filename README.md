# ZotaPay Integration

This project demonstrates a simple integration with the ZotaPay payment provider. The application implements a deposit (non-credit card) request to the ZotaPay API.

## Prerequisites

Ensure you have the following installed:

- Go (version 1.15 or later)
- Git

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/your-repo.git
   cd your-repo
   ```
Install Dependencies

There are no external dependencies for this project other than the Go standard library.

Functions
generateSignature(endpointID string, merchantOrderID string, orderAmount string, customerEmail string, merchantSecretKey string) string: Generates a SHA-256 signature required for the deposit request.
makeDepositRequest(request *DepositRequest) (*DepositResponse, error): Makes a deposit request to the ZotaPay API and returns the response.

Example Output
