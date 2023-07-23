package respository

import "github.com/phuhao00/wddd/blog/internal/app/domain/valueobject"

type IBlog interface {
	Update(post *valueobject.Post)
	Get(id string) *valueobject.Post
	Delete(id string)
}
