package main

import "sort"

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

// метод для расчёта рейтинга
func (p *Player) calculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals) + float64(p.Assists)/2
	} else {
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2) / float64(p.Misses)
	}
}

// сортировка по количеству голов
func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals == players[j].Goals {
			return players[i].Name < players[j].Name
		}
		return players[i].Goals > players[j].Goals
	})
	return players
}

// сортировка по рейтингу
func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Rating == players[j].Rating {
			return players[i].Name < players[j].Name
		}
		return players[i].Rating > players[j].Rating
	})
	return players
}

// сортировка по отношению голов к промахам
func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		// вычисляем отношение для игрока i
		ratioI := float64(players[i].Goals)
		if players[i].Misses > 0 {
			ratioI /= float64(players[i].Misses)
		}

		// вычисляем отношение для игрока j
		ratioJ := float64(players[j].Goals)
		if players[j].Misses > 0 {
			ratioJ /= float64(players[j].Misses)
		}

		// сравнение по отношению, при равенстве — по имени
		if ratioI == ratioJ {
			return players[i].Name < players[j].Name
		}
		return ratioI > ratioJ
	})
	return players
}

func NewPlayer(name string, goals, misses, assists int) Player {
	nPlayer := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}
	nPlayer.calculateRating()
	return nPlayer
}
