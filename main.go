package main

import (
	"encoding/json"

	"fmt"
)

type student struct {
	name    string
	email   string
	class   string
	subject subject
	dob     dob
}
type dob struct {
	day   int
	month int
	year  int
}
type subject struct {
	subj1 string
	subj2 string
	subj3 string
}

//Now to create interface for calsculating areas of circle and and square

type Shape interface {
    Area() float64
}
type Circle struct {
    radius float64
}

type Square struct {
    length float64
}
// Circle implements Shape
func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}
// Square implements Shape
func (s Square) Area() float64 {
    return s.length * s.length
}
type Triangle struct {
    base   float64
    height float64
}
func (t Triangle) Area() float64 {
    return 0.5 * t.base * t.height
}
func calculateArea(listOfShapes []Shape) {
    for _, shape := range listOfShapes {
        fmt.Println("Area: ", shape.Area())
    }
}




var records map[int]student

func main() {

	c1 := Circle{radius: 5}
s1 := Square{length: 6}
t1 := Triangle{base: 6, height: 8}
shapes := []Shape{c1, s1, t1}
    calculateArea(shapes)


	records = make(map[int]student)

	records[1] = student{
		name:  "seyi",
		email: "seyi@yahoo.com",
		class: "ss2",
		subject: subject{
			subj1: "math",
			subj2: "english",
			subj3: "economics",
		},
		dob: dob{
			day:   3,
			month: 5,
			year:  2012,
		},
	}
	records[2] = student{
		name:  "femi johnson",
		email: "femi@yahoo.com",
		class: "ss2",
		subject: subject{
			subj1: "math",
			subj2: "english",
			subj3: "government",
		},
		dob: dob{
			day:   13,
			month: 6,
			year:  2012,
		},
	}
	records[3] = student{
		name:  "bunmi oluwa",
		email: "bunmi@yahoo.com",
		class: "ss2",
		subject: subject{
			subj1: "math",
			subj2: "english",
			subj3: "arabic",
		},
		dob: dob{
			day:   31,
			month: 4,
			year:  2011,
		},
	}

	// collate the key value and data contentst of the students records
	for t, q := range records {

		fmt.Println(t, q.name, q.email, q.class, q.subject, q.dob)
	}
	johnJSON, err := json.Marshal(records)
	if err == nil {
		fmt.Println(johnJSON)
	} else {
		fmt.Println(err)
	}
}
