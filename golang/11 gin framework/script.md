[Intro]
"Hey everyone! Let's quickly set up a simple web API with the Gin framework and add Swagger for easy documentation! Here's how!"

[Step 1: Install Gin]
"First, install the Gin framework. Run this command in your terminal:"

bash
Copy code
go get github.com/gin-gonic/gin
[Step 2: Install Swagger]
"Next, install Swagger by running this:"

bash
Copy code
go install github.com/swaggo/swag/cmd/swag@latest
[Step 3: Add Swagger Comments]
"Now, let's add Swagger comments to your handlers. For example, above a simple GET endpoint, write this:"

go
Copy code
// @Summary Get all users
// @Tags users
// @Success 200 {array} User
// @Router /users [get]
r.GET("/users", func(c *gin.Context) {
    // Handler code here
})
[Step 4: Generate Docs]
"Generate Swagger documentation with this command:"

bash
Copy code
swag init
[Step 5: Serve Swagger UI]
"To serve the Swagger UI, add this line to your main file:"

go
Copy code
import "github.com/swaggo/gin-swagger"
import "github.com/swaggo/gin-swagger/swaggerFiles"

// In your main function:
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
[Start the Server]
"Now run your server and head over to /swagger/index.html to see your docs in action!"

[Outro]
"And that's it! You've got your API docs up and running with Gin and Swagger. Like and follow for more quick tips!"
