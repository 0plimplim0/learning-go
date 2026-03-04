package main

import "fmt"

type Anime struct {
	Title    string
	Episodes int
}

var Animes = []Anime{
	{Title: "Mono", Episodes: 12},
	{Title: "Hell's Paradise", Episodes: 24},
	{Title: "Chainsaw Man", Episodes: 12},
	{Title: "Dan Da Dan", Episodes: 24},
	{Title: "Do It Yourself!!", Episodes: 12},
}

func filter(list []Anime, min int) []Anime {
	filtered := make([]Anime, 0, len(list))

	for _, v := range list {
		if v.Episodes > min {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func main() {
	fmt.Print("Welcome to the Slice filter! (For animes)\n\nNumber of episodes: ")
	var num int
	fmt.Scanln(&num)
	filtered := filter(Animes, num)
	fmt.Printf("\nAnimes that have more than %d episodes:\n\n", num)
	for _, v := range filtered {
		fmt.Printf("- %s | %d episodes\n", v.Title, v.Episodes)
	}
}
