package services

import (
	"testing"
)

type MockProvider struct{}

func (*MockProvider) CallAPI(busStop int) (bytes []byte) {

}

func TestSumcGenerator_GenerateSchedule(t *testing.T) {
	busStop := 1
	provider := MockProvider{}
	generator := NewGenerator(&provider)

	schedule := generator.GenerateSchedule(busStop)


}
