package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// Create a new blog
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var blog Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO blogs (title, content, author) VALUES (?, ?, ?)"
	_, err := DB.Exec(query, blog.Title, blog.Content, blog.Author)
	if err != nil {
		http.Error(w, "Failed to create blog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Blog created successfully"})
}

// Fetch a specific blog by ID
func BlogByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/blog/"):]

	switch r.Method {
	case http.MethodGet:
		getBlogByID(w, id)
	case http.MethodPut:
		updateBlogByID(w, r, id)
	case http.MethodDelete:
		deleteBlogByID(w, id)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getBlogByID(w http.ResponseWriter, id string) {
	var blog Blog
	query := "SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?"
	err := DB.QueryRow(query, id).Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	if err == sql.ErrNoRows {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch blog", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func updateBlogByID(w http.ResponseWriter, r *http.Request, id string) {
	var blog Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?"
	_, err := DB.Exec(query, blog.Title, blog.Content, blog.Author, id)
	if err != nil {
		http.Error(w, "Failed to update blog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Blog updated successfully"})
}

func deleteBlogByID(w http.ResponseWriter, id string) {
	query := "DELETE FROM blogs WHERE id = ?"
	_, err := DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete blog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Blog deleted successfully"})
}

// Fetch all blogs
func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, title, content, author, timestamp FROM blogs ORDER BY timestamp DESC")
	if err != nil {
		http.Error(w, "Failed to fetch blogs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var blogs []Blog
	for rows.Next() {
		var blog Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp); err != nil {
			http.Error(w, "Failed to scan blog", http.StatusInternalServerError)
			return
		}
		blogs = append(blogs, blog)
	}

	json.NewEncoder(w).Encode(blogs)
}
