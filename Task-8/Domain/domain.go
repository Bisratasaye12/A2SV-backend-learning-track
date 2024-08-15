package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"due_date" json:"due_date"`
	Status      string             `bson:"status" json:"status"`
}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username" validate:"required,min=3,max=30"`
	Email     string             `json:"email" bson:"email" validate:"email"`
	Password  string             `json:"password" bson:"password" validate:"required,min=8"`
	Role      string             `json:"role" bson:"role"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}


type TaskRepository interface {
    GetAllTasks(ctx context.Context) ([]Task, error)
    GetTaskByID(ctx context.Context, id primitive.ObjectID) (Task, error)
    AddTask(ctx context.Context, task *Task) (Task, error)
    UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask *Task) (Task, error)
    DeleteTask(ctx context.Context, id primitive.ObjectID) error
}


type TaskUseCase interface {
    GetAllTasks(ctx context.Context) ([]Task, error)
    GetTaskByID(ctx context.Context, id primitive.ObjectID) (Task, error)
    AddTask(ctx context.Context, task *Task) (Task, error)
    UpdateTask(ctx context.Context, id primitive.ObjectID, updatedTask *Task) (Task, error)
    DeleteTask(ctx context.Context, id primitive.ObjectID) error
}

type TaskController interface {
	GetAllTasks(c *gin.Context)
	GetTaskByID(c *gin.Context)
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type UserRepository interface {
	NoUsers(ctx context.Context) (bool, error)
    Register(ctx context.Context, user *User) (User, error)
    Login(ctx context.Context, username string) (*User, error)
	PromoteUser(ctx context.Context, id primitive.ObjectID) error
}

type UserUseCase interface {
    Register(ctx context.Context, user *User) (User, error)
    Login(ctx context.Context, user *User) (string, error)
    PromoteUser(ctx context.Context, id primitive.ObjectID) error
}

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	PromoteUser(c *gin.Context)
}

type Infrastructure interface {
	AuthMiddleware(requiredRoles ...string) gin.HandlerFunc
	EncryptPassword(password string) (string, error)
	JWT_Auth(existingUser *User, user *User) (string, error)
}