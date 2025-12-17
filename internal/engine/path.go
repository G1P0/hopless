package engine

import "github.com/G1P0/hopless/internal/domain"

// FindPath ищет путь From -> To и возвращает список узлов
func FindPath(w domain.World, from, to domain.Node) ([]domain.Node, bool) {
	queue := []domain.Node{from}
	prev := map[domain.Node]domain.Node{}
	visited := map[domain.Node]bool{from: true}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == to {
			break
		}

		for _, l := range w.Links {
			if l.From == current && !visited[l.To] {
				visited[l.To] = true
				prev[l.To] = current
				queue = append(queue, l.To)
			}
		}
	}

	if !visited[to] {
		return nil, false
	}

	// восстановление пути
	path := []domain.Node{to}
	for n := to; n != from; {
		n = prev[n]
		path = append([]domain.Node{n}, path...)
	}

	return path, true
}
