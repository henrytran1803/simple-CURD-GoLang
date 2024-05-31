package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"column:id"`
	RoleId    int       `json:"role_id" gorm:"column:role_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"column:email"`
	Address   string    `json:"addr" gorm:"column:addr"`
	Status    int       `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
type Role struct {
	Id        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Status    int       `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
type Account struct {
	Id        int       `json:"id" gorm:"column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	Password  string    `json:"password" gorm:"column:pass"`
	Status    int       `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (User) TableName() string    { return "users" }
func (Role) TableName() string    { return "roles" }
func (Account) TableName() string { return "account" }
func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{}, &Role{}, &Account{})
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.GET("/user/:id", func(c *gin.Context) {
		var user User
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db.Table(User{}.TableName()).Where("id = ?", id).First(&user)
		c.JSON(http.StatusOK, user)

	})
	v1.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db.Create(&user)
		c.JSON(http.StatusOK, user)
	})
	v1.DELETE("/user/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		db.Delete(&User{}, id)
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})
	v1.PUT("/user/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		var input User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user User
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		user.Name = input.Name
		user.Email = input.Email
		user.RoleId = input.RoleId
		user.Status = input.Status
		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, user)
	})
	//v1.PATCH("/user/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	}
	//	var input User
	//	if err := c.ShouldBindJSON(&input); err != nil {
	//
	//	}
	//})
	r.Run()
}
