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
	"main/model/saveDTO"
	"net/http"
	"strconv"
)

// @Summary      GetProduct
// @Description  Получение продукта
// @Param        id path int true	"id продукта"
// @Tags         Product
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.ProductView
// @Router       /product/{id} [get]
func GetProductHandler(c *gin.Context) {

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := GetProductById(uint(id))

	if result == nil {
		log.Println(fmt.Sprintf("фильм с id = %v не удалось найти", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToProductView(result),
	})
}

func GetProductById(id uint) *model.Product {

	var Product *model.Product

	result := database.GetDB().Model(&model.Product{}).Preload("Categories").
		Where("id = ?", id).Find(&Product)

	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))
		return nil
	}

	return Product
}

// @Summary      GetCategory
// @Description  Получение жанра
// @Param        id path int true	"id category"
// @Tags         Category
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.CategoryView
// @Router       /category/{id} [get]
func GetCategoryHandler(c *gin.Context) {

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := GetCategoryById(uint(id))
	if result == nil {
		log.Println(result, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToCategoryView(result),
	})
}

func GetCategoryById(id uint) *model.Category {
	var category *model.Category
	result := database.GetDB().First(&category, id)
	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))

		return nil
	}

	return category

}

// @Summary      GetCategory
// @Description  Получение списка категорий
// @Tags         Category
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.CategoryView
// @Router       /all-categories [get]
func GetAllCategoryHandler(c *gin.Context) {
	var category []*model.Category

	result := database.GetDB().Find(&category)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToCategoryViews(category),
	})
}

// @Summary      CreateProduct
// @Description  Добавление продукта в базу
// @Param        input body saveDTO.ProductDTO  true  "создание продукта"
// @Tags         Product
// @Accept       json
// @Produce      json
// @Router       /product/add/ [post]
func CreateProductHandler(c *gin.Context) {
	var Product *model.Product
	var createdProduct *saveDTO.ProductDTO
	if err := c.ShouldBindJSON(&createdProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var linkedGenres []*model.Category
	for _, categoryID := range createdProduct.Categories {
		linkedGenres = append(linkedGenres, &model.Category{Id: categoryID})
	}
	Product = &model.Product{
		Name:        createdProduct.Name,
		Description: createdProduct.Description,
		Categories:  linkedGenres}

	result := database.GetDB().Omit("category").Create(&Product)
	if result.Error != nil {
		log.Println(result.Error, &Product)
		c.JSON(400, gin.H{
			"message": "Error adding to the database",
		})
		return
	} else {
		newProduct := GetProductById(Product.Id)
		c.JSON(200, gin.H{
			"result": mappers.MapToProductView(newProduct),
		})
	}

}

// @Summary      CreateCategory
// @Description  Добавление категории в базу
// @Param        input body saveDTO.CategoryDTO  true  "создание категории"
// @Tags         Category
// @Accept       json
// @Produce      json
// @Router       /category/add/ [post]
func CreateCategoryHandler(c *gin.Context) {
	var category *model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		log.Println(err.Error(), &category)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	row := database.GetDB().Where("name = ?", category.Name).First(&category)
	if row.Error != nil {
		result := database.GetDB().Create(&category)

		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(400, gin.H{
				"message": "Error when adding",
			})
		} else {
			newGenre := GetCategoryById(category.Id)
			c.JSON(200, gin.H{
				"result": mappers.MapToCategoryView(newGenre),
			})
		}
	} else {
		c.JSON(409, gin.H{
			"message": "data duplication",
		})
	}
}

// @Summary      DeleteProduct
// @Description  Удаление продукта из базы
// @Param        id path int true	"id продукта"
// @Tags         Product
// @Accept       json
// @Produce      json
// @Router       /product/{id}  [delete]
func DeleteProductHandler(c *gin.Context) {
	var product *model.Product

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := database.GetDB().Model(&model.Product{}).Preload("Categories").
		Where("id = ?", id).Find(&product)
	database.GetDB().Model(product).Association("Categories").Clear()
	database.GetDB().Delete(&product)

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

// @Summary      DeleteCategory
// @Description  Удаление категории из базы
// @Param        id path int true	"id категории"
// @Tags         Category
// @Accept       json
// @Produce      json
// @Router       /category/{id}  [delete]
func DeleteGenreHandler(c *gin.Context) {
	var Category []*model.Category

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := database.GetDB().Model(&model.Category{}).Preload("Product").
		Where("id = ?", id).Find(&Category)
	database.GetDB().Model(Category).Association("Product").Clear()
	database.GetDB().Delete(&Category)

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

// @Summary      UpdateProduct
// @Description  Обновление данных продукта
// @Param		 id path int true	"id продукта"
// @Param        input body saveDTO.ProductDTO  true  "Новые значения"
// @Tags         Product
// @Accept       json
// @Produce      json
// @Router       /product/{id}  [put]
func UpdateProductHandler(c *gin.Context) {
	var product *model.Product = &model.Product{}
	var upProduct *model.Product = &model.Product{}
	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", value).First(&product)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"result": result,
		})
	} else {
		if upProduct.Name != "" {
			product.Name = upProduct.Name
		}
		if upProduct.Description != "" {
			product.Description = upProduct.Description
		}
		database.GetDB().Save(&product)
		c.JSON(200, gin.H{
			"resulst": result,
		})
	}

}

// @Summary      UpdateCategory
// @Description  Обновление данных категории
// @Param		 id path int true	"id категории"
// @Param        input body saveDTO.CategoryDTO  true  "Новые значения"
// @Tags         Category
// @Accept       json
// @Produce      json
// @Router       /category/{id}  [put]
func UpdateCategoryHandler(c *gin.Context) {
	var category *model.Category = &model.Category{}
	var upCategory *model.Category = &model.Category{}
	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", value).First(&category)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	} else {
		if upCategory.Name != "" {
			category.Name = upCategory.Name
		}

		database.GetDB().Save(&category)
		c.JSON(200, gin.H{
			"result": result,
		})
	}
}
