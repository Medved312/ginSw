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

// @Summary      GetMark
// @Description  Получение марки
// @Param        id path int true	"id марки"
// @Tags         Mark
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.MarkView
// @Router       /mark/{id} [get]
func GetMarkHandler(c *gin.Context) {

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := GetMarkById(uint(id))
	if result == nil {
		log.Println(result, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToMarkView(result),
	})
}

func GetMarkById(id uint) *model.Mark {
	var mark *model.Mark
	result := database.GetDB().First(&mark, id)
	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))

		return nil
	}

	return mark

}

// @Summary      GetMark
// @Description  Получение списка марок
// @Tags         Mark
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.MarkView
// @Router       /all-mark [get]
func GetAllMarkHandler(c *gin.Context) {
	var marks []*model.Mark

	result := database.GetDB().Find(&marks)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToMarkViews(marks),
	})
}

// @Summary      CreateMark
// @Description  Добавление марки в базу
// @Param        input body saveDTO.MarkDTO  true  "создание марки"
// @Tags         Mark
// @Accept       json
// @Produce      json
// @Router       /mark/add/ [post]
func CreateMarkHandler(c *gin.Context) {
	var mark *model.Mark

	if err := c.ShouldBindJSON(&mark); err != nil {
		log.Println(err.Error(), &mark)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	row := database.GetDB().Where("name = ?", mark.Name).First(&mark)
	if row.Error != nil {
		result := database.GetDB().Create(&mark)

		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(400, gin.H{
				"message": "Error when adding",
			})
		} else {
			newMark := GetMarkById(mark.ID)
			c.JSON(200, gin.H{
				"result": mappers.MapToMarkView(newMark),
			})
		}
	} else {
		c.JSON(409, gin.H{
			"message": "data duplication",
		})
	}
}

// @Summary      DeleteMark
// @Description  Удаление марки из базы
// @Param        id path int true	"id марки"
// @Tags         Mark
// @Accept       json
// @Produce      json
// @Router       /mark/{id}  [delete]
func DeleteMarkHandler(c *gin.Context) {
	var mark []*model.Mark

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := database.GetDB().Where("id = ?", id).Find(&mark)
	database.GetDB().Delete(&mark)

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

// @Summary      UpdateMark
// @Description  Обновление данных марки
// @Param		 id path int true	"id марки"
// @Param        input body saveDTO.MarkDTO  true  "Новые значения"
// @Tags         Mark
// @Accept       json
// @Produce      json
// @Router       /mark/{id}  [put]
func UpdateMarkHandler(c *gin.Context) {
	var mark *model.Mark = &model.Mark{}
	var upMark *model.Mark = &model.Mark{}
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upMark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", id).First(&mark)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	} else {
		mark = upMark
		mark.ID = uint(id)

		result = database.GetDB().Save(&mark)
		c.JSON(200, gin.H{
			"result": result,
		})
	}
}
