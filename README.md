# Receipt Processor API

This project is a Receipt Processor API built in Go. It includes two main endpoints:
- **POST /receipts/process**: Processes a receipt and calculates points.
- **GET /receipts/{id}/points**: Retrieves points for a processed receipt by ID.

## Prerequisites

To run this application, ensure you have the following installed:
- **Go** (version 1.16 or later)
- **Docker** (if you want to use Docker for running the application)

## Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/sslcandoit/receipt-processor.git
   cd receipt-processor
   ```
2. **Install Dependencies**:
   ```bash
   go mod download
   ```
3. **Run the application locally:**

    To start the application on your local machine, use:

    ```bash
    go run .
    ```
4. **Run with Docker (optional):**

    To run the application within a Docker container, follow these steps:

    ```bash
    docker build -t receipt-processor .
    docker run -p 8080:8080 receipt-processor
    ```
