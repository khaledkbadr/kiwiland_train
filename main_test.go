package main

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"
)

var graph = map[string]map[string]int{}

type DistanceTestCase struct {
	Input    []string
	Expected int
	Error    error
}

func init() {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for _, node := range strings.Split(string(dat), ",") {
		trimmed := strings.TrimSpace(node)
		if len(trimmed) < 3 {
			panic(errors.New("invalid input"))
		}

		if _, ok := graph[string(trimmed[0])]; !ok {
			graph[string(trimmed[0])] = map[string]int{
				string(trimmed[1]): int(trimmed[2]) - 48,
			}
		} else {
			graph[string(trimmed[0])][string(trimmed[1])] = int(trimmed[2]) - 48
		}
	}
}

func TestDistanceBetween(t *testing.T) {
	railRoad := RailRoad{
		graph: graph,
	}

	distances := []DistanceTestCase{
		{Input: []string{"A", "B", "C"}, Expected: 9},
		{Input: []string{"A", "D"}, Expected: 5},
		{Input: []string{"A", "D", "C"}, Expected: 13},
		{Input: []string{"A", "E", "B", "C", "D"}, Expected: 22},
		{Input: []string{"A", "E", "D"}, Expected: 0, Error: ErrNoRoute},
	}

	for _, tc := range distances {
		distance, err := railRoad.DistanceBetween(tc.Input)
		if distance != tc.Expected || err != tc.Error {
			t.Errorf("test failed. expected %d, got %d", tc.Expected, distance)
		}
	}

}

type RoutesTestCase struct {
	Source      string
	Destination string
	Expected    int
	StopsNumber int
	Distance    int
	Error       error
}

func TestNumberOfRoutesMaxStops(t *testing.T) {
	railRoad := RailRoad{
		graph: graph,
	}

	distances := []RoutesTestCase{
		{Source: "C", Destination: "C", StopsNumber: 3, Expected: 2},
	}

	for _, tc := range distances {
		distance, err := railRoad.NumberOfRoutesMaxStops(tc.Source, tc.Destination, tc.StopsNumber)
		if distance != tc.Expected || err != tc.Error {
			t.Errorf("test failed. expected %d, got %d", tc.Expected, distance)
		}
	}

}

func TestNumberOfRoutesExactStops(t *testing.T) {
	railRoad := RailRoad{
		graph: graph,
	}

	distances := []RoutesTestCase{
		{Source: "A", Destination: "C", StopsNumber: 4, Expected: 3},
	}

	for _, tc := range distances {
		distance, err := railRoad.NumberOfRoutesExactStops(tc.Source, tc.Destination, tc.StopsNumber)
		if distance != tc.Expected || err != tc.Error {
			t.Errorf("test failed. expected %d, got %d", tc.Expected, distance)
		}
	}

}

func TestShortestPath(t *testing.T) {
	railRoad := RailRoad{
		graph: graph,
	}

	distances := []RoutesTestCase{
		{Source: "A", Destination: "C", Expected: 9},
		{Source: "B", Destination: "B", Expected: 9},
	}

	for _, tc := range distances {
		distance, err := railRoad.ShortestPath(tc.Source, tc.Destination)
		if distance != tc.Expected || err != tc.Error {
			t.Errorf("test failed. expected %d, got %d", tc.Expected, distance)
		}
	}

}

func TestRoutesMaxDistance(t *testing.T) {
	railRoad := RailRoad{
		graph: graph,
	}

	distances := []RoutesTestCase{
		{Source: "C", Destination: "C", Distance: 30, Expected: 7},
	}

	for _, tc := range distances {
		distance, err := railRoad.RoutesMaxDistance(tc.Source, tc.Destination, tc.Distance)
		if distance != tc.Expected || err != tc.Error {
			t.Errorf("test failed. expected %d, got %d", tc.Expected, distance)
		}
	}

}
