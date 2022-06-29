package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type AppForm struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Password  string `json:"password" form:"password"`
	Gender    string `json:"gender" form:"gender"`
}

func getCorsConfig() gin.HandlerFunc {
	var origins = []string{"*"}

	return cors.New(cors.Config{

		AllowOrigins: origins,
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Accept", "Accept-Encoding",
			"Accept-Language", "Connection", "Content-Length",
			"Content-Type", "Host", "Origin", "Referer", "User-Agent", "transformRequest"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

var forms = []AppForm{
	{
		FirstName: "",
		LastName:  "",
		Password:  "",
		Gender:    "",
	},
}

// func fetchResponse(url string) string {
// 	resp, _ := http.Get(url)
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	return string(body)
// }

func main() {
	// resp := fetchResponse("http://localhost:8000/person")
	// fmt.Println(resp)
	router := gin.Default()
	router.Use(getCorsConfig())
	db, err := sql.Open("mysql", "root:Prabhat@2022@tcp(127.0.0.1:3306)/form?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connection Successfully!!!")

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	// router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	router.POST("/person", func(c *gin.Context) {
		// FirstName := c.Request.FormValue("FirstName")
		// fmt.Println(FirstName)
		// LastName := c.Request.FormValue("LastName")
		// fmt.Println(LastName)
		// Password := c.Request.FormValue("Password")
		// Gender := c.Request.FormValue("Gender")
		// var userForm AppForm

		// _, err := db.Exec("INSERT INTO regForm(FirstName,LastName,Password,Gender)VALUES(?,?,?,?)", FirstName, LastName, Password, Gender)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// if err := c.Bind(&userForm); err == nil {
		// 	fmt.Printf("user form - %+v\n", userForm)
		// 	fmt.Println("insert regForm FirstName{}", FirstName)
		// 	msg := fmt.Sprintf("insert successful %s", FirstName)
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"msg":    msg,
		// 		"status": "ok",
		// 		"data":   userForm,
		// 	})
		// } else {
		// 	fmt.Printf("user form -%+v\n", err)
		// }

		// fmt.Println("insert regForm FirstName{}", FirstName)
		// msg := fmt.Sprintf("insert successful %s", FirstName)
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": msg,
		// })
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": msg,
		// 	"status": "ok",
		// 	"data":   userForm,
		// })

		// var user AppForm
		// c.Bind(&user)
		// log.Println(user)
		// if user.FirstName != "" && user.LastName != "" && user.Password != "" && user.Gender != ""{
		// 	if insert, _ := db.Exec(`INSERT INTO user (FirstName, LastName, Password, Gender) VALUES (?, ?, ?, ?)`, user.FirstName, user.LastName, user.Password, user.Gender); insert != nil{
		// 		_, err := insert.LastInsertId()
		// 	if err == nil {
		// 		content := &AppForm{
		// 			Username:  user.FirstName,
		// 			LastName:  user.LastName,
		// 			Password: user.Password,
		// 			Gender:  user.Gender,
		// 		}
		// 		c.JSON(201, content)
		// 	} else {
		// 		checkErr(err, "Insert failed")
		// 	}
		// } else {
		// 	c.JSON(400, gin.H{"error": "Fields are empty"})
		// }

		// }
		var user AppForm
		c.Bind(&user)
		log.Println(user)
		if user.FirstName != "" && user.LastName != "" && user.Password != "" && user.Gender != "" {
			if insert, _ := db.Exec(`INSERT INTO regForm(FirstName, LastName, Password, Gender) VALUES(?, ?, ?, ?)`, user.FirstName, user.LastName, user.Password, user.Gender); insert != nil {
				_, err := insert.LastInsertId()
				if err == nil {
					content := &AppForm{
						FirstName: user.FirstName,
						LastName:  user.LastName,
						Password:  user.Password,
						Gender:    user.Gender,
					}
					c.JSON(http.StatusOK, gin.H{
						"status": "ok",
						"data":   content,
					})
				}
			}
		}

	})
	// router.POST("/person", func(c *gin.Context) {
	// 	var newform AppForm
	// 	if err := c.BindJSON(&newform); err != nil {
	// 		return
	// 	}
	// 	forms = append(forms, newform)
	// 	_, err := db.Exec("INSERT INTO regform(appforms) VALUES(?,?,?,?)", newform)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	c.IndentedJSON(http.StatusCreated, newform)

	// })

	router.GET("/persons", func(c *gin.Context) {
		rows, err := db.Query("SELECT FirstName,LastName,Password,Gender FROM regForm")
		defer rows.Close()

		if err != nil {
			log.Fatalln(err)
		}

		appforms := make([]AppForm, 0)

		for rows.Next() {
			var appform AppForm
			rows.Scan(&appform.FirstName, &appform.LastName, &appform.Password, &appform.Gender)
			appforms = append(appforms, appform)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"appforms": appforms,
		})
	})
	router.Run(":8000")
}
