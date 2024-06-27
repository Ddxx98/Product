# Product

# API Documentation

## FastAPI Endpoints

### Add Product

- **URL**: `/product`
- **Method**: POST
- **Authorization**: TOKEN
- **Request Body**: {
    "name":"Book",
    "price": 20,
}
- **Response**: {
    "name":"Book",
    "price": 20,
}

### Get Product

- **URL**: `/product`
- **Method**: GET
- **Authorization**: TOKEN
- **Response**: [
    {
        "name": "Book",
        "price": 30
    },
    {
        "name": "Paper",
        "price": 20
    }
]

### Get Token

- **URL**: `/token`
- **Method**: GET
- **Request Body**: {
    "name":"Book",
}
- **Response**:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJsb2dnZWRJbkFzIjoiYWRtaW4iLCJpYXQiOjE0MjI3Nzk2Mzh9gzSraSYS8EXBxLN_oWnFSRgCzcmJmMjLiuyu5CSpyHI"

