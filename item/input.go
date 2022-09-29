package item

import "test_aac/user"

type GetItemDetailInput struct {
	ID int `uri:"id" binding: "required"`
}

type CreateItemInput struct {
	Name  string `json: "name" binding:"required"`
	Price string `json: "price" binding:"required"`
	Cost  string `json: "cost" binding:"required"`
	User  user.User
}
