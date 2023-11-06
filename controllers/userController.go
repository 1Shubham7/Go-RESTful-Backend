package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	database "github.com/1shubham7/jwt/database"
	helper "github.com/1shubham7/jwt/helpers"
	models "github.com/1shubham7/jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func Hashpassword()

func VerifyPassword()

func SignUp() {
	
}

func Login()

func GetUsers()

func GetUserById()