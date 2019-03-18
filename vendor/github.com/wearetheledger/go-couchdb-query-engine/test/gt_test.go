package test

import (
	"github.com/wearetheledger/go-couchdb-query-engine/query"
	"testing"
)

func TestGt(t *testing.T) {

	t.Run("Size should be greater than 3", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$gt": 3,
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("TestGtNotEq should have returned 1 result")
		}
	})

	t.Run("Size should be greater than 0", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$gt": 1,
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 2 {
			t.Error("TestGtMultiple should have returned 3 result")
		}
	})
}
