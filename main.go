package main

import (
	"fmt"
	"image/color"
	"log"

	"gonum.org/v1/plot/vg/draw"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func main() {

	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) plotter.XYs {
		pts := make(plotter.XYs, 4)
		pts[0].X = 1
		pts[0].Y = 1
		pts[1].X = 1
		pts[1].Y = 2
		pts[2].X = 2
		pts[2].Y = 1
		pts[3].X = 2
		pts[3].Y = 2

		return pts
	}

	n := 15
	scatterData := randomPoints(n)

	p, err := plot.New()
	p.HideAxes()
	if err != nil {
		log.Panic(err)
	}

	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		log.Panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Shape = draw.BoxGlyph{}
	s.GlyphStyle.Radius = 20

	p.Add(s)
	fmt.Println(scatterData)
	err = p.Save(200, 200, "scatter.png")
	if err != nil {
		log.Panic(err)
	}
}
