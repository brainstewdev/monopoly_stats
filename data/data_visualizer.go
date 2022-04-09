package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

/*
	this script allows you to analyze the game data generated
	it can:
		- tell you which is the most landed square on the games
        - percentage of landing of cells
*/

func main() {
	// load all the data from jsons
	files, err := ioutil.ReadDir("./jsons")
	if err != nil {
		fmt.Println("error occured:", err)
		return
	}
	gamesValues := make(map[int]int)

	for _, v := range files {
		if !v.IsDir() {
			// create temporary map
			currentValues := make(map[int]int)
			// open file
			path := "./jsons/"
			filename := filepath.Join(path, v.Name())
			f, err := os.Open(filename)
			if err != nil {
				fmt.Println(err)
			}
			byteValue, err := ioutil.ReadAll(f)
			if err == nil {
				json.Unmarshal(byteValue, &currentValues)
				// for each element inside current value sum it to the general map
				for k, v := range currentValues {
					gamesValues[k] += v
				}
			} else {
				fmt.Println(err)
			}
			f.Close()
		}
		fmt.Println()
	}
	fmt.Println("general values:")
	// get the keys from the map and add them to a vector
	var keys []int
	var values []int
	for k, v := range gamesValues {
		fmt.Println("square", k, ":", v)
		keys = append(keys, k)
		values = append(values, v)
	}
	sort.Ints(keys)
	// create a new bar instance
	bar := charts.NewBar()

	// Set global options
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "most landed on square",
		Subtitle: "",
	}))

	// Put data into instance
	bar.SetXAxis(keys).
		AddSeries("square landed", generateBarItems(values))
	f, _ := os.Create("bar.html")
	_ = bar.Render(f)
}

func generateBarItems(values []int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(values); i++ {
		items = append(items, opts.BarData{Value: values[i]})
	}
	return items
}
