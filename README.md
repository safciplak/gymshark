# Pack Size Calculator API

This project is a solution for calculating optimal pack sizes for product orders. The application provides a REST API built with Go that determines the most efficient combination of packs to fulfill customer orders.

## Problem Description

Given a set of standard pack sizes (250, 500, 1000, 2000, 5000 items), the application calculates the optimal combination of packs for any order quantity following these rules:

1. Only whole packs can be sent. Packs cannot be broken open.
2. Send out the least amount of items to fulfill the order (primary priority).
3. Send out as few packs as possible to fulfill each order (secondary priority).

### Example Scenarios

| Items Ordered | Correct Packs         | Why This is Optimal                    |
|--------------|----------------------|---------------------------------------|
| 1            | 1 x 250             | Smallest possible pack size           |
| 250          | 1 x 250             | Exact match                          |
| 251          | 1 x 500             | Minimum items above order            |
| 501          | 1 x 500, 1 x 250    | Minimum combination above order      |
| 12001        | 2 x 5000, 1 x 2000, 1 x 250 | Optimal combination         |

## Features

- RESTful API endpoints for pack calculation
- Flexible configuration for pack sizes
- Web UI for easy interaction
- Comprehensive unit tests
- Docker support

## Technology Stack

- Backend: Go

## Getting Started

### Prerequisites

- Go 1.x
- [Other prerequisites]

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/safciplak/gymshark.git
   ```

   ```bash
   cd gymshark
   ```

2. Build the Docker image:
   ```bash
   docker build -t gymshark:latest .
   ```

3. Run the container:
   ```bash
   docker run -d \
     --name gymshark \
     -p 8081:8081 \
     --restart unless-stopped \
     gymshark:latest
   ```

4. Verify the application is running:
   ```bash
   http://localhost:8081
   ```