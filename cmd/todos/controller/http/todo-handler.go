package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/devlulcas/todoom/cmd/todos"
	"github.com/devlulcas/todoom/db"
)

type TodoHandler struct {
	TodoUseCase todos.TodoUseCase
}

func NewTodoHandler(todoUseCase todos.TodoUseCase) *TodoHandler {
	return &TodoHandler{todoUseCase}
}

func (c *TodoHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var todo todos.TodoInput

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	id, err := c.TodoUseCase.Create(todo)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusInternalServerError)
		return
	}

	var resp = map[string]int64{"id": id}

	sendResponseAsJson(w, resp, http.StatusCreated)
}

func (c *TodoHandler) HandleGetById(w http.ResponseWriter, r *http.Request) {
	id, err := getTodoId(r)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusBadRequest)
		return
	}

	todo, err := c.TodoUseCase.FindOne(id)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusInternalServerError)
		return
	}

	sendResponseAsJson(w, todo, http.StatusOK)
}

func (c *TodoHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	todo, pagination, err := c.TodoUseCase.FindAll(db.PaginationInput{
		Page:  1,
		Limit: 10,
	})

	if err != nil {
		sendErrorAsJson(w, err, http.StatusInternalServerError)
		return
	}

	var resp = map[string]interface{}{
		"todos":      todo,
		"pagination": pagination,
	}

	sendResponseAsJson(w, resp, http.StatusOK)
}

func (c *TodoHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := getTodoId(r)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusBadRequest)
		return
	}

	var todo todos.TodoInput

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err = c.TodoUseCase.Update(id, todo)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusInternalServerError)
		return
	}

	sendResponseAsJson(w, nil, http.StatusNoContent)
}

func (c *TodoHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := getTodoId(r)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusBadRequest)
		return
	}

	err = c.TodoUseCase.Delete(id)

	if err != nil {
		sendErrorAsJson(w, err, http.StatusInternalServerError)
		return
	}

	sendResponseAsJson(w, nil, http.StatusNoContent)
}

func sendErrorAsJson(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError := map[string]string{"error": err.Error()}
	json.NewEncoder(w).Encode(jsonError)

	log.Println(err)

	return
}

func sendResponseAsJson(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var resp = map[string]interface{}{"data": data}
	json.NewEncoder(w).Encode(resp)
}

func getTodoId(r *http.Request) (int64, error) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		return 0, err
	}

	return int64(id), nil
}
