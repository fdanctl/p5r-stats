package services

import (
	"fmt"
	"math"
	"strings"
)

type Metric struct {
	Label string
	Value float64
}

func BuildRadarSVG(width, height int, metrics []Metric, minValue, maxValue float64) string {
	cx := float64(width) / 2
	cy := float64(height) / 2
	radius := math.Min(cx, cy) * 0.7

	n := len(metrics)
	if n == 0 {
		return `<svg xmlns="http://www.w3.org/2000/svg"></svg>`
	}

	rangeValue := maxValue - minValue

	points := make([][2]float64, 0, n*2)
	for i := 0; i < n*2; i++ {
		angle := 2 * math.Pi * float64(i) / float64(n*2)
		if i%2 == 0 {
			idx := i / 2
			normalized := (metrics[idx].Value - minValue) / rangeValue
			r := normalized * radius
			x := cx + r*math.Sin(angle)
			y := cy - r*math.Cos(angle)
			points = append(points, [2]float64{x, y})
		} else {
			idx := i / 2
			min := math.Min(metrics[idx%n].Value, metrics[(idx+1)%n].Value)
			normalizedMin := (min - minValue) / rangeValue
			depthFactor := 0.5

			r := normalizedMin * radius * depthFactor
			x := cx + r*math.Sin(angle)
			y := cy - r*math.Cos(angle)
			points = append(points, [2]float64{x, y})
		}
	}

	out := fmt.Sprintf(
		`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`,
		width, height, width, height,
	)

	// Draw grid (optional: 4 concentric levels)
	levels := 4
	for l := 1; l <= levels; l++ {
		r := radius * float64(l) / float64(levels)
		out = fmt.Sprint(out, polygonRadial(cx, cy, r, n, `fill="none" stroke="#ddd" stroke-width="1"`))
	}

	for i := 0; i < n*2; i++ {
		if i%2 == 0 {
			angle := 2 * math.Pi * float64(i) / float64(n*2)
			x := cx + radius*math.Sin(angle)
			y := cy - radius*math.Cos(angle)
			out += fmt.Sprintf(
				`<line x1="%f" y1="%f" x2="%f" y2="%f" stroke="#ccc" stroke-width="1"/>`,
				cx, cy, x, y,
			)

			idx := i / 2
			labelR := radius * 1.12
			lx := cx + labelR*math.Sin(angle)
			ly := cy - labelR*math.Cos(angle)
			out += fmt.Sprintf(
				`<text x="%f" y="%f" font-size="12" text-anchor="middle" dominant-baseline="middle" fill="white" font-family="sans-serif">%s</text>`,
				lx, ly, metrics[idx].Label,
			)
		}
	}

	out += "<polygon points=\""
	for _, p := range points {
		out += fmt.Sprintf("%f,%f ", p[0], p[1])
	}
	out += `" fill="rgba(193, 154, 63, 0.7)" stroke="rgba(193, 154, 63, 1)" stroke-width="2"/>`

	out += "</svg>"
	return out
}

func polygonRadial(cx, cy, radius float64, n int, attrs string) string {
	points := ""
	for i := 0; i < n*2; i++ {
		if i%2 == 0 {
			angle := 2 * math.Pi * float64(i) / float64(n*2)
			x := cx + radius*math.Sin(angle)
			y := cy - radius*math.Cos(angle)
			points = fmt.Sprint(points, fmt.Sprintf("%f,%f ", x, y))
		} else {
			depthFactor := 0.5
			innerRadius := radius * depthFactor
			angle := 2 * math.Pi * float64(i) / float64(n*2)
			x := cx + innerRadius*math.Sin(angle)
			y := cy - innerRadius*math.Cos(angle)
			points = fmt.Sprint(points, fmt.Sprintf("%f,%f ", x, y))
		}
	}
	return fmt.Sprintf(`<polygon points="%s" %s />`, strings.TrimSpace(points), attrs)
}
