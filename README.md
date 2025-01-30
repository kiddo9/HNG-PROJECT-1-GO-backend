# HNG-PROJECT-1-GO-Backend

## Description
This is an internship project for the backend track on HNG. The project retrieves intern details, including:
- Email
- Current date (dynamically generated in **ISO 8601 format (UTC)**)
- Project GitHub URL

---

## Setup Instructions

### Prerequisites
Ensure you have the following installed:
- **Go** (`1.18` or later) → [Download Go](https://go.dev/dl/)
- **Git** → [Download Git](https://git-scm.com/downloads)
- **VS Code** (Recommended) → [Download VS Code](https://code.visualstudio.com/)

### Installation & Running the Project

1. **Clone the repository**:
   ```sh
   git clone https://github.com/kiddo9/HNG-PROJECT-1-GO-backend.git
2. **Navigate to the project directory**:
    ```sh
    cd HNG-PROJECT-1-GO-backend

3. **Start the server locally**
    ```sh
    go run main.go
4. **Test the API**:
    *Open Postman or run the following cURL command*:
    ```sh
    curl -X GET http://localhost:8080


***API Documentation***:
1. Local URL:
    ```
    http://localhost:8080

2. Online URL:
    ```
    https://hng-project-1-go-backend-production.up.railway.app/

3. Get Intern Details
    
    Endpoint: 
    ```
    GET '/'


 4.  Response Format:
        ```
        {
            "email": "your-email@example.com",
            "current_datetime": "2025-01-30T09:30:00Z",
            "github_url": "https://github.com/yourusername/your-repo"
        }

5. Example Usage (cURL):
    ```
        curl -X GET http://localhost:8080

Folder Structure

    HNG-PROJECT-1-GO-backend/
    │── main.go
    │── go.mod
    │── go.sum
    │── README.md

Programming Language/Stack backlink:
```sh
https://hng.tech/hire/golang-developers

