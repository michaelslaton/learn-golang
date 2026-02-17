package maps

import "fmt"

func maps() {
	websites := map[string]string{
		"Google":              "http://google.com",
		"Amazon Web Services": "http://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "http://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}