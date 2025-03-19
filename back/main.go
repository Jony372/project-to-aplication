package main

import (
	// "fmt"
	"context"
	"myproject/api"
	"myproject/database"

	// "io/ioutil"
	"log"
	"net/http"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connet()

	// 🔹 Servir archivos estáticos de Vue (JS, CSS, imágenes)
	r.Static("/static", "./dist/assets") 

	// 🔹 Servir el archivo index.html como la página principal
	r.LoadHTMLFiles("dist/index.html")

	//* 🔹 Ruta principal para cargar Vue
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//* 🔹 Ruta para la API
	r.GET("/api", func(c *gin.Context) {
		api.GetItems(c)
	})

	r.GET("/api/grp/:id", func(c *gin.Context) {
		var porcent float64

		err = crdbpgx.ExecuteTx(context.Background(), database.DB, pgx.TxOptions{}, func(tx pgx.Tx) (error) {
			p, err := database.GrowthPorcentage(context.Background(), tx, c.Param("id"))
			if err != nil {
				return err
			}
			porcent = p
			return nil
		})
		if err != nil {
			println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"porcent": porcent})
	})

	r.GET("/api/grp", func(c *gin.Context) {
		var items []database.ChartItem

		err = crdbpgx.ExecuteTx(context.Background(), database.DB, pgx.TxOptions{}, func(tx pgx.Tx) (error) {
			p, err := database.GetGrowthItems(context.Background(), tx)
			if err != nil {
				return err
			}
			items = p
			return nil
		})
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	})

	r.GET("/api/to-buy", func(c *gin.Context) {
		var items []database.ChartItem

		err = crdbpgx.ExecuteTx(context.Background(), database.DB, pgx.TxOptions{}, func(tx pgx.Tx) (error) {
			p, err := database.GetToBuyItems(context.Background(), tx)
			if err != nil {
				return err
			}
			items = p
			return nil
		})
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	})

	r.GET("/api/dcp", func(c *gin.Context) {
		var items []database.ChartItem

		err = crdbpgx.ExecuteTx(context.Background(), database.DB, pgx.TxOptions{}, func(tx pgx.Tx) (error) {
			p, err := database.GetDecreaseItems(context.Background(), tx)
			if err != nil {
				return err
			}
			items = p
			return nil
		})
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	})

	// 🔹 Si un usuario entra a una ruta desconocida (ej. "/dashboard"), cargar index.html
	// r.NoRoute(func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/index/")
	// 	// c.HTML(http.StatusOK, "index.html", nil)
	// })

	r.Run(":8080") // Ejecutar servidor en el puerto 8080
}
