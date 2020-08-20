package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/alexeyco/simpletable"
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

var (
	data = [][]interface{}{
		{1, "iOS", "3 jobs"},
		{2, "Golang", "5 jobs"},
		{3, "ReactJS", "2 jobs"},
		{4, "VueJS", "1 job"},
		{5, "Android", "2 jobs"},
	}
)


func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
		c.String(200, Green("hello")+ Purple(" world!"))

	})

	router.GET("/table", func(c *gin.Context) {
		table := simpletable.New()

		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "#"},
				{Align: simpletable.AlignCenter, Text: "NAME"},
				{Align: simpletable.AlignCenter, Text: "TAX"},
			},
		}

		// subtotal := float64(0)
		for _, row := range data {
			r := []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
				{Text: row[1].(string)},
				{Align: simpletable.AlignRight, Text: row[2].(string)},
			}

			table.Body.Cells = append(table.Body.Cells, r)
			// subtotal += row[2].(string)
		}

		// table.Footer = &simpletable.Footer{
		// 	Cells: []*simpletable.Cell{
		// 		{},
		// 		{Align: simpletable.AlignRight, Text: "Subtotal"},
		// 		{Align: simpletable.AlignRight, Text: fmt.Sprintf("$ %.2f", subtotal)},
		// 	},
		// }

		table.SetStyle(simpletable.StyleCompactLite)
		c.String(200, table.String())
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
