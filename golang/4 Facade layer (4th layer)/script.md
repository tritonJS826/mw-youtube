[Intro]
"Hey everyone! Let's quickly dive into the Facade Layer in software architecture. Here's a simple breakdown of how it works!"

[Step 1: What is the Facade Layer?]
"The Facade Layer is like a gateway. It simplifies complex systems by providing a unified interface. Instead of interacting with multiple subsystems, you talk to just one!"

[Step 2: Why Use It?]
Let’s say you have different services: userService and Order servide, but you just need a simple method to interact with them."

[Step 3: Example in Go]
We’ll create a UserFacade to interact with user-related operations."

```go
package main

import (
    "fmt"
    "errors"
)

type UserService struct{}

func (s *UserService) GetUser(id int) string {
    return fmt.Sprintf("User ID: %d", id)
}

type OrderService struct{}

func (o *OrderService) GetOrder(id int) string {
    return fmt.Sprintf("Order ID: %d", id)
}

type UserFacade struct {
    userService  *UserService
    orderService *OrderService
}

func (f *UserFacade) GetUserDetails(id int) (string, error) {
    user := f.userService.GetUser(id)
    if user == "" {
        return "", errors.New("user not found")
    }
    order := f.orderService.GetOrder(id)
    return fmt.Sprintf("%s with %s", user, order), nil
}

func main() {
    userFacade := &UserFacade{
        userService:  &UserService{},
        orderService: &OrderService{},
    }
    
    details, err := userFacade.GetUserDetails(1)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(details)
    }
}```

[Step 4: Benefits of Facade Layer]
"As you can see, instead of dealing with both UserService and OrderService directly, we interact with a single UserFacade. This makes our code more organized!"

[Outro]
"And that's the Facade Layer in action! It simplifies complex systems. Don’t forget to like and follow for more tips!"



slides:
1. 3-Layers diagram
2. 4-layers diagram
3. print userService
4. print orderServie
5. create facade in the tree
6. print facade code
7. remove services logic from handler and replace it on facade 
