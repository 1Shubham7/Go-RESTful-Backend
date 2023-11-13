package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/akhil/golang-jwt-project/database"
	jwt "github.com/dgrijalva/jwt-go"  // golang driver for jwt
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GenerateAllTokens(){

}