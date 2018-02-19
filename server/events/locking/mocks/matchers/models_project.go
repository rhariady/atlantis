package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsProject() models.Project {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.Project))(nil)).Elem()))
	var nullValue models.Project
	return nullValue
}

func EqModelsProject(value models.Project) models.Project {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.Project
	return nullValue
}
