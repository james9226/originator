package resumer

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lendotopia.com/originator/services/firestore"
)

func LoadQuoteByApplicationID(c *gin.Context) {
	IDString := c.Param("id")

	applicationID, err := uuid.Parse(IDString)
	if err != nil {
		log.Fatal("Invalid Application ID Format")
		c.AbortWithStatus(400)
		return
	}

	client := firestore.Client

	client.Collection("quotes")

	c.IndentedJSON(http.StatusOK, gin.H{"Application ID": applicationID})

}
