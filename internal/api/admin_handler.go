package api

import (
	"encoding/json"
	"fmt"
	"github.com/abelmalu/CafeteriaAccessControl/internal/core"
	"github.com/abelmalu/CafeteriaAccessControl/internal/models"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type AdminHandler struct {
	service core.AdminService
}

type StandardResponse struct {
	Status  any         `json:"status"`
	Message interface{} `json:"message"`
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
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	w.Write(response)
}

// CreateCafeteria handles the api/admin/create/Cafeteria route
func (h *AdminHandler) CreateCafeteria(w http.ResponseWriter, r *http.Request) {
	var cafeteria models.Cafeteria
	err := json.NewDecoder(r.Body).Decode(&cafeteria)

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{
			"status":  "error",
			"message": err.Error(),
		}
		json, _ := json.Marshal(response)
		w.Write(json)

		return
	}
	validationError := cafeteria.Validate()
	if validationError != nil {
		response := StandardResponse{
			Status:  "error",
			Message: validationError,
		}

		respondWithJSON(w, 400, response)

		return

	}

	_, err = h.service.CreateCafeteria(r.Context(), &cafeteria)

	if err != nil {

		response := StandardResponse{
			Status:  "error",
			Message: err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		jsonResponse, _ := json.Marshal(response)

		w.Write(jsonResponse)

		return
	}

	response := map[string]string{
		"status":  "Success",
		"message": "Cafeteria Created Successfully",
	}

	json.NewEncoder(w).Encode(response)

}

func (h *AdminHandler) CreateBatch(w http.ResponseWriter, r *http.Request) {
	var batch models.Batch
	// decode the request body
	err := json.NewDecoder(r.Body).Decode(&batch)
	if err != nil {

		response := map[string]string{

			"status":  "error",
			"message": err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)

		return
	}
	validationError := batch.Validate()
	if validationError != nil {
		reponse := StandardResponse{

			Status:  "error",
			Message: validationError,
		}

		respondWithJSON(w, 400, reponse)

	}
	_, serviceErr := h.service.CreateBatch(r.Context(), &batch)
	if serviceErr != nil {

		response := StandardResponse{

			Status:  "error",
			Message: serviceErr.Error(),
		}
		respondWithJSON(w, 400, response)

		return

	}

	w.Write([]byte("successfully created a batch"))

}

func (h *AdminHandler) CreateMeal(w http.ResponseWriter, r *http.Request) {
	var meal models.Meal

	decodingErr := json.NewDecoder(r.Body).Decode(&meal)
	if decodingErr != nil {
		response := StandardResponse{
			Status:  "error",
			Message: decodingErr.Error(),
		}

		respondWithJSON(w, 400, response)
		return

	}

	_, err := h.service.CreateMeal(r.Context(), &meal)

	if err != nil {

		response := StandardResponse{
			Status:  "error",
			Message: err.Error(),
		}

		respondWithJSON(w, 400, response)
		return

	}
	json.NewEncoder(w).Encode("Successfully created a meal")

}

func (h *AdminHandler) RegisterDevice(w http.ResponseWriter, r *http.Request) {

	var device models.Device

	// log.Fatal("request reached here ")

	decodingErr := json.NewDecoder(r.Body).Decode(&device)

	if decodingErr != nil {
		response := StandardResponse{
			Status:  "error",
			Message: decodingErr.Error(),
		}

		respondWithJSON(w, 400, response)
		return

	}
	validationError := device.Validate()
	if validationError != nil {
		response := StandardResponse{
			Status:  "error",
			Message: validationError,
		}

		respondWithJSON(w, 400, response)

		return

	}
	_, err := h.service.RegisterDevice(r.Context(), &device)

	if err != nil {

		errorResponse := StandardResponse{

			Status:  "Error",
			Message: err.Error(),
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

	err := r.ParseMultipartForm(10 << 20)
	fmt.Printf("Multipart values: %+v\n", r.MultipartForm.Value)

	fmt.Println("hellow")

	if err != nil {

		response := StandardResponse{

			Status:  "error",
			Message: "couldn't parse the request",
		}
		responseJson, _ := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)

		return
	}

	student.MiddleName = r.FormValue("middle_name")
	student.LastName = r.FormValue("last_name")
	student.RFIDTag = r.FormValue("rfidTag")
	student.BatchId, _ = strconv.Atoi(r.FormValue("batch_id"))
	file, handler, err := r.FormFile("photo")
	validationError := student.Validate()
	if validationError != nil {
		response := StandardResponse{
			Status:  "error",
			Message: validationError,
		}

		respondWithJSON(w, 400, response)

		return

	}

	defer file.Close()

	uniqueID := uuid.New().String()
	extension := filepath.Ext(handler.Filename)
	newFilename := uniqueID + extension //  example image1 + .png = image1.png

	uploadsDir := os.Getenv("UPLOAD_DIR")

	photoPath := filepath.Join(uploadsDir, newFilename)
	fmt.Println("printing the photo's file name ")
	fmt.Println(newFilename)

	errr := os.MkdirAll(uploadsDir, 0755)
	if errr != nil {
		response := StandardResponse{
			Status:  "error",
			Message: errr.Error(),
		}

		respondWithJSON(w, 500, response)
		return
	}

	dst, err := os.Create(photoPath)
	if err != nil {
		response := StandardResponse{
			Status:  "error",
			Message: err.Error(),
		}

		respondWithJSON(w, 500, response)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		response := StandardResponse{
			Status:  "error",
			Message: err.Error(),
		}

		respondWithJSON(w, 500, response)
		return
	}

	student.ImageURL = newFilename
	
	_, err = h.service.CreateStudent(r.Context(), &student)
	if err != nil {

		response := StandardResponse{
			Status:  "error",
			Message: err.Error(),
		}
		respondWithJSON(w, 400, response)

		return
	}

	successResponse := StandardResponse{
		Status:  "success",
		Message: "Student Registered Successfully",
	}
	respondWithJSON(w, 200, successResponse)

}
