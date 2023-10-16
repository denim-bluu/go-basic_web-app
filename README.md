# GoGo

GoGo is a simple Go web server that serves a "Hello, World!" message on port 3000.

## Installation

To install GoGo, you need to have Go installed on your machine. Once you have Go installed, you can run the following command to install GoGo:

```bash
go get github.com/gogo
```

## Usage

To start the GoGo server, run the following command:

```bash
go run main.go
```

This will start the server on port 3000. You can then access the server by navigating to <http://localhost:3000> in your web browser.

## Endpoints (Routes)

1. **Home Route**:  
   - Endpoint: `/`
   - HTTP Method: GET
   - Function: `showHome` This route serves the `index.html` file from the `./static/` directory.

2. **Get All Customers**:  
   - Endpoint: `/customers`
   - HTTP Method: GET
   - Function: `getCustomerDict` This route returns a JSON of all customers in the `customerDict` map.

3. **Get Single Customer**:  
   - Endpoint: `/customers/{id}`
   - HTTP Method: GET
   - Function: `getSignleCustomerDetail` This route fetches the details of a single customer based on the provided ID and returns them in JSON. If the customer ID does not exist, it returns a 404 status code.

4. **Create Single Customer**:  
   - Endpoint: `/customers`
   - HTTP Method: POST
   - Function: `createSingleCustomer` This route creates a new customer with a unique ID and adds it to the `customerDict` map. It then returns the updated `customerDict` in JSON.

5. **Update Single Customer**:  
   - Endpoint: `/customers/{id}`
   - HTTP Method: PATCH
   - Function: `updateSingleCustomer` This route updates the details of an existing customer based on the provided ID. It replaces the customer details in the `customerDict` map and returns the updated dictionary.

6. **Delete Single Customer**:  
   - Endpoint: `/customers/{id}`
   - HTTP Method: DELETE
   - Function: `deleteSingleCustomer`  This route deletes a customer based on the provided ID and returns the updated `customerDict` in JSON.
