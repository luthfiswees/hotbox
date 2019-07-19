package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"fmt"

	"github.com/luthfiswees/hotbox/model"
	"github.com/luthfiswees/hotbox/db"
	"github.com/minio/minio-go/v6"
)

// Response
type Response struct {
	Message string
}

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize bucket vars
	s3Host := os.Getenv("HOTBOX_S3_HOST")
	s3AccessKey := os.Getenv("HOTBOX_S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("HOTBOX_S3_SECRET_KEY")
	s3Bucket := os.Getenv("HOTBOX_S3_BUCKET")
	useSSL := false

	// Set response header to json
	w.Header().Set("Content-Type", "application/json")
	var respBody []byte

	// Check if request is POST
	if r.Method != "POST" {
		respBody, _ = json.Marshal(&Response{Message: "Only POST is Supported"})
		w.Write(respBody)
		return
	}

	// Parse form from request
	r.ParseMultipartForm(0)
	
	// Extract value from form
	serviceName := r.FormValue("service_name")
	status      := r.FormValue("status")
	timestamp   := r.FormValue("epoch_timestamp")
	if (serviceName == "" || status == "" || timestamp == "") {
		respBody, _ = json.Marshal(&Response{Message: "service_name, status, and timestamp have to be provided as params"})
		w.Write(respBody)
		return
	}

	// Extract file
	fileUploaded, fileHeader, err := r.FormFile("report")
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error receiving files"})
		w.Write(respBody)
		return
	}

	// Get filename
	filepath := "report/" + serviceName + "/" + timestamp + "/" + fileHeader.Filename

	// Connect to gorm db
	dbInstance, err := db.GetDatabaseInstance()
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error connecting to db"})
		w.Write(respBody)
		fmt.Println(err)
		return
	}

	// Store it using gorm
	reportEntry := model.ReportEntry{ServiceName: serviceName, Timestamp: timestamp, Status: status, ReportFile: s3Host + "/" + s3Bucket + "/" + filepath}
	dbInstance.Create(&reportEntry)
	dbInstance.Close()

	// Initialize minio client object.
	minioClient, err := minio.New(s3Host, s3AccessKey, s3SecretKey, useSSL)
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error connecting to s3"})
		w.Write(respBody)
		fmt.Println(err)
		return
	}

	// Open file
	fileContent, err := fileHeader.Open()
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error opening file header"})
		w.Write(respBody)
		fmt.Println(err)
		return
	}
	filesize, err := fileContent.Seek(0, 2)
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error opening file stat"})
		w.Write(respBody)
		fmt.Println(err)
		return
	}
	
	// Upload file to s3
	_, err = minioClient.PutObject(s3Bucket, filepath, fileUploaded, filesize, minio.PutObjectOptions{})
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error uploading file to s3"})
		w.Write(respBody)
		fmt.Println(err)
		return
	}

	// Return success message
	respBody, _ = json.Marshal(&Response{Message: "Success storing " + serviceName + " report to db"})
	w.Write(respBody)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header to json
	w.Header().Set("Content-Type", "application/json")
	var respBody []byte

	// Parse form
	r.ParseForm()

	// Get value from request
	serviceName := r.FormValue("service_name")
	status      := r.FormValue("status")

	// Connect to gorm db
	dbInstance, err := db.GetDatabaseInstance()
	defer dbInstance.Close()
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error connecting to db"})
		w.Write(respBody)
		fmt.Println(err)
		return
	}

	// Get data from gorm
	var reportEntries []model.ReportEntry
	_ = dbInstance.Where(&model.ReportEntry{ServiceName: serviceName, Status: status}).Find(&reportEntries)
	respBody, err = json.Marshal(reportEntries)
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error parsing json data from DB"})
		w.Write(respBody)
		fmt.Println(err)
		return
	}

	// Return data to client
	w.Write(respBody)
}