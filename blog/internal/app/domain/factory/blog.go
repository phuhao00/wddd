package factory

import (
	"github.com/phuhao00/wddd/blog/internal/app/domain"
)

type Blog struct {
}

func (b *Blog) GeneratePost() *domain.Blog {
	return &domain.Blog{}
}
