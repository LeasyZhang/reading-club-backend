package api

import (
	"net/http"
	featureService "reading-club-backend/service/feature"

	"github.com/gin-gonic/gin"
)

// EnableFeature : enable feature
func EnableFeature(c *gin.Context) {
	feature := c.Param("name")
	enabled := featureService.EnableFeature(feature)

	if enabled == false {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not enable feature!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "successfully enabled!!"})
	}
}

// DisableFeature : disable feature by name
func DisableFeature(c *gin.Context) {
	feature := c.Param("name")
	enabled := featureService.DisableFeature(feature)

	if enabled == false {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not enable feature!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "successfully disabled!!"})
	}
}

// AddFeature : add a new feature
func AddFeature(c *gin.Context) {
	feature := c.Param("name")
	_, err := featureService.AddFeature(feature)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can not add feature!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "successfully enabled!!"})
	}
}
