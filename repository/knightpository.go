package repository

import (
	"chess-move-api/domain/entities"

	"errors"
	"strconv"
)
var (
	xyMax = 8
	xyMin = 1
)
type Knightrepository struct {

}

func (repo *Knightrepository) GetNextsMoves(position string) (entities.NextMoves, error) {
	nextMoves := entities.NextMoves { }
	if position == "" {
		return nextMoves, errors.New("position is empty")
	}
	if position[0] < 'A' || position[0] > 'H' {
		return nextMoves, errors.New("Invalid position: ("+position+")")
	}
	sY := string(position[1])
	Y, err := strconv.Atoi(sY)
	if err != nil {
		return nextMoves, errors.New("Invalid position: ("+position+")")
	}
	if Y < 1 || Y > 8 {
		return nextMoves, errors.New("Invalid position: ("+position+")")
	}

	X := int(position[0] - 64)

	deslocamento := 1
	offset := 3

	positions := make(map[string]string)

	//HORIZONTAL => RIGHT
	for deslocamento <= 2 {
		sY = ""
		offset--

		basePos := X + deslocamento
		if basePos > xyMax {
			deslocamento++
			continue
		}
		var b byte = byte(64 + basePos)
		if basePos <= xyMax {
			sY = string(b) + strconv.Itoa(Y+offset)
		}
		if sY != "" {
			if _, ok := positions[sY]; !ok {
				positions[sY] = sY
			}
		}
		if basePos >= xyMin {
				sY = string(b) + strconv.Itoa(Y-offset)
		}
		if sY != "" {
			if _, ok := positions[sY]; !ok {
				positions[sY] = sY
			}
		}
		deslocamento++
	}
	////////////////////////////////////////////////////////////////////////////

	deslocamento = 1
	offset = 3

	//HORIZONTAL => LEFT
	for deslocamento <= 2 {
		sY = ""
		offset--

		basePos := X - deslocamento

		if basePos < xyMin {
			deslocamento++
			continue
		}
		var b byte = byte(64 + basePos)
		if (basePos <= xyMax) && ((Y+offset) <= xyMax) {
			sY = string(b) + strconv.Itoa(Y+offset)
		} 
		if sY != "" {
			if _, ok := positions[sY]; !ok {
				positions[sY] = sY
			}
		}
		if (basePos >= xyMin) && ((Y-offset) >= xyMin) {
			sY = string(b) + strconv.Itoa(Y-offset)
		}
		if sY != "" {
			if _, ok := positions[sY]; !ok {
				positions[sY] = sY
			}
		}
		deslocamento++
	}
	////////////////////////////////////////////////////////////////////////////
	deslocamento = 1
	offset = 3

	//VERTICAL => UP
	for deslocamento <= 2 {
		sY = ""
		offset--

		basePos := Y + deslocamento

		if basePos > xyMax {
			deslocamento++
			continue
		}
		var b byte = byte(64 + X - offset)
		if basePos >= xyMin {
			sY = string(b) + strconv.Itoa(basePos)
		} 
		if sY != "" {
			if _, ok := positions[sY]; !ok {
				positions[sY] = sY
			}
		}
		if basePos >= xyMax {
			b = byte(64 + X + offset)
			sY = string(b) + strconv.Itoa(basePos)
		}
		if sY != "" {
			if _, ok := positions[sY]; !ok {
				positions[sY] = sY
			}
		}
		deslocamento++
	}


	////////////////////////////////////////////////////////////////////////////
	deslocamento = 1
	offset = 3

	//VERTICAL => DOWN
	for deslocamento <= 2 {
		sY = ""
		offset--

		basePos := Y - deslocamento

		if basePos < xyMin {
			deslocamento++
			continue
		}
		var b byte = byte(64 + X - offset)
		if basePos >= xyMin {
			sY = string(b) + strconv.Itoa(basePos)
		}
		if sY != "" {
			if _, ok := positions[sY]; !ok {
				positions[sY] = sY
			}
		}
		if basePos <= xyMax {
			b = byte(64 + X + offset)
			sY = string(b) + strconv.Itoa(basePos)
		}
		if sY != "" {
			if _, ok := positions[sY]; !ok {
				positions[sY] = sY
			}
		}
		deslocamento++
	}
	////////////////////////////////////////////////////////////////////////////
	if len(positions) <= 0 {
		return nextMoves, errors.New("no moves for: ("+position+")")
	}
	arr := []string{}
	for _, v := range positions {
		arr = append(arr, v)
	}
	nextMoves.Moves = arr
	
	return nextMoves, nil
}

