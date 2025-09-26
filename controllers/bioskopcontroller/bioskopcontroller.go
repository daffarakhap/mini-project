package bioskopcontroller

import (
	"mini-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /bioskops
func Index(c *gin.Context) {
	var bioskops []models.Bioskop
	models.DB.Find(&bioskops)
	c.JSON(http.StatusOK, gin.H{"bioskops": bioskops})
}

// GET /bioskops/:id
func Show(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	if err := models.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bioskop": bioskop})
}

// POST /bioskops
func Create(c *gin.Context) {
	var input models.Bioskop

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi manual
	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	// Simpan bioskop baru
	bioskop := models.Bioskop{
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}
	models.DB.Create(&bioskop)

	c.JSON(http.StatusOK, gin.H{"bioskop": bioskop})
}


// PUT /bioskops/:id
func Update(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	// Cari bioskop
	if err := models.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
		return
	}

	// Bind data baru
	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update field
	models.DB.Model(&bioskop).Updates(models.Bioskop{
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	})

	c.JSON(http.StatusOK, gin.H{"bioskop": bioskop})
}

// DELETE /bioskops/:id
func Delete(c *gin.Context) {
	var bioskop models.Bioskop
	id := c.Param("id")

	// Cari bioskop
	if err := models.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
		return
	}

	// Hapus bioskop
	models.DB.Delete(&bioskop)

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop deleted"})
}
