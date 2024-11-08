package facade

import (
	"errors"
	"fmt"
)

type UserFacade struct {
	userService  *userService
	orderService *OrderService
}

func (f *UserFacade) GetUserDetails(id string) (string, error) {
	user := f.userService.GetUser(id)
	if user == "" {
		return "", errors.New("user not found")
	}

	order := f.orderService.GetOrder(id)
	return fmt.Sprintf("%s with %s", user, order), nil
}
