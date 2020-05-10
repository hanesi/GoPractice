package main

import (
	"fmt"
	"os"

	// Need to install this and gonum.org/v1/gonum, but no import required for gonum
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	fp := "/Users/ianhanes/documents/github/gopractice/dataframestuff/nfl_data/NFL_data.csv"
	csv, err := os.Open(fp)
	if err != nil {
		fmt.Println(err)
	}

	df := dataframe.ReadCSV(csv)

	homeSeries := df.Col("home_score")
	awaySeries := df.Col("away_score")

	homeWin := homeWin(homeSeries, awaySeries)
	df = df.Mutate(homeWin)

	fmt.Println(df)
}

func homeWin(hs, as series.Series) series.Series {
	homeWSlice := []bool{}
	for i := 0; i < hs.Len(); i++ {
		if hs.Val(i).(int) > as.Val(i).(int) {
			homeWSlice = append(homeWSlice, true)
		} else {
			homeWSlice = append(homeWSlice, false)
		}
	}
	return series.New(homeWSlice, series.Bool, "homeWinSeries")
}
