package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	list  []int      
	mutex sync.Mutex 
)

func handleNumber(c *gin.Context) {
	var input struct {
		Number int `json:"number"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		log.Println("Error: invalid input")
		return
	}

	number := input.Number
	if number == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "zero not allowed"})
		log.Println("Error: zero not allowed")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if len(list) == 0 {
		list = append(list, number)
		log.Printf("List was empty, appended: %d", number)
		c.JSON(http.StatusOK, gin.H{"updated_list": list})
		return
	}

	if (list[0] > 0 && number > 0) || (list[0] < 0 && number < 0) {
		list = append(list, number)
		log.Printf("Same sign, appended %d → %v", number, list)
	} else {
		toReduce := abs(number)
		log.Printf("Opposite sign, reducing %d from %v", toReduce, list)

		for toReduce > 0 && len(list) > 0 {
			if abs(list[0]) > toReduce {
				if list[0] > 0 {
					list[0] -= toReduce
				} else {
					list[0] += toReduce
				}
				toReduce = 0
			} else {
				toReduce -= abs(list[0])
				list = list[1:]
			}
		}

		if toReduce > 0 {
			list = append(list, number/abs(number)*toReduce)
		}
		log.Printf("After reduction → %v", list)
	}

	c.JSON(http.StatusOK, gin.H{"updated_list": list})
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	r := gin.Default()
	r.GET("/number", func(c *gin.Context) {
    c.JSON(200, gin.H{"current_list": list})
})
	r.POST("/number", handleNumber)
	r.Run(":8080")
}
