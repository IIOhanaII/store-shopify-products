# Store Shopify Products

This Go project is designed to fetch product data from a Shopify store and store it in a database for easier bulk modification. Once data is stored, users can manipulate it before sending the modified data back to Shopify using the companion project [Update Shopify Products](https://github.com/IIOhanaII/update-shopify-products).

## Features

- **Fetch Shopify product data**: Retrieve all product information from your Shopify store and store it in a database.
- **Bulk modification**: Modify the product data directly within the database for faster bulk operations.
- **Seamless integration**: Easily push changes back to your Shopify store using the companion project.

## Prerequisites

1. **Go**: Ensure that you have Go installed. You can download it [here](https://golang.org/dl/).
2. **PostgreSQL Database**: Set up a PostgreSQL database where the Shopify product data will be stored.
3. **Shopify Store**: You must have a Shopify store with the necessary API credentials.
4. **.env File**: Create a `.env` file in the root directory of your project to store sensitive variables like API keys, database credentials, etc.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/IIOhanaII/store-shopify-products.git
    cd store-shopify-products
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Set up your `.env` file:

    Create a `.env` file in the root directory of the project and include the following sensitive variables:

    ```
    SHOP_NAME=your_shopify_store_name
    SHOPIFY_ACCESS_TOKEN=your_shopify_access_token
    POSTGRES_DBNAME=your_database_name
    POSTGRES_USER=your_database_user
    POSTGRES_PASSWORD=your_database_password
    POSTGRES_HOST=your_database_host
    POSTGRES_PORT=your_database_port
    ```

   **Explanation of variables**:
   - `SHOP_NAME`: Your Shopify store name (e.g., `yourstore` within `yourstore.myshopify.com`).
   - `SHOPIFY_ACCESS_TOKEN`: The access token for authenticating with the Shopify API.
   - `POSTGRES_DBNAME`: Name of the PostgreSQL database.
   - `POSTGRES_USER`: PostgreSQL database user.
   - `POSTGRES_PASSWORD`: Password for the PostgreSQL database user.
   - `POSTGRES_HOST`: Host address of your PostgreSQL database (e.g., `localhost`).
   - `POSTGRES_PORT`: Port on which the PostgreSQL database is running (default is `5432`).

4. Run the project:

    ```bash
    go run main.go
    ```

## Usage

1. **Fetching Shopify products**:
   Once the project is running, it will retrieve all product data from your Shopify store and store it in the connected PostgreSQL database.

2. **Modifying product data**:
   After the data is fetched, you can manually modify it in the database (e.g., via SQL queries, a database management tool, or scripts).

3. **Sending modified data back to Shopify**:
   After making bulk modifications, use the companion project [Update Shopify Products](https://github.com/IIOhanaII/update-shopify-products) to push the changes back to your Shopify store.

## Project Structure

```bash
├── .env              # Contains sensitive environment variables
├── main.go           # Main entry point of the application
├── db                # Contains database connection and data handling logic
├── shopify           # Shopify API interaction logic
├── go.mod            # Go module dependencies
└── README.md         # Project documentation
```

## Environment Variables

All sensitive information should be stored in the `.env` file for security purposes. The project uses the [godotenv](https://github.com/joho/godotenv) package to load environment variables from the `.env` file into the application.

## Database Configuration

The project is designed to work with PostgreSQL, but feel free to modify it to work with other databases if needed. Ensure that your PostgreSQL connection information is properly set up in the `.env` file.

Example PostgreSQL connection string in the `.env`:

```
POSTGRES_DBNAME=shopify_db
POSTGRES_USER=postgres
POSTGRES_PASSWORD=yourpassword
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
```

## Contributing

If you'd like to contribute to this project, feel free to submit a pull request or file an issue. Any contributions, such as bug fixes, new features, or improvements to the documentation, are welcome.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Related Projects

- [Update Shopify Products](https://github.com/IIOhanaII/update-shopify-products): Use this project to send your modified Shopify product data back to the store after making bulk changes.

---

Feel free to modify this as needed!
