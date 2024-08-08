
# Task Management API Documentation

## Base URL

```
http://localhost:8080/api
```

## Endpoints

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
- **Description**: Adds a new task to the database.
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
- **Description**: Updates an existing task in the database.
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
  - **Status Code**: `404 Not Found`
    - **Body**:
      ```json
      {
        "message": "Task with ID <id> not found"
      }
      ```

### 5. **Delete a Task**

- **Method**: `DELETE`
- **URL**: `/tasks/:id`
- **Description**: Deletes a task from the database by its ID.
- **Request Parameters**:
  - **Path**: `id` (string) - The ID of the task to delete.
- **Response**:
  - **Status Code**: `204 No Content`
  - **Content-Type**: `application/json`
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