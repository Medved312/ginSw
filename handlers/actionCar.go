package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"main/mappers"
	"main/model"
	"main/model/saveDTO"
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

// @Summary      CreateCar
// @Description  Добавление автомобиля в базу
// @Param        input body saveDTO.CarDTO  true  "создание автомобиля"
// @Tags         Car
// @Accept       json
// @Produce      json
// @Router       /car/add/ [post]
func CreateCarHandler(c *gin.Context) {
	var Car *model.Car
	var createdCar *saveDTO.CarDTO
	if err := c.ShouldBindJSON(&createdCar); err != nil {
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
	var car *model.Car = &model.Car{}
	var upCar *model.Car = &model.Car{}
	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", value).First(&car)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"result": result,
		})
	} else {
		if upCar.Description != "" {
			car.Description = upCar.Description
		}
		if upCar.Description != "" {
			car.Description = upCar.Description
		}
		database.GetDB().Save(&car)
		c.JSON(200, gin.H{
			"resulst": result,
		})
	}

}
