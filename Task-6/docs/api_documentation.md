# Task Management API Documentation

## Base URL

```
http://localhost:8080/api
```

## Task Endpoints

### 1. **Get All Tasks**

- **Method**: `GET`
- **URL**: `/tasks`
- **Description**: Retrieves a list of all tasks from the database.
- **Request Parameters**: None
- **Response**:
  - **Status Code**: `200 OK`
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "tasks": [
        {
          "id": "ObjectId",
          "title": "Task Title",
          "description": "Task Description",
          "due_date": "2024-08-09T00:37:00Z",
          "status": "completed"
        },
        ...
      ]
    }
    ```

### 2. **Get Task by ID**

- **Method**: `GET`
- **URL**: `/tasks/:id`
- **Description**: Retrieves a single task by its ID.
- **Request Parameters**:
  - **Path**: `id` (string) - The ID of the task to retrieve.
- **Response**:
  - **Status Code**: `200 OK`
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "task": {
        "id": "ObjectId",
        "title": "Task Title",
        "description": "Task Description",
        "due_date": "2024-08-09T00:37:00Z",
        "status": "completed"
      }
    }
    ```
  - **Status Code**: `400 Bad Request`
    - **Body**:
      ```json
      {
        "message": "Invalid Task ID"
      }
      ```
  - **Status Code**: `404 Not Found`
    - **Body**:
      ```json
      {
        "message": "Task with ID <id> not found"
      }
      ```

### 3. **Add a New Task**

- **Method**: `POST`
- **URL**: `/tasks`
- **Description**: Adds a new task to the database. Requires admin role.
- **Request Body**:
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2024-08-09T00:37:00Z",
      "status": "completed"
    }
    ```
- **Response**:
  - **Status Code**: `201 Created`
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "id": "ObjectId",
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2024-08-09T00:37:00Z",
      "status": "completed"
    }
    ```
  - **Status Code**: `400 Bad Request`
    - **Body**:
      ```json
      {
        "message": "Invalid input data"
      }
      ```

### 4. **Update an Existing Task**

- **Method**: `PUT`
- **URL**: `/tasks/:id`
- **Description**: Updates an existing task in the database. Requires admin role.
- **Request Parameters**:
  - **Path**: `id` (string) - The ID of the task to update.
- **Request Body**:
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "due_date": "2024-08-09T00:37:00Z",
      "status": "in-progress"
    }
    ```
- **Response**:
  - **Status Code**: `200 OK`
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "id": "ObjectId",
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "due_date": "2024-08-09T00:37:00Z",
      "status": "in-progress"
    }
    ```
  - **Status Code**: `400 Bad Request`
    - **Body**:
      ```json
      {
        "message": "Invalid Task ID or Bad Request"
      }
      ```

### 5. **Delete a Task**

- **Method**: `DELETE`
- **URL**: `/tasks/:id`
- **Description**: Deletes a task from the database by its ID. Requires admin role.
- **Request Parameters**:
  - **Path**: `id` (string) - The ID of the task to delete.
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    {
      "message": "Task Removed Successfully!"
    }
    ```
  - **Status Code**: `400 Bad Request`
    - **Body**:
      ```json
      {
        "message": "Invalid Task ID"
      }
      ```
  - **Status Code**: `404 Not Found`
    - **Body**:
      ```json
      {
        "message": "Task with ID <id> not found"
      }
      ```

## User Endpoints

### 1. **Register User**

- **Method**: `POST`
- **URL**: `/register`
- **Description**: Registers a new user in the system.
- **Request Body**:
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "username": "Username",
      "email": "user@example.com",
      "password": "password123"
    }
    ```
- **Response**:
  - **Status Code**: `201 Created`
  - **Body**:
    ```json
    {
      "message": "User successfully registered!"
    }
    ```
  - **Status Code**: `400 Bad Request`
    - **Body**:
      ```json
      {
        "message": "Validation error message"
      }
      ```

### 2. **Login User**

- **Method**: `POST`
- **URL**: `/login`
- **Description**: Authenticates a user and returns a JWT token.
- **Request Body**:
  - **Content-Type**: `application/json`
  - **Body**:
    ```json
    {
      "username": "Username",
      "password": "password123"
    }
    ```
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    {
      "message": "User logged in successfully",
      "token": "JWT token"
    }
    ```
  - **Status Code**: `401 Unauthorized`
    - **Body**:
      ```json
      {
        "error": "Invalid username or password"
      }
      ```

### 3. **Get All Users**

- **Method**: `GET`
- **URL**: `/users`
- **Description**: Retrieves a list of all registered users. Requires user or admin role.
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    {
      "message": "List of users"
    }
    ```
  - **Status Code**: `404 Not Found`
    - **Body**:
      ```json
      {
        "message": "No registered user found."
      }
      ```

### 4. **Promote User to Admin**

- **Method**: `PUT`
- **URL**: `/users/promote/:id`
- **Description**: Promotes a user to admin role. Requires admin role.
- **Request Parameters**:
  - **Path**: `id` (string) - The ID of the user to promote.
- **Response**:
  - **Status Code**: `200 OK`
  - **Body**:
    ```json
    {
      "message": "User promoted to admin"
    }
    ```
  - **Status Code**: `403 Forbidden`
    - **Body**:
      ```json
      {
        "error": "Only admins can promote users"
      }
      ```
  - **Status Code**: `400 Bad Request`
    - **Body**:
      ```json
      {
        "error": "Invalid user ID"
      }
      ```

## Middleware

- **AuthMiddleware**: Ensures that the request contains a valid JWT token and checks user roles for access control.

## Environment Variables

- **MONGODB_URI**: MongoDB connection string.
- **GO111MODULE**: Go module setting.
- **JWT_SECRET**: Secret key for signing JWT tokens.

## Postman API Documentation

[Postman API Documentation](https://documenter.getpostman.com/view/32287741/2sA3s3HB9Y))
