package controllers

import (
	"api-crypto/crypto"
	"api-crypto/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckAccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data" : "the service is available"})
}

func Crypto(c *gin.Context) {
	var data models.CryptoData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data" : err.Error()})
		return
	}

	if data.DecData != "" {
		dataSlice := strings.Split(data.DecData, "")
		decrypted, err := crypto.Decrypt(dataSlice)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data" : err.Error()})
			return
		}
		decryptedString := strings.Join(decrypted, "")
		c.JSON(http.StatusOK, gin.H{"decrypted" : decryptedString})
		return
	}
	if data.EncData != "" {
		dataSlice := strings.Split(data.EncData, "")
		encrypted, err := crypto.Encrypt(dataSlice)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data" : err.Error()})
			return
		}
		encryptedString := strings.Join(encrypted, "")
		c.JSON(http.StatusOK, gin.H{"encrypted" : encryptedString})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"data" : "wrong request"})
}
