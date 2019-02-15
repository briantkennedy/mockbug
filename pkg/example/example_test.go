package example_test

import (
	"github.com/briantkennedy/mockbug/pkg/example"
	"github.com/briantkennedy/mockbug/pkg/mock_example"
	"github.com/golang/mock/gomock"
	"testing"
)


func TestNilVsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectNilSliceArg := &example.ThingApplyArg{}

	m := mock_example.NewMockThing(ctrl)
	m.EXPECT().
		Apply(gomock.Eq(expectNilSliceArg))

	emptySliceArg := &example.ThingApplyArg{
		SomeList: []string{},
	}
	m.Apply(emptySliceArg)
}

