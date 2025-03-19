package api

import (
	"context"
	"encoding/json"
	"io"
	"myproject/database"
	"net/http"
	"net/url"
	"os"

	// "github.com/cockroachdb/cockroach-go/crdb/crdbpgx"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)


func GetItems(c *gin.Context) {
	page := c.Query("page")

	parsedUrl, err := url.Parse(os.Getenv("API_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := parsedUrl.Query()
	query.Set("next_page", page)
	parsedUrl.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", parsedUrl.String(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Authorization", os.Getenv("API_KEY"))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData struct{
		Items []database.Item `json:"items"`
		NextPage string `json:"next_page"`
	}

	// Deserializar el body en una variable de tipo interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responseData)
	// Enviar la respuesta como JSON
	for i := range responseData.Items {
		err = crdbpgx.ExecuteTx(context.Background(), database.DB, pgx.TxOptions{}, func(tx pgx.Tx) error {
			return database.InsertRows(context.Background(), tx, responseData.Items[i])
    })
		if err != nil {
			println(err.Error())
			return
		}
	}
}