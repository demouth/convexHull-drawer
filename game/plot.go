package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hamao0820/convexHull-drawer/graham"
)

var (
	NormalColor = color.Black
	ConvexColor = color.RGBA{255, 0, 0, 255}
)

type Plot struct {
	*graham.Point
	c        color.Color
	isConvex bool
	id       int
	r        int

	VX float64
	VY float64
}

var (
	plotID = 0
)

func NewPlot(x, y int) *Plot {
	plotID++
	return &Plot{
		Point:    graham.NewPoint(x, y),
		c:        NormalColor,
		isConvex: false,
		id:       plotID,
		r:        4,
	}
}

func (p *Plot) Update() error {
	if p.isConvex {
		p.c = ConvexColor
	} else {
		p.c = NormalColor
	}
	return nil
}

func (p *Plot) Draw(screen *ebiten.Image) {
	w, h := screen.Size()
	vector.DrawFilledCircle(screen, float32(p.X()+w/2), float32(p.Y()+h/2), float32(p.r), p.c, true)
}

func (p *Plot) near(x, y int) bool {
	return p.X()-p.r < x && x < p.X()+p.r && p.Y()-p.r < y && y < p.Y()+p.r
}
