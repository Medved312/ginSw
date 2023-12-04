package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "gorm.io/gorm"
	"log"
	"main/database"
	_ "main/docs"
	"main/mappers"
	"main/model"
	"net/http"
	"strconv"
)

// @Summary      GetTransmission
// @Description  Получение трансмиссии
// @Param        id path int true	"id трансмиссии"
// @Tags         Transmission
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.TransmissionView
// @Router       /transmission/{id} [get]
func GetTransmissionHandler(c *gin.Context) {

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := GetTransmissionById(uint(id))
	if result == nil {
		log.Println(result, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToTransmissionView(result),
	})
}

func GetTransmissionById(id uint) *model.Transmission {
	var model *model.Transmission
	result := database.GetDB().First(&model, id)
	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))

		return nil
	}

	return model

}

// @Summary      GetTransmission
// @Description  Получение списка трансмиссий
// @Tags         Transmission
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.TransmissionView
// @Router       /all-transmission [get]
func GetAllTransmissionHandler(c *gin.Context) {
	var transmissions []*model.Transmission

	result := database.GetDB().Find(&transmissions)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToTransmissionViews(transmissions),
	})
}

// @Summary      CreateTransmission
// @Description  Добавление трансмиссии в базу
// @Param        input body saveDTO.TransmissionDTO  true  "создание трансмиссии"
// @Tags         Transmission
// @Accept       json
// @Produce      json
// @Router       /transmission/add/ [post]
func CreateTransmissionHandler(c *gin.Context) {
	var transmission *model.Transmission

	if err := c.ShouldBindJSON(&transmission); err != nil {
		log.Println(err.Error(), &transmission)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	row := database.GetDB().Where("name = ?", transmission.Name).First(&transmission)
	if row.Error != nil {
		result := database.GetDB().Create(&transmission)

		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(400, gin.H{
				"message": "Error when adding",
			})
		} else {
			newTransmission := GetTransmissionById(transmission.ID)
			c.JSON(200, gin.H{
				"result": mappers.MapToTransmissionView(newTransmission),
			})
		}
	} else {
		c.JSON(409, gin.H{
			"message": "data duplication",
		})
	}
}

// @Summary      DeleteTransmission
// @Description  Удаление модели из базы
// @Param        id path int true	"id трансмиссии"
// @Tags         Transmission
// @Accept       json
// @Produce      json
// @Router       /transmission/{id}  [delete]
func DeleteTransmissionHandler(c *gin.Context) {
	var transmission []*model.Transmission

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := database.GetDB().Where("id = ?", id).Find(&transmission)
	database.GetDB().Delete(&transmission)

	//result := database.GetDB().Model(&model.Mark{}).Preload("Product").
	//	Where("id = ?", id).Find(&mark)
	//database.GetDB().Model(mark).Association("Model").Clear()
	//database.GetDB().Delete(&mark)

	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": "status OK",
	})
}

// @Summary      UpdateTransmission
// @Description  Обновление данных трансмиссии
// @Param		 id path int true	"id трансмиссии"
// @Param        input body saveDTO.TransmissionDTO  true  "Новые значения"
// @Tags         Transmission
// @Accept       json
// @Produce      json
// @Router       /transmission/{id}  [put]
func UpdateTransmissionHandler(c *gin.Context) {
	var oldTransmission *model.Transmission = &model.Transmission{}
	var upTransmission *model.Transmission = &model.Transmission{}
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upTransmission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", id).First(&oldTransmission)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	} else {
		oldTransmission = upTransmission
		oldTransmission.ID = uint(id)

		result = database.GetDB().Save(&oldTransmission)
		c.JSON(200, gin.H{
			"result": result,
		})
	}
}
