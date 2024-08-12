# API Documentation

## Overview

This document provides details on the RESTful API for task management. The API is built using Go and Gin and provides endpoints to manage tasks.

## Base URL

```
http://localhost:8088

```

## Endpoints

### Get All Tasks

- **URL**: `/tasks`
- **Method**: `GET`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "tasks": [
        {
          "id": "1",
          "title": "Task Manager Project",
          "description": "Add/View/Delete Tasks",
          "due_date": "2024-08-07T12:00:00Z",
          "status": "In Progress"
        },
        ...
      ]
    }
    ```
- **Error Response**:
  - **Code**: 404 Not Found
  - **Content**:
    ```json
    {
      "message": "No available Tasks"
    }
    ```

### Get Task By ID

- **URL**: `/tasks/:id`
- **Method**: `GET`
- **URL Params**:
  - **Required**: `id=[string]`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "task": {
        "id": "1",
        "title": "Task Manager Project",
        "description": "Add/View/Delete Tasks",
        "due_date": "2024-08-07T12:00:00Z",
        "status": "In Progress"
      }
    }
    ```
- **Error Response**:
  - **Code**: 404 Not Found
  - **Content**:
    ```json
    {
      "message": "Task Not Found"
    }
    ```

### Add New Task

- **URL**: `/tasks`
- **Method**: `POST`
- **Body**:
  - **Content**:
    ```json
    {
      "title": "New Task Title",
      "description": "Task Description",
      "due_date": "2024-08-08T12:00:00Z",
      "status": "Pending"
    }
    ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**:
    ```json
    {
      "id": "5",
      "title": "New Task Title",
      "description": "Task Description",
      "due_date": "2024-08-08T12:00:00Z",
      "status": "Pending"
    }
    ```
- **Error Response**:
  - **Code**: 400 Bad Request
  - **Content**:
    ```json
    {
      "message": "Bad Request"
    }
    ```

### Update Task By ID

- **URL**: `/tasks/:id`
- **Method**: `PUT`
- **URL Params**:
  - **Required**: `id=[string]`
- **Body**:
  - **Content**:
    ```json
    {
      "title": "Updated Task Title",
      "description": "Updated Description",
      "due_date": "2024-08-09T12:00:00Z",
      "status": "Completed"
    }
    ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**:
    ```json
    {
      "id": "1",
      "title": "Updated Task Title",
      "description": "Updated Description",
      "due_date": "2024-08-09T12:00:00Z",
      "status": "Completed"
    }
    ```
- **Error Response**:
  - **Code**: 400 Bad Request
  - **Content**:
    ```json
    {
      "message": "Bad Request"
    }
    ```
  - **Code**: 404 Not Found
  - **Content**:
    ```json
    {
      "message": "No such task"
    }
    ```

### Delete Task By ID

- **URL**: `/tasks/:id`
- **Method**: `DELETE`
- **URL Params**:
  - **Required**: `id=[string]`
- **Success Response**:
  - **Code**: 204 No Content
  - **Content**:
    ```json
    {
      "message": "Task Removed Successfully!"
    }
    ```
- **Error Response**:
  - **Code**: 404 Not Found
  - **Content**:
    ```json
    {
      "message": "Task Not Found"
    }
    ```

## Models

### Task

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "due_date": "string",
  "status": "string"
}
```

- **id**: Unique identifier for the task.
- **title**: Title of the task.
- **description**: Description of the task.
- **due_date**: Due date of the task in ISO 8601 format.
- **status**: Current status of the task (e.g., "In Progress", "Completed").

## Utility Functions

### giveId()

- **Description**: Generates a random ID for new tasks.
- **Returns**: A string representing the generated ID.

## Error Handling

The API provides error messages in the response body with an appropriate HTTP status code.
## Postman Documentation
postman docs(https://documenter.getpostman.com/view/32287741/2sA3s4mVqM)
