package service

import (
	"errors"
	"go-blog-management-system/model"
	"go-blog-management-system/repository"
)

type BlogService struct {
	BlogRepo *repository.BlogRepository
}

func NewBlogService(blogRepo *repository.BlogRepository) *BlogService {
	return &BlogService{BlogRepo: blogRepo}
}

// Add a new blog
func (service *BlogService) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	if blog.Title == "" || blog.Content == "" || blog.Author == "" {
		return nil, errors.New("invalid blog data")
	}

	createdBlog, err := service.BlogRepo.CreateBlog(blog)
	if err != nil {
		return nil, err
	}

	return createdBlog, nil
}

// Fetch blog by ID
func (service *BlogService) GetBlog(id int) (*model.Blog, error) {
	if id <= 0 {
		return nil, errors.New("invalid blog ID")
	}

	blog, err := service.BlogRepo.GetBlog(id)
	if err != nil {
		return nil, err
	}
	if blog == nil {
		return nil, errors.New("blog not found")
	}

	return blog, nil
}

// Fetch all blogs with pagination
func (service *BlogService) GetAllBlogs(page, limit int) ([]model.Blog, error) {
	if page <= 0 || limit <= 0 {
		return nil, errors.New("invalid pagination parameters")
	}

	blogs, err := service.BlogRepo.GetAllBlogs(page, limit)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

// Update blog details
func (service *BlogService) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	if blog.ID <= 0 || blog.Title == "" || blog.Content == "" || blog.Author == "" {
		return nil, errors.New("invalid blog data")
	}

	existingBlog, err := service.BlogRepo.GetBlog(blog.ID)
	if err != nil {
		return nil, err
	}
	if existingBlog == nil {
		return nil, errors.New("blog not found")
	}

	updatedBlog, err := service.BlogRepo.UpdateBlog(blog)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

// Delete a blog by ID
func (service *BlogService) DeleteBlog(id int) error {
	if id <= 0 {
		return errors.New("invalid blog ID")
	}

	existingBlog, err := service.BlogRepo.GetBlog(id)
	if err != nil {
		return err
	}
	if existingBlog == nil {
		return errors.New("blog not found")
	}

	return service.BlogRepo.DeleteBlog(id)
}
