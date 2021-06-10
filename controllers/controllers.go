package controllers

import (
	"api-crypto/crypto"
	"api-crypto/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func CheckAccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data" : "service is available"})
}

func Decrypt(c *gin.Context) {
	var data models.DecData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data" : err.Error()})
	}
	log.Println(data)

	decrypted := strings.Join(crypto.Decrypt(strings.Split(data.DecryptString, "")), "")

	c.JSON(http.StatusOK, gin.H{"decrypted" : decrypted})
}
