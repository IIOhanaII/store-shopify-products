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

func connectDB() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("POSTGRES_HOST"), 
		os.Getenv("POSTGRES_PORT"), 
		os.Getenv("POSTGRES_USER"), 
		os.Getenv("POSTGRES_PASSWORD"), 
		os.Getenv("POSTGRES_DBNAME"))
    
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
    ID               int64  `json:"id"`
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
    //     return nil, err
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
        INSERT INTO products (id, title, body_html, vendor, product_type, created_at, updated_at, handle, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (id) DO NOTHING;
    `
    _, err := db.Exec(query, product.ID, product.Title, product.BodyHTML, product.Vendor, product.ProductType, product.CreatedAt, product.UpdatedAt, product.Handle, product.Status)
    if err!= nil {
        return err
    }

    // Insert variants data
    query = `
        INSERT INTO variants (product_id, id, price, inventory_quantity)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (id) DO NOTHING;
    `
    for _, variant := range product.Variants {
        _, err := db.Exec(query, product.ID, variant.ID, variant.Price, variant.InventoryQuantity)
        if err!= nil {
            return err
        }
    }

    return nil
}

func main() {
    err := godotenv.Load()
    if err != nil {
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
