package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output(){
	fmt.Println(m)
}

func main() {
	// websites := map[string]string{
	// 	"Google":              "https://google.com",
	// 	"Amazon Web Services": "https://aws.com",
	// }
	// fmt.Println(websites)
	// fmt.Println(websites["Google"])

	// websites["Linkedin"] = "https://linkedin.com"
	// fmt.Println(websites)

	// delete(websites, "Google")
	// fmt.Println(websites)

	// courseRatings := map[string]float64{}
	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.6
	courseRatings["nodejs"] = 4.9

	// fmt.Println(courseRatings)
	courseRatings.output()
}