package repository

import (
	"database/sql"
	"go-blog-management-system/model"
	"log"
)

type BlogRepository struct {
	DB *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{DB: db}
}

// ----------- Blog Repository Methods ----------- //

func (repo *BlogRepository) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO blogs (title, content, author, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(blog.Title, blog.Content, blog.Author, blog.Timestamp)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	blog.ID = int(id)
	return blog, nil
}

func (repo *BlogRepository) GetBlog(id int) (*model.Blog, error) {
	row := repo.DB.QueryRow("SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?", id)
	blog := &model.Blog{}
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (repo *BlogRepository) GetAllBlogs(page, limit int) ([]model.Blog, error) {
	offset := (page - 1) * limit
	query := "SELECT id, title, content, author, timestamp FROM blogs ORDER BY timestamp DESC LIMIT ? OFFSET ?"
	rows, err := repo.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.Blog
	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
		if err != nil {
			log.Fatal(err)
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (repo *BlogRepository) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("UPDATE blogs SET title = ?, content = ?, author = ?, timestamp = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author, blog.Timestamp, blog.ID)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (repo *BlogRepository) DeleteBlog(id int) error {
	stmt, err := repo.DB.Prepare("DELETE FROM blogs WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
