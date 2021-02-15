package main

import (
	"fmt"
	"math"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// getArea gets the area inside from a given shape
func getArea(points []Point) (area float64) {

	a := 0.0
	b := 0.0
	
	for i:=0 ; i<len(points) - 3; i++ {
	
		if doIntersect(points[i], points[i+1], points[i+2], points[i+3]) == true {

			fmt.Println("No valid area")
			return 0
		
		}
	}
	for i:=1 ; i<len(points); i++ {
	
		a = points[i-1].X*points[i].Y + a
		b = -(points[len(points) - i].X * points[len(points)-i-1].Y) + b
		
		
	}
	
	c := points[len(points) -1].X*points[0].Y
	d := -points[0].X*points[len(points) -1].Y
	area = math.Abs((a+b+c+d))/2
	return area
	
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	
	for i:=0 ; i<len(points) - 3; i++ {
	
		if doIntersect(points[i], points[i+1], points[i+2], points[i+3]) == true {

			fmt.Println("No valid perimeter")
			return 0
		
		}
	}
	Perimeter:=0.0
	for i:=1 ; i<len(points); i++ {
		Perimeter = math.Sqrt(math.Pow(points[i].X - points[i-1].X,2) + math.Pow(points[i].Y - points[i-1].Y,2)) + Perimeter
		fmt.Println(Perimeter)
	}
	FinalLength := math.Sqrt(math.Pow(points[len(points)-1].X - points[len(points)-2].X,2) + math.Pow(points[len(points)-1].Y - points[len(points)-2].Y,2))
	TotalPerimeter := Perimeter + FinalLength
	return TotalPerimeter
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome friend to the Remote Shapes Analyzer\n")
	response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
	response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
	response += fmt.Sprintf(" - Area            : %v\n", area)

	// Send response to client
	fmt.Fprintf(w, response)
}

func isOnSegment(p,q,r Point) bool {

	if q.X <= math.Max(p.X, r.X) && q.X >= math.Min(p.X, r.X) && q.Y <= math.Max(p.Y, r.Y) && q.Y >= math.Min(p.Y, r.Y) {

		return true

	}else {
		return false
	}

}

func orientation(p,q,r Point) int {

	val := (q.Y - p.Y)*(r.X-q.X) - (q.X - p.X)*(r.Y-q.Y)
	if val == 0 {
		return 0
	}else if val > 0 {
		return 1
	}else {
		return 2
	}

}

func doIntersect(p1, q1, p2, q2 Point) bool {

	o1 := orientation(p1, q1, p2) 
	o2 := orientation(p1, q1, q2) 
	o3 := orientation(p2, q2, p1) 
	o4 := orientation(p2, q2, q1) 

	if o1 != o2 && o3 != o4 {
		return true
	}

	if o1 == 0 && isOnSegment(p1, p2, q1) {
		return true
	}

	if o2 == 0 && isOnSegment(p1, q2, q1) {
		return true
	}

	if o3 == 0 && isOnSegment(p2,p1,q2) {
		return true
	}

	if o4 == 0 && isOnSegment(p2,q1,q2) {
		return true
	}

	return false

}