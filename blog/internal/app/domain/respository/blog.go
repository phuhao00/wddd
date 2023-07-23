package respository

import "github.com/phuhao00/wddd/blog/internal/app/domain/valueobject"

type IBlog interface {
	UpdatePost(post *valueobject.Post)
	GetPost(id string) *valueobject.Post
	DeletePost(id string)
}
