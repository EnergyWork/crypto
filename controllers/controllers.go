package controllers

import (
	"api-crypto/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data" : "service is available"})
}

func Crypto(c *gin.Context) {

	var data models.DecData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data" : err.Error()})
	}

	//decrypted := strings.Join(crypto.Decrypt(strings.Split(data.DecryptString, "")), "")

	c.JSON(http.StatusOK, gin.H{"decrypted" : "decrypted"})
}
