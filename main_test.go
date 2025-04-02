package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasks(t *testing.T) {
	req := httptest.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()

	getTasks(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", w.Code)
	}
}

func TestCreateTask(t *testing.T) {
	task := Task{Name: "Test Task", Status: 0}
	body, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
	w := httptest.NewRecorder()

	createTask(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status OK, got %v", w.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	task := Task{Name: "Test Task", Status: 0}
	tasks[1] = task

	updatedTask := Task{Name: "Updated Task", Status: 1}
	body, _ := json.Marshal(updatedTask)
	req := httptest.NewRequest("PUT", "/tasks/1", bytes.NewReader(body))
	w := httptest.NewRecorder()

	updateTask(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", w.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	task := Task{Name: "Test Task", Status: 0}
	tasks[1] = task

	req := httptest.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()

	deleteTask(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("Expected status No Content, got %v", w.Code)
	}
}
