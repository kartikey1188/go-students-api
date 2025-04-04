package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/kartikey1188/go-students-api/internal/types"
	"github.com/kartikey1188/go-students-api/internal/utils/getters"
	"github.com/kartikey1188/go-students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// validator := validator.New()
		// if err := validator.Struct(student); err != nil {

		// 	validatateErrs := err.(validator.ValidationErrors)
		// 	response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validatateErrs))
		// 	return
		// }

		//request validation

		if validation := response.MissingFields(student); validation != "" {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("%s", validation)))
			return
		}

		ageisthis := getters.GetAge(student)
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK", "age": strconv.Itoa(ageisthis)})
	}
}
