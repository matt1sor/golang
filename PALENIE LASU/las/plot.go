package las

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"os"
	"strconv"
)

func PlotData() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Błąd podczas otwierania pliku:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	p := plot.New()

	data := make(plotter.XYs, len(records))
	for i, record := range records {
		x, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			panic(err)
		}
		data[i].X = x
		data[i].Y = y
	}

	scatter, err := plotter.NewScatter(data)
	if err != nil {
		panic(err)
	}
	scatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	scatter.GlyphStyle.Radius = vg.Points(4)

	// Add the scatter plot to the plot and set the axes labels
	p.Add(scatter)
	p.Title.Text = "Spalone i poczatkowe drzewa"
	p.X.Label.Text = "Poczatkowe drzewa "
	p.Y.Label.Text = "Spalone Drzewa"

	// Save the plot to a PNG file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "scatter.png"); err != nil {
		panic(err)
	}

	//p.Title.Text = "histogram plot"
	//
	//hist, err := plotter.NewHist(records, 20)
	//if err != nil {
	//	panic(err)
	//}
	//p.Add(hist)
	//
	//if err := p.Save(3*vg.Inch, 3*vg.Inch, "hist.png"); err != nil {
	//	panic(err)
	//}
}
