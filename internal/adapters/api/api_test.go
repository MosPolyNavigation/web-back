package api

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestApi_GetPlan(t *testing.T) {
	type test struct {
		Campus string
		Corpus string
		Floor  int
	}

	var tests []test
	tests = append(tests,
		test{
			Campus: "BS",
			Corpus: "B",
			Floor:  1,
		},
		test{
			Campus: "BS",
			Corpus: "N",
			Floor:  1,
		},
	)

	api := New(logrus.New())

	for _, tst := range tests {
		plan, err := api.GetPlan(tst.Campus, tst.Corpus, tst.Floor)
		if err != nil {
			t.Error(err)
			return
		}
		logrus.Info(plan)
	}
}
