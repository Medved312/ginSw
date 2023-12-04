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

// @Summary      GetModel
// @Description  Получение модели
// @Param        id path int true	"id модели"
// @Tags         Model
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.ModelView
// @Router       /model/{id} [get]
func GetModelHandler(c *gin.Context) {

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := GetModelById(uint(id))
	if result == nil {
		log.Println(result, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToModelView(result),
	})
}

func GetModelById(id uint) *model.Model {
	var model *model.Model
	result := database.GetDB().First(&model, id)
	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))

		return nil
	}

	return model

}

// @Summary      GetAllModel
// @Description  Получение списка марок
// @Tags         Model
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.ModelView
// @Router       /all-model [get]
func GetAllModelHandler(c *gin.Context) {
	var models []*model.Model

	result := database.GetDB().Find(&models)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToModelViews(models),
	})
}

// @Summary      CreateModel
// @Description  Добавление марки в базу
// @Param        input body saveDTO.ModelDTO  true  "создание марки"
// @Tags         Model
// @Accept       json
// @Produce      json
// @Router       /model/add/ [post]
func CreateModelHandler(c *gin.Context) {
	var model *model.Model

	if err := c.ShouldBindJSON(&model); err != nil {
		log.Println(err.Error(), &model)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	row := database.GetDB().Where("name = ?", model.Name).First(&model)
	if row.Error != nil {
		result := database.GetDB().Create(&model)

		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(400, gin.H{
				"message": "Error when adding",
			})
		} else {
			newModel := GetModelById(model.ID)
			c.JSON(200, gin.H{
				"result": mappers.MapToModelView(newModel),
			})
		}
	} else {
		c.JSON(409, gin.H{
			"message": "data duplication",
		})
	}
}

// @Summary      DeleteModel
// @Description  Удаление модели из базы
// @Param        id path int true	"id модели"
// @Tags         Model
// @Accept       json
// @Produce      json
// @Router       /model/{id}  [delete]
func DeleteModelHandler(c *gin.Context) {
	var model []*model.Model

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := database.GetDB().Where("id = ?", id).Find(&model)
	database.GetDB().Delete(&model)

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
// @Description  Обновление данных модели
// @Param		 id path int true	"id модели"
// @Param        input body saveDTO.ModelDTO  true  "Новые значения"
// @Tags         Model
// @Accept       json
// @Produce      json
// @Router       /model/{id}  [put]
func UpdateModelHandler(c *gin.Context) {
	var oldModel *model.Model = &model.Model{}
	var upModel *model.Model = &model.Model{}
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", id).First(&oldModel)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	} else {
		oldModel = upModel
		oldModel.ID = uint(id)

		result = database.GetDB().Save(&oldModel)
		c.JSON(200, gin.H{
			"result": result,
		})
	}
}
