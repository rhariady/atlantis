package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsRepo() models.Repo {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.Repo))(nil)).Elem()))
	var nullValue models.Repo
	return nullValue
}

func EqModelsRepo(value models.Repo) models.Repo {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.Repo
	return nullValue
}
