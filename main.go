package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pet struct {
	ID         int    `json:"id"`
	UserName   string `json:"userName"`
	PetName    string `json:"petName"`
	PetImage   string `json:"petImage"`
	IsFriendly bool   `json:"isFriendly"`
}

func main() {
	r := gin.Default()

	pets := []Pet{
		{
			ID:         1,
			UserName:   "Posma Jan",
			PetName:    "Bret",
			PetImage:   "https://cdn.shibe.online/shibes/04be82add971f2b8490b5ec2308d3426e8494ad9.jpg",
			IsFriendly: true,
		},
		{
			ID:         2,
			UserName:   "Zidane Magaba",
			PetName:    "Antonette",
			PetImage:   "https://cdn.shibe.online/shibes/655be60d14f2a0dfb3060dae20c4aad0194b3393.jpg",
			IsFriendly: true,
		},
		{
			ID:         3,
			UserName:   "Rizky Fauzan",
			PetName:    "Samantha",
			PetImage:   "https://cdn.shibe.online/shibes/2d2e282fe2af42373600b819f80b869e1e7b18e9.jpg",
			IsFriendly: false,
		},
		{
			ID:         4,
			UserName:   "Fatol",
			PetName:    "Karianne",
			PetImage:   "https://cdn.shibe.online/shibes/8945cfa8f6a37efd24fa21fae6353b8c3537c381.jpg",
			IsFriendly: true,
		},
		{
			ID:         5,
			UserName:   "Blutak",
			PetName:    "Kamren",
			PetImage:   "https://cdn.shibe.online/shibes/0f305bd1f711737046266435fdbd1eac201329f7.jpg",
			IsFriendly: false,
		},
		{
			ID:         6,
			UserName:   "Anggun Prakoso",
			PetName:    "Aris",
			PetImage:   "https://cdn.shibe.online/shibes/6ddc38b0e28399ba27cd53019b0a5998cdbfdbed.jpg",
			IsFriendly: true,
		},
		{
			ID:         7,
			UserName:   "Reztol",
			PetName:    "Elwyn.Skiles",
			PetImage:   "https://cdn.shibe.online/shibes/6d05b676bb40d18686543ccbd92bff1f329ebd16.jpg",
			IsFriendly: false,
		},
		{
			ID:         8,
			UserName:   "Piandst",
			PetName:    "Maxime_Nienow",
			PetImage:   "https://cdn.shibe.online/shibes/c5bf11ac2dc4f660d15b7d31c37ce56203d73513.jpg",
			IsFriendly: true,
		},
	}

	r.GET("/pets", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": pets,
		})
	})

	r.Run(":6262")
}
