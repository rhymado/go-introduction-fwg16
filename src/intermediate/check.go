package intermediate

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

func Check() {
	urls := []string{
		"https://stackoverflow.com",
		"https://github.com",
		"https://www.linkedin.com",
		"https://medium.com",
		"https://golang.org",
		"https://www.fazztrack.com",
		"https://www.coursera.org",
		"https://wesionary.team",
	}

	for _, url := range urls {
		wg.Add(1)
		go webChecker(url)
		// webChecker(url)
	}
	wg.Wait()
}

func webChecker(url string) {
	start := time.Now()
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(url, " terjadi Error")
		return
	}
	fmt.Printf("[%d] %s is UP, started at %d\n", res.StatusCode, url, start.Second())
}

func Cube(side float64) (float64, float64, float64) {
	surfaceArea := 6 * math.Pow(side, 2)
	circumference := 4 * side
	volume := math.Pow(side, 3)
	return surfaceArea, circumference, volume
}
