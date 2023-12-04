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

// @Summary      GetCar
// @Description  Получение автомобиля
// @Param        id path int true	"id автомобиля"
// @Tags         Car
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.CarView
// @Router       /car/{id} [get]
func GetCarHandler(c *gin.Context) {

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := GetCarById(uint(id))

	if result == nil {
		log.Println(fmt.Sprintf("автомобиль с id = %v не удалось найти", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToCarView(result),
	})
}

func GetCarById(id uint) *model.Car {

	var Car *model.Car

	result := database.GetDB().Model(&model.Car{}).
		Where("id = ?", id).Find(&Car)

	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))
		return nil
	}

	return Car
}

// @Summary      GetCarRange
// @Description  Получение списка машин
// @Param        offset path int true	"начало списка"
// @Param		 limit path int true	"конец списка"
// @Tags         Car
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.CarView
// @Router       /range-car/{offset}/{limit} [get]
func GetRangeCarHandler(c *gin.Context) {
	var cars []*model.Car

	offset, err := strconv.ParseUint(c.Params.ByName("offset"), 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	limit, err := strconv.ParseUint(c.Params.ByName("limit"), 10, 32)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result := database.GetDB().Offset(int(offset)).Limit(int(limit)).Find(&cars)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToCarViews(cars),
	})
}

// @Summary      CreateCar
// @Description  Добавление автомобиля в базу
// @Param        input body saveDTO.CarDTO  true  "создание автомобиля"
// @Tags         Car
// @Accept       json
// @Produce      json
// @Router       /car/add/ [post]
func CreateCarHandler(c *gin.Context) {
	var Car *model.Car

	if err := c.ShouldBindJSON(&Car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Create(&Car)
	if result.Error != nil {
		log.Println(result.Error, &Car)
		c.JSON(400, gin.H{
			"message": "Error adding to the database",
		})
		return
	} else {
		newCar := GetCarById(Car.ID)
		c.JSON(200, gin.H{
			"result": mappers.MapToCarView(newCar),
		})
	}

}

// @Summary      DeleteCar
// @Description  Удаление автомобиля из базы
// @Param        id path int true	"id автомобиля"
// @Tags         Car
// @Accept       json
// @Produce      json
// @Router       /car/{id}  [delete]
func DeleteCarHandler(c *gin.Context) {
	var car *model.Car

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := database.GetDB().Model(&model.Car{}).Preload("Categories").
		Where("id = ?", id).Find(&car)
	database.GetDB().Model(car).Association("Categories").Clear()
	database.GetDB().Delete(&car)

	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": "Status OK",
	})
}

// @Summary      UpdateCar
// @Description  Обновление данных автомобиля
// @Param		 id path int true	"id автомобиля"
// @Param        input body saveDTO.CarDTO  true  "Новые значения"
// @Tags         Car
// @Accept       json
// @Produce      json
// @Router       /car/{id}  [put]
func UpdateCarHandler(c *gin.Context) {
	var oldCar *model.Car = &model.Car{}
	var upCar *model.Car = &model.Car{}
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", id).First(&oldCar)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"result": result,
		})
	} else {
		oldCar = upCar
		oldCar.ID = uint(id)

		database.GetDB().Save(&oldCar)
		c.JSON(200, gin.H{
			"resulst": result,
		})
	}

}
