package main

import "fmt"

func main() {
	fmt.Println("Maps in golang") //Maps in golang

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages) //List of all languages map[JS:javascript PY:Python RB:Ruby]
	fmt.Println("JS shotrs for: ", languages["JS"]) //JS shotrs for:  javascript

	delete(languages, "RB")
	fmt.Println("List of all languages", languages) //List of all languages map[JS:javascript PY:Python]

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value) //for key JS, value is javascript ,for key PY, value is Python

	}
}
