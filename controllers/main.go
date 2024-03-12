package controllers

import (
	"sort"
	"summary-mvc/models"

	"github.com/gin-gonic/gin"
)

func ProcessArrays(c *gin.Context) {
	var body models.Body
	c.Bind(&body)
	var CombineArray []string
	var UniqueArray []string
	duplicate := map[string]int{}

	for i := 0; i < len(body.Ar1); i++ {
		if duplicate[body.Ar1[i]] == 0 {
			CombineArray = append(CombineArray, body.Ar1[i])
		}
		duplicate[body.Ar1[i]]++

	}
	for i := 0; i < len(body.Ar2); i++ {
		if duplicate[body.Ar2[i]] == 0 {
			CombineArray = append(CombineArray, body.Ar2[i])
		}
		duplicate[body.Ar2[i]]++
	}

	for key := range duplicate {

		if duplicate[key] == 1 {
			UniqueArray = append(UniqueArray, key)
		}
	}
	sort.Strings(CombineArray) //sort ให้เรียงตาม input ครับ
	sort.Strings(UniqueArray)
	c.JSON(200, gin.H{
		"combined": CombineArray,
		"unique":   UniqueArray,
	})
}
