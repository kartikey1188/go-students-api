package getters

import "github.com/kartikey1188/go-students-api/internal/types"

func GetAge(s types.Student) int {
	if s.Age == nil {
		return -1 // Return a default value (or handle it differently)
	}
	return *s.Age
}
