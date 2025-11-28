package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	//"fmt"
	//"fmt"
	"io"
	"os"

	//"strconv"
	"path/filepath"

	//"log"
	"net/http"

	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
	"github.com/google/uuid"
)

type AdminHandler struct {
	service core.AdminService
}

type StandardResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewAdminHandler(service core.AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

// --- 2. JSON Utility Function ---

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status": "error", "message": "Internal JSON encoding error"}`))
		return
	}

	// MANDATORY: Set the header
	w.Header().Set("Content-Type", "application/json")

	// Set the status code
	w.WriteHeader(code)

	// Write the JSON body
	w.Write(response)
}

// CreateCafeteria handles the api/admin/create/Cafeteria route
func (h *AdminHandler) CreateCafeteria(w http.ResponseWriter, r *http.Request) {
	var cafeteria models.Cafeteria
	err := json.NewDecoder(r.Body).Decode(&cafeteria)

	if err != nil {

		http.Error(w, "Invalid data sent", http.StatusBadRequest)
		w.Header().Set("Content-Type", ":application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json, _ := json.Marshal([]byte("Invalid data sent"))
		w.Write(json)

		return
	}

	created, err := h.service.CreateCafeteria(r.Context(), &cafeteria)

	if err != nil {
		errorString := err.Error()

		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}
	message := "Cafeteria successfully created"

	json.NewEncoder(w).Encode(created)
	json.NewEncoder(w).Encode(message)

}

func (h *AdminHandler) CreateBatch(w http.ResponseWriter, r *http.Request) {
	var batch models.Batch
	// decode the request body
	err := json.NewDecoder(r.Body).Decode(&batch)
	if err != nil {

		http.Error(w, "Bad Request", http.StatusBadRequest)

		return
	}
	created, serviceErr := h.service.CreateBatch(r.Context(), &batch)
	if serviceErr != nil {

		errorString := serviceErr.Error()

		http.Error(w, errorString, http.StatusBadRequest)

		return

	}

	w.Write([]byte("successfully created a batch"))
	json.NewEncoder(w).Encode(created)

}

func (h *AdminHandler) CreateMeal(w http.ResponseWriter, r *http.Request) {
	var meal models.Meal

	decodingErr := json.NewDecoder(r.Body).Decode(&meal)
	if decodingErr != nil {
		errorString := decodingErr.Error()

		http.Error(w, errorString, http.StatusBadRequest)
		return

	}
	_, err := h.service.CreateMeal(r.Context(), &meal)

	if err != nil {

		errorString := err.Error()

		http.Error(w, errorString, http.StatusBadRequest)
		return

	}
	json.NewEncoder(w).Encode([]byte("Successfully created a meal"))

}

func (h *AdminHandler) RegisterDevice(w http.ResponseWriter, r *http.Request) {

	var device models.Device

	// log.Fatal("request reached here ")

	decodingErr := json.NewDecoder(r.Body).Decode(&device)

	if decodingErr != nil {
		errorString := decodingErr.Error()
		http.Error(w, errorString, http.StatusBadRequest)
		return

	}
	_, err := h.service.RegisterDevice(r.Context(), &device)

	if err != nil {

		errorResponse := StandardResponse{

			Status:  "Error",
			Message: "Invalid Cafeteria ID",
		}

		// http.Error(w, errString, http.StatusBadRequest)
		respondWithJSON(w, http.StatusBadRequest, errorResponse)
		return

	}

	successResponse := StandardResponse{
		Status:  "success",
		Message: "Device Registered Successfully",
	}

	respondWithJSON(w, http.StatusCreated, successResponse)

}

// CreateStudent handles the api/admin/create/student route
func (h *AdminHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student

	// if err := json.NewDecoder(r.Body).Decode(&student); err != nil {

	// 	http.Error(w, "invalid input", http.StatusBadRequest)

	// 	return
	// }

	// created, err := h.service.CreateStudent(r.Context(), &student)
	// if err != nil {
	// 	errString := err.Error()
	// 	http.Error(w, errString, http.StatusInternalServerError)
	// 	return
	// }

	// json.NewEncoder(w).Encode(created)

	err := r.ParseMultipartForm(10)

	if err != nil {

		response := StandardResponse{

			Status:  "error",
			Message: "couldn't parse the request",
		}
		responseJson, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
	}

	student.FirstName = r.FormValue("first_name")
	student.MiddleName = r.FormValue("middle_name")
	student.LastName = r.FormValue("last_name")
	student.RFIDTag = r.FormValue("rfidTag")
	student.BatchId, _ = strconv.Atoi(r.FormValue("batch_id"))
	// 3. Extract the uploaded file
	file, handler, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "photo is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique filename using uuid and preserve the original extension.
	uniqueID := uuid.New().String()
	extension := filepath.Ext(handler.Filename)
	newFilename := uniqueID + extension

	// Save the image in the static folder using the unique filename
	uploadsDir := filepath.Join("..", "..", "uploads")
	photoPath := filepath.Join(uploadsDir, newFilename)

	// 1. Ensure the directory exists (and create it if it doesn't)
	errr := os.MkdirAll(uploadsDir, 0755)
	if errr != nil {
		fmt.Println(errr)
		http.Error(w, "failed to create directory: "+errr.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Now attempt to create the file inside the confirmed existing directory
	dst, err := os.Create(photoPath)
	if err != nil {
		fmt.Println(err)
		// This error is now specifically about file creation within the existing directory
		http.Error(w, "failed to save photo", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file contents to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)
		http.Error(w, "failed to save photo", http.StatusInternalServerError)
		return
	}

	// 5. Store photo path in DB (store the filename or relative path as desired)
	student.ImageURL = newFilename

	// 6. Call service

	// 6. Call service
	created, err := h.service.CreateStudent(r.Context(), &student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 7. Return JSON
	json.NewEncoder(w).Encode(created)

}
