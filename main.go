package main

import (
    "database/sql"
	"encoding/json"
    "fmt"
    "log"
	"net/http"
    // "io/ioutil"
	"time"
    "os"
    
	_ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    dbname   = "postgres"
)

func connectDB() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, os.Getenv("POSTGRES_PASSWORD"), dbname)
    
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Successfully connected to PostgreSQL!")
    return db, nil
}

type Product struct {
    ID             int64     `json:"id"`
    Title          string    `json:"title"`
    BodyHTML       string    `json:"body_html"`
    Vendor         string    `json:"vendor"`
    ProductType    string    `json:"product_type"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
    Handle         string    `json:"handle"`
    Status         string    `json:"status"`
    Variants       []Variant `json:"variants"`
}

type Variant struct {
    Price            string `json:"price"`
    InventoryQuantity int   `json:"inventory_quantity"`
}

func getShopifyProducts() ([]Product, error) {
    shopName := os.Getenv("SHOP_NAME")
    accessToken := os.Getenv("SHOPIFY_ACCESS_TOKEN")
    url := fmt.Sprintf("https://%s.myshopify.com/admin/api/2024-07/products.json", shopName)

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Add("X-Shopify-Access-Token", accessToken)
    
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read the response body
    // body, err := ioutil.ReadAll(resp.Body)
    // if err!= nil {
    //     fmt.Println(err)
    //     return
    // }

    // Display the response body
    //fmt.Println("Body:", string(body))

    var result struct {
        Products []Product `json:"products"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result.Products, nil
}

func insertProduct(db *sql.DB, product Product) error {
    // Insert product data
    query := `
        INSERT INTO products (title, body_html, vendor, product_type, created_at, updated_at, handle, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id;
    `
    var productId int
    err := db.QueryRow(query, product.Title, product.BodyHTML, product.Vendor, product.ProductType, product.CreatedAt, product.UpdatedAt, product.Handle, product.Status).Scan(&productId)
    if err!= nil {
        return err
    }

    // Insert variants data
    query = `
        INSERT INTO variants (product_id, price, inventory_quantity)
        VALUES ($1, $2, $3);
    `
    for _, variant := range product.Variants {
        _, err := db.Exec(query, productId, variant.Price, variant.InventoryQuantity)
        if err!= nil {
            return err
        }
    }
    return nil
}

func main() {
    err := godotenv.Load()
    if err!= nil {
        log.Fatal(err)
    }

    db, err := connectDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    products, err := getShopifyProducts()
    if err != nil {
        log.Fatal(err)
    }

    for _, product := range products {
        // fmt.Printf("Product: %s\n", product.Title)
        err := insertProduct(db, product)
        if err != nil {
            log.Printf("Error inserting product %s: %v", product.Title, err)
        } else {
            fmt.Printf("Inserted product: %s\n", product.Title)
        }
    }
}
