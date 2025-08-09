package handlers

import (
    "database/sql"
    "net/http"
    "taskmanager/internal/db"
    "taskmanager/internal/models"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }

    user.Password = string(hashedPassword)
    if err := db.InsertUser(db.DB, user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        var creds models.User
        if err := c.ShouldBindJSON(&creds); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        user, err := db.GetUserByName(db.DB, creds.Username)
        if err != nil {
            if err == sql.ErrNoRows {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
            return
        }

        expirationTime := time.Now().Add(15 * time.Minute)
        claims := &Claims{
            Username: creds.Username,
            RegisteredClaims: jwt.RegisteredClaims{
                ExpiresAt: jwt.NewNumericDate(expirationTime),
            },
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        tokenString, err := token.SignedString([]byte(jwtSecret))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": tokenString})
    }
}
