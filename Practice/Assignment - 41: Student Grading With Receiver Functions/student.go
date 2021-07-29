package main

// Declare a struct type to handle student marks with his identity
type StudentMarks struct {
	id          int
	name        string
	m_english   int
	m_maths     int
	m_science   int
	m_computers int
}

func (s *StudentMarks) validateResults() string {

	if s.m_computers < 35 || s.m_english < 35 || s.m_maths < 35 || s.m_science < 35 {
		return "Fail"
	}
	return "Pass"
}

func (s *StudentMarks) getTotalMarks() int {

	return (s.m_computers + s.m_english + s.m_maths + s.m_science)
}

func (s *StudentMarks) updateGrade() string {

	switch t := s.getTotalMarks(); {
	case t > 280:
		return "Distinction"
	case t > 240:
		return "First Class"
	case t > 160:
		return "Second Class"
	default:
		return "-"
	}
}
