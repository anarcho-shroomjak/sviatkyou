package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ContactsHandler struct {
	store ContactsStore
}

func NewContactsHandler(s ContactsStore) *ContactsHandler {
	return &ContactsHandler{
		store: s,
	}
}

//func (h *ContactsHandler) Register(router *gin.RouterGroup) {}

func (h *ContactsHandler) Get(c *gin.Context) {
	contacts, err := h.store.GetAll()
	if err != nil {
	}
	c.JSON(http.StatusOK, contacts)
}

func (h *ContactsHandler) Create(c *gin.Context) {
	var contact Contact
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(contact.Email)
	if err := h.store.Add(contact); err != nil {
	}
	c.JSON(http.StatusCreated, contact)
}

func main() {
	router := gin.Default()

	s := NewStore()
	contactsHandler := NewContactsHandler(s)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	gr := router.Group("/contacts")
	gr.GET("/", contactsHandler.Get)
	gr.POST("/", contactsHandler.Create)

	router.Run(":1488")
}
