package main

import "fmt"

// Declare a struct type to handle student marks with his identity
type StudentMarks struct {
	id          int
	name        string
	m_english   int
	m_maths     int
	m_science   int
	m_computers int
}

func main() {

	// Create students list and their marks. We will store it in a slice.
	students := []StudentMarks{
		{id: 1, name: "Srikanth", m_english: 34, m_maths: 67, m_science: 45, m_computers: 90},
		{id: 2, name: "Rajiv", m_english: 89, m_maths: 90, m_science: 67, m_computers: 57},
		{id: 3, name: "Gokul", m_english: 40, m_maths: 47, m_science: 44, m_computers: 49},
		{id: 4, name: "Shyam", m_english: 34, m_maths: 67, m_science: 45, m_computers: 90},
		{id: 5, name: "Rahul", m_english: 38, m_maths: 71, m_science: 91, m_computers: 83},
		{id: 6, name: "Arvind", m_english: 48, m_maths: 56, m_science: 78, m_computers: 49},
		{id: 7, name: "Nivas", m_english: 90, m_maths: 100, m_science: 80, m_computers: 95},
		{id: 8, name: "Lalitha", m_english: 89, m_maths: 100, m_science: 87, m_computers: 89},
		{id: 9, name: "Rameez", m_english: 70, m_maths: 66, m_science: 89, m_computers: 54},
		{id: 10, name: "Sonia", m_english: 80, m_maths: 56, m_science: 56, m_computers: 34},
	}

	// declare a variable to hold max marks
	maxMarks := 0
	// declare a variable to hold toper
	var topper StudentMarks
	// iterate through the slice to find the toper
	for _, s := range students {
		// calculate total marks
		total := s.m_computers + s.m_english + s.m_maths + s.m_science
		// compare if total marks are higher than maxMarks we have already seen
		if total > maxMarks {
			// if it higher, then replace maxMarks with current total
			maxMarks = total
			// also track the toper
			topper = s
		}
	}
	// print the result at the end
	fmt.Printf("Toper Id: %d | Name: %s | Top Marks: %d\n", topper.id, topper.name, maxMarks)
}
