package usecase

import (
	"github.com/phuhao00/wddd/blog/internal/app/application/service"
	"github.com/phuhao00/wddd/blog/internal/app/domain/valueobject"
)

func Update(svr service.BlogService, post *valueobject.Post) {
	svr.Update(post)
}
func Delete(svr service.BlogService, id string) {
	svr.Delete(id)
}
func Get(svr service.BlogService, id string) {
	svr.Get(id)
}
