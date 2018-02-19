package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyPtrToModelsProjectLock() *models.ProjectLock {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*models.ProjectLock))(nil)).Elem()))
	var nullValue *models.ProjectLock
	return nullValue
}

func EqPtrToModelsProjectLock(value *models.ProjectLock) *models.ProjectLock {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *models.ProjectLock
	return nullValue
}
