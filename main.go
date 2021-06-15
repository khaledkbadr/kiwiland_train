package main

import (
	"errors"
	"math"
)

func main() {

}

var ErrNoRoute = errors.New("NO SUCH ROUTE")

type RailRoad struct {
	graph map[string]map[string]int
}

func (r *RailRoad) DistanceBetween(stops []string) (int, error) {
	var distances int

	for i, node := range stops {
		if i+1 < len(stops) {
			nextStop := stops[i+1]
			distance, ok := r.graph[node][nextStop]
			if !ok {
				return 0, ErrNoRoute
			}
			distances += distance
		}
	}
	return distances, nil
}

func (r *RailRoad) NumberOfRoutesMaxStops(start, end string, maxStops int) (int, error) {
	return r.DFSMaxStops(start, end, maxStops)
}

func (r *RailRoad) NumberOfRoutesExactStops(start, end string, exactStops int) (int, error) {
	return r.DFSExactStops(start, end, exactStops)
}

func (r *RailRoad) ShortestPath(start, end string) (int, error) {
	_, ok := r.graph[start]
	if !ok {
		return 0, ErrNoRoute
	}

	_, ok = r.graph[end]
	if !ok {
		return 0, ErrNoRoute
	}

	// table will have shortest paths between start to all the other nodes
	pathTable := make(map[string]int)
	for key := range r.graph {
		pathTable[key] = math.MaxInt64
	}
	pathTable[start] = 0

	visited := make(map[string]struct{})
	queue := []string{start}

	for len(queue) > 0 {
		parent, q := queue[0], queue[1:]
		queue = q
		for child, weight := range r.graph[parent] {
			if _, ok := visited[child]; ok {
				continue
			}
			relaxedWeight := pathTable[parent] + weight

			if start == end && pathTable[child] == 0 {
				pathTable[child] = relaxedWeight
			} else if pathTable[child] > relaxedWeight {
				pathTable[child] = relaxedWeight
			}
			queue = append(queue, child)
		}

		if parent != start {
			visited[parent] = struct{}{}
		}
	}

	return pathTable[end], nil
}

func (r *RailRoad) RoutesMaxDistance(start, end string, maxDistance int) (int, error) {
	_, ok := r.graph[start]
	if !ok {
		return 0, ErrNoRoute
	}

	_, ok = r.graph[end]
	if !ok {
		return 0, ErrNoRoute
	}

	if maxDistance <= 0 {
		return 0, nil
	}

	var routes int

	for next, weight := range r.graph[start] {
		if next == end && maxDistance-weight > 0 {
			routes++
		}

		recRoutes, err := r.RoutesMaxDistance(next, end, maxDistance-weight)
		if err != nil {
			return 0, err
		}
		routes += recRoutes
	}

	return routes, nil
}

func (r *RailRoad) DFSMaxStops(start, end string, maxStops int) (int, error) {
	_, ok := r.graph[start]
	if !ok {
		return 0, ErrNoRoute
	}

	_, ok = r.graph[end]
	if !ok {
		return 0, ErrNoRoute
	}

	if maxStops == 0 {
		return 0, nil
	}

	var routes int

	for next := range r.graph[start] {
		if next == end {
			routes++
			continue
		}

		recRoutes, err := r.DFSMaxStops(next, end, maxStops-1)
		if err != nil {
			return 0, err
		}
		routes += recRoutes
	}

	return routes, nil
}

func (r *RailRoad) DFSExactStops(start, end string, exactStops int) (int, error) {
	_, ok := r.graph[start]
	if !ok {
		return 0, ErrNoRoute
	}

	_, ok = r.graph[end]
	if !ok {
		return 0, ErrNoRoute
	}

	if exactStops == 0 {
		return 0, nil
	}

	var routes int

	for next := range r.graph[start] {
		if next == end && exactStops == 1 {
			routes++
			continue
		}

		recRoutes, err := r.DFSExactStops(next, end, exactStops-1)
		if err != nil {
			return 0, err
		}
		routes += recRoutes
	}

	return routes, nil
}
