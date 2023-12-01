package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CloudsIntersection(t *testing.T) {
	testCases := []struct {
		desc                  string
		clouds                Clouds
		expectedIntersections int
	}{
		{
			desc:                  "Horizontal & vertical",
			clouds:                []Cloud{{Point{1, 3}, Point{3, 3}}, {Point{3, 3}, Point{3, 4}}},
			expectedIntersections: 1,
		},
		{
			desc:                  "Vertical & Horizontal",
			clouds:                []Cloud{{Point{3, 3}, Point{3, 4}}, {Point{1, 3}, Point{3, 3}}},
			expectedIntersections: 1,
		},
		{
			desc:                  "Vertical multiple intersection",
			clouds:                []Cloud{{Point{3, 2}, Point{3, 4}}, {Point{3, 3}, Point{3, 2}}},
			expectedIntersections: 2,
		},
		{
			desc:                  "Horizontal multiple intersection",
			clouds:                []Cloud{{Point{2, 4}, Point{5, 4}}, {Point{3, 4}, Point{4, 4}}},
			expectedIntersections: 2,
		},
		{
			desc:                  "three intersections",
			clouds:                []Cloud{{Point{2, 4}, Point{5, 4}}, {Point{3, 4}, Point{4, 4}}, {Point{3, 3}, Point{3, 4}}},
			expectedIntersections: 2,
		},
		{
			desc:                  "vertical intersections",
			clouds:                []Cloud{{Point{2, 4}, Point{5, 4}}, {Point{5, 5}, Point{4, 4}}},
			expectedIntersections: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.expectedIntersections, tC.clouds.NumberOfIntersections())
		})
	}
}
