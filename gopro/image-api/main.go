package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Image struct to represent the image data
type Image struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Function to handle image upload
func uploadImage(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	r.ParseMultipartForm(10 << 20) // 10MB max

	// Get form data
	name := r.FormValue("name")
	description := r.FormValue("description")
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error retrieving image from form data: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save file to local directory
	filePath := "./uploads/" + handler.Filename
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error saving image: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	// Insert image data to MongoDB Atlas
	image := Image{
		Name:        name,
		Description: description,
		URL:         filePath, // Change to your cloud storage URL
	}
	insertImage(image)

	fmt.Fprintf(w, "Image uploaded successfully!")
}

// Function to insert image data into MongoDB Atlas
func insertImage(image Image) {
	// Connect to MongoDB Atlas
	clientOptions := options.Client().ApplyURI("mongodb+srv://snorpiii:cKFUGSfTgBDFbBrw@snorpiii.bjg2a6f.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB Atlas: ", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Get a handle for your collection
	collection := client.Database("<database>").Collection("go-api")

	// Insert image data
	_, err = collection.InsertOne(context.Background(), image)
	if err != nil {
		fmt.Println("Error inserting image data: ", err)
		return
	}
	fmt.Println("Image data inserted successfully!")
}

// Function to handle image deletion
func deleteImage(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Connect to MongoDB Atlas
	clientOptions := options.Client().ApplyURI("mongodb+srv://snorpiii:cKFUGSfTgBDFbBrw@snorpiii.bjg2a6f.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB Atlas: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	// Get a handle for your collection
	collection := client.Database("<database>").Collection("go-api")

	// Delete image data
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		fmt.Println("Error deleting image data: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Image deleted successfully!")
}

// Function to list all images
func listImages(w http.ResponseWriter, r *http.Request) {
	// Connect to MongoDB Atlas
	clientOptions := options.Client().ApplyURI("mongodb+srv://snorpiii:cKFUGSfTgBDFbBrw@snorpiii.bjg2a6f.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB Atlas: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	// Get a handle for your collection
	collection := client.Database("<database>").Collection("go-api")

	// Find all images
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("Error retrieving images: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var images []Image
	if err := cursor.All(context.Background(), &images); err != nil {
		fmt.Println("Error decoding images: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Display images
	for _, image := range images {
		fmt.Fprintf(w, "ID: %s, Name: %s, Description: %s, URL: %s\n", image.ID, image.Name, image.Description, image.URL)
	}
}

func main() {
	// Perform automatic migration if needed
	migrateDB()

	// Initialize router
	router := mux.NewRouter()

	// Route handler for image upload
	router.HandleFunc("/upload", uploadImage).Methods("POST")

	// Route handler for image deletion
	router.HandleFunc("/delete/{id}", deleteImage).Methods("DELETE")

	// Route handler for listing all images
	router.HandleFunc("/list", listImages).Methods("GET")

	// Serve index.html
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Start server
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}

// Function to perform automatic migration of the database schema
func migrateDB() {
	// Connect to MongoDB Atlas
	clientOptions := options.Client().ApplyURI("mongodb+srv://snorpiii:cKFUGSfTgBDFbBrw@snorpiii.bjg2a6f.mongodb.net/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB Atlas: ", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Get a handle for your collection
	collection := client.Database("<database>").Collection("go-api")

	// Create index on ID field if not exists
	indexView := collection.Indexes()
	_, err = indexView.CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.M{"id": 1},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		fmt.Println("Error creating index on ID field: ", err)
		return
	}

	fmt.Println("Database migration completed successfully!")
}
