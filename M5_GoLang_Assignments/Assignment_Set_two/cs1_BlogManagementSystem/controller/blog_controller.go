package controller

import (
	"encoding/json"
	"go-blog-management-system/model"
	"go-blog-management-system/service"
	"net/http"
	"strconv"
	"strings"
)

type BlogController struct {
	BlogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
	return &BlogController{BlogService: blogService}
}

// Helper to write JSON responses
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Create a new blog
func (controller *BlogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	createdBlog, err := controller.BlogService.CreateBlog(&blog)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusOK, createdBlog)
}

// Get a blog by ID
func (controller *BlogController) GetBlog(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/blogs/")
	blogID, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid blog ID"})
		return
	}

	blog, err := controller.BlogService.GetBlog(blogID)
	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusOK, blog)
}

// Get all blogs with pagination
func (controller *BlogController) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || page <= 0 {
		limit = 10
	}

	// if page <= 0 || limit <= 0 {
	// 	writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Page and limit must be positive integers"})
	// 	return
	// }

	blogs, err := controller.BlogService.GetAllBlogs(page, limit)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusOK, blogs)
}

// Update a blog
func (controller *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/blogs/")
	blogID, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid blog ID"})
		return
	}

	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	blog.ID = blogID
	updatedBlog, err := controller.BlogService.UpdateBlog(&blog)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusOK, updatedBlog)
}

// Delete a blog
func (controller *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/blogs/")
	blogID, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid blog ID"})
		return
	}

	err = controller.BlogService.DeleteBlog(blogID)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]string{"message": "Blog deleted successfully"})
}
