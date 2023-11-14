package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/1shubham7/jwt/database"
	"github.com/akhil/golang-jwt-project/database"
	jwt "github.com/dgrijalva/jwt-go" // golang driver for jwt
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email string
	First_name string
	Last_name string
	Uid string
	User_type string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// btw we sould have our secret key in .env for production 
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error){
	claims := &SignedDetails{
		Email : email,
		First_name: firstName,
		Last_name: lastName,
		Uid : uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			// setting the expiry time
			ExpiresAt: time.Now().Local().Add(time.Hour *time.Duration(120)).Unix(),
		},
	}

		// refreshClaims is used to get a new token if th eprevious once is expired.
	
		refreshClaims := &SignedDetails{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Local().Add(time.Hour *time.Duration(172)).Unix(),
		},
	}
}