package service

import "github.com/phuhao00/wddd/blog/internal/app/domain/valueobject"

type BlogService interface {
	Update(post *valueobject.Post)
	Get(id string) *valueobject.Post
	Delete(id string)
}
