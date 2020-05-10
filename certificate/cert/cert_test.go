package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2020-06-08")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got=nil")
	}
	if c.Course != "Golang" {
		t.Errorf("Course name is not valid. Excepted='Golang', got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "bob", "2020-08-06")
	if err != nil {
		t.Error("Error should be returned on an empty course")
	}

}

func TestCoursetooLong(t *testing.T) {
	course := "abcdefghijklmnopqrstuvwxyz"
	_, err := New(course, "Bob", "2020-05-08")
	if err != nil {
		t.Errorf("Error should be returned on a too long course name (course=%s)", course)
	}
}

func TestNameEmptyValue(t *testing.T) {

}
