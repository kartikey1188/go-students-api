package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

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

		if validation := response.MissingFields(student); validation != "" {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("%s", validation)))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			*student.Age,
		)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		slog.Info("student created created successfully", slog.String("userID", fmt.Sprint(lastId)))

		response.WriteJson(w, http.StatusCreated, map[string]int64{"ID": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a student", slog.String("id", id))

		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id %s", id)))
			return
		}
		student, err := storage.GetStudent(intID)
		if err != nil {
			slog.Error("failed to get student", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all students")

		students, err := storage.GetStudents()
		if err != nil {
			slog.Error("failed to get students")
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		if len(students) == 0 {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("no students found")))
			return
		}
		response.WriteJson(w, http.StatusOK, students)
	}
}
