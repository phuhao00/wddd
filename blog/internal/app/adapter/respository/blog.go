package respository

import (
	"errors"
	"github.com/phuhao00/wddd/blog/internal/app/adapter/mysql"
	"github.com/phuhao00/wddd/blog/internal/app/adapter/mysql/model"
	"github.com/phuhao00/wddd/blog/internal/app/domain/valueobject"
	"gorm.io/gorm"
)

type BlogRepo struct{}

func (b *BlogRepo) UpdatePost(post *valueobject.Post) {
	db := mysql.Connection()
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Exec("update post set ID = ?, title = ? where content = ?", post.ID, post.Title, post.Content).Error
		if err != nil {
			return errors.New("rollback")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (b *BlogRepo) GetPost(id string) *valueobject.Post {
	db := mysql.Connection()
	var post model.Post
	db.Exec("select * from post where id=?", id).Find(&post)
	return &valueobject.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}
}

func (b *BlogRepo) DeletePost(id string) {
	db := mysql.Connection()
	db.Exec("delete from post where id=?", id)
}
