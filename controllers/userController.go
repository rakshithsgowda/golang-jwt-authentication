package controllers

import (
	"context"
	"fmt"
	"golang-jwt-project/helpers"
	helper "golang-jwt-project/helpers"
	"golang-jwt-project/models"
	"jwt-authentication/database"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection=database.OpenCollection(database.Client,"user")
var validate = validator.New()
func HashPassword()
func VerifyPassword()
func Signup()
func Login()
func GetUsers()
func GetUser()