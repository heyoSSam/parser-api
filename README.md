## **Installation & Setup**
### **1. Clone the Repository**
```sh
git clone https://github.com/heyoSSam/parser-api
cd parser-api
```

### **2. Install Dependencies**
Ensure you have Go installed (1.20+).
```sh
go mod tidy
```

### **3. Set Up Configuration**
Create a `config.yaml` file in the root directory:
```sh
touch config.yaml
```
Use the following template for configuration:
```yaml
server:
  port: 8080

front:
  URL: "YOUR_URL"
```

### **4. Run the Server**
```sh
go run cmd/main.go
```
The API will be available at `http://localhost:8080`.

---


