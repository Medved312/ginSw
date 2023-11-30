package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"main/mappers"
	"main/model"
	"net/http"
	"strconv"
)

// @Summary      GetModel
// @Description  Получение марки
// @Param        id path int true	"id марки"
// @Tags         Model
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.ModelView
// @Router       /Model/{id} [get]
func GetModelHandler(c *gin.Context) {

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

func GetModelById(id uint) *model.Mark {
	var mark *model.Mark
	result := database.GetDB().First(&mark, id)
	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))

		return nil
	}

	return mark

}

// @Summary      GetModel
// @Description  Получение списка марок
// @Tags         Model
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.ModelView
// @Router       /all-model [get]
func GetAllModelHandler(c *gin.Context) {
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

// @Summary      CreateModel
// @Description  Добавление марки в базу
// @Param        input body saveDTO.ModelDTO  true  "создание марки"
// @Tags         Model
// @Accept       json
// @Produce      json
// @Router       /Model/add/ [post]
func CreateModelHandler(c *gin.Context) {
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

// @Summary      DeleteModel
// @Description  Удаление марки из базы
// @Param        id path int true	"id марки"
// @Tags         Model
// @Accept       json
// @Produce      json
// @Router       /model/{id}  [delete]
func DeleteModelHandler(c *gin.Context) {
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

// @Summary      UpdateModel
// @Description  Обновление данных марки
// @Param		 id path int true	"id марки"
// @Param        input body saveDTO.ModelDTO  true  "Новые значения"
// @Tags         Model
// @Accept       json
// @Produce      json
// @Router       /model/{id}  [put]
func UpdateModelHandler(c *gin.Context) {
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
