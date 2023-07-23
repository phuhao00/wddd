package domain

import "github.com/phuhao00/wddd/blog/internal/app/domain/valueobject"

type Blog struct {
	Posts []*valueobject.Post
}
