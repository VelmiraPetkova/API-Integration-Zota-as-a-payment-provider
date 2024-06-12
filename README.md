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
<h3>Install Dependencies</h3>

<p>There are no external dependencies for this project other than the Go standard library.</p>

<h3>Functions</h3>
<p>GenerateSignature(endpointID string, merchantOrderID string, orderAmount string, customerEmail string, merchantSecretKey string) string: 
Generates a SHA-256 signature required for the deposit request.</p>
<p>MakeDepositRequest(request *DepositRequest) (*DepositResponse, error): Makes a deposit request to the ZotaPay API and returns the response.</p>

<h3> Example Output </h3>
<img width="773" alt="image" src="https://github.com/VelmiraPetkova/API-Integration-Zota-as-a-payment-provider/assets/104137851/73043e5e-93a5-45de-8ea8-02578b346cfe">

