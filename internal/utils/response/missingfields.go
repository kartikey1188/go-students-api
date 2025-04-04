package response

import (
	"fmt"
	"strings"

	"github.com/kartikey1188/go-students-api/internal/types"
)

func MissingFields(student types.Student) string {
	missingFields := []string{}

	if student.Name == "" {
		missingFields = append(missingFields, "name")
	}
	if student.Email == "" {
		missingFields = append(missingFields, "email")
	}
	if student.Age == nil {
		missingFields = append(missingFields, "age")
	}

	if len(missingFields) == 0 {
		return "" // No missing fields
	}

	return fmt.Sprintf("Missing Fields: %s", strings.Join(missingFields, ", "))
}
