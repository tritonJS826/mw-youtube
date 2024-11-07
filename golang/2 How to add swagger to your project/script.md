[Intro]

"Hey everyone! Let’s quickly add Swagger to a Go project using the swag library."

[Step 1: Install Swag]
"First, install swag by running this in your terminal:"
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

[Step 2: Add Swagger Comments]
"Then add comments to your handlers. For example, above a GET /users endpoint, add:"
```go
// @Summary Get all users
// @Tags users
// @Success 200 {array} User
// @Router /users [get]
```

[Step 3: Generate Docs]
"Generate Swagger files with this command:"
```bash
swag init
```

[Step 4: Serve Swagger UI]
"To see it in action, add this to your main file for Gin projects:"
```go
r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
```

[Step 5: Start the Server]
"Run your server and go to /swagger/index.html to view your docs!"

[Outro]
"And there you go! Swagger is live. Like and follow for more quick tips!"
