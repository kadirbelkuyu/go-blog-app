package repository

import (
	"database/sql"
	"errors"
	"go-blog-app/internal/domain"
	"log"
)

type ArticleRepository interface {
	CreateArticle(article domain.Article) error
	FindArticleById(id uint) (domain.Article, error)
	UpdateArticle(id uint, authorId uint, u domain.Article) error
	GetArticles() ([]domain.Article, error)
	RemoveArticle(id uint, authorId uint) error
}

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r articleRepository) CreateArticle(article domain.Article) error {
	query := `INSERT INTO articles (title, content, author_id) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, article.Title, article.Content, article.AuthorID)
	if err != nil {
		log.Printf("create article error %v", err.Error())
		return errors.New("failed to create article")
	}
	return nil
}

func (r articleRepository) FindArticleById(id uint) (domain.Article, error) {
	var article domain.Article
	query := `SELECT * FROM articles WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content, &article.AuthorID, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		log.Printf("find article by id error %v", err)
		return domain.Article{}, errors.New("article does not exist")
	}

	return article, nil
}

func (r articleRepository) UpdateArticle(id uint, authorId uint, a domain.Article) error {
	query := `UPDATE articles SET title = $1, content = $2 WHERE id = $3 AND author_id = $4`
	result, err := r.db.Exec(query, a.Title, a.Content, id, authorId)
	if err != nil {
		log.Printf("error on update %v", err.Error())
		return errors.New("failed update article")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("error on get rows affected %v", err.Error())
		return errors.New("failed update article")
	}

	if rowsAffected < 1 {
		return errors.New("article not found or you are not the author of the article")
	}

	return nil
}

func (r articleRepository) GetArticles() ([]domain.Article, error) {
	var articles []domain.Article
	query := `SELECT * FROM articles`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("get articles error: %v", err)
		return nil, errors.New("failed to get articles")
	}

	for rows.Next() {
		var article domain.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.AuthorID, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			log.Printf("error scan article %v", err)
			return nil, errors.New("failed to get articles")
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r articleRepository) RemoveArticle(id uint, authorId uint) error {
	query := `DELETE FROM articles WHERE id = $1 AND author_id = $2`
	result, err := r.db.Exec(query, id, authorId)
	if err != nil {
		log.Printf("error on delete %v", err.Error())
		return errors.New("failed delete to article")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("error on get rows affected %v", err.Error())
		return errors.New("failed delete article")
	}

	if rowsAffected < 1 {
		return errors.New("article not found or you are not the author of the article")
	}

	return err
}
