package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"se_cli/database"
	"se_cli/model"
	"se_cli/transformer"

	"github.com/spf13/cobra"
)

const url = "https://tiki.vn/api/v2/products?page=1&q=%20Th%E1%BB%9Di%20Trang%20Nam&limit=47"


var crawlProductCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl products from tiki api",
	Long:  `Series of interesing commands are waiting for you`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// Connect to DB
		database.Connect(getDBConfig(cmd))

		// Crawl products
		crawlProducts()
	},
}

func init() {
	crawlProductCmd.Flags().String("user", "root", "user")
	crawlProductCmd.Flags().String("password", "123456", "password")
	crawlProductCmd.Flags().String("ip", "localhost", "user")
	crawlProductCmd.Flags().String("port", "3306", "user")
	crawlProductCmd.Flags().String("db", "ecommerce", "database name")
}


func getDBConfig(cmd *cobra.Command) database.Config {
	user, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	ip, _ := cmd.Flags().GetString("ip")
	port, _ := cmd.Flags().GetString("port")
	db, _ := cmd.Flags().GetString("db")
	return database.Config{
		User: user,
		Password: password,
		Ip: ip,
		Port: port,
		Database: db,
	}
}

func crawlProducts() {
	log.Println("Start crawling products")

	// Making http request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error when making http request to crawl", "err", err)
		return
	}
	defer resp.Body.Close()

	data := model.Response{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal("error decoding", err)
		return
	}

	// transform data
	products  := transformer.TransformDataToSeProducts(data) 
	database.MysqlDB.CreateInBatches(products, 100)


	log.Println("Products have been crawled successfully!!!")
}