package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/kartikey1188/go-students-api/internal/storage"
	"github.com/kartikey1188/go-students-api/internal/types"
	"github.com/kartikey1188/go-students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
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

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			*student.Age,
		)

		slog.Info("student created created successfully", slog.String("userID", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"ID": lastId})
	}
}
