package categories

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteCategory(c *gin.Context) {
	ctx := c.Request.Context()
	categoryIDStr := c.Param("categoryID")
	categoryID, err := strconv.Atoi(categoryIDStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	errCode, err := h.categorySvc.DeleteCategory(ctx, int64(categoryID))
	if err != nil {
		c.JSON(errCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully deleted category",
	})
}
