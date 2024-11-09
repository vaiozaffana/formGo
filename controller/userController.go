package controller

import (
	database "formGo/Database"
	model "formGo/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderForm(c *gin.Context) {
	htmlContent := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Form Input Data</title>
	</head>
	<body>
		<h2>Form Input Data</h2>
		<form action="/submit" method="post">
			<label for="name">Nama:</label><br>
			<input type="text" id="name" name="name" required><br><br>

			<label for="email">Email:</label><br>
			<input type="email" id="email" name="email" required><br><br>

			<label for="password">Passowrd:</label><br>
			<input type="password" id="password" name="password" required><br><br>

			<button type="submit">Kirim</button>
		</form>
	</body>
	</html>
	`
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, htmlContent)
}

func SubmitForm(c *gin.Context) {
	var data model.User
	data.Name = c.PostForm("name")
	data.Email = c.PostForm("email")
	data.Password = c.PostForm("password")

	if err := database.DB.Create(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil disimpan"})
}
