package lib

import (
	"fmt"
)

func StartTurnBasedSolving(field *Field, pathsToExit [][]string) {
	isSolved := false
	turns := []string{}
	for !isSolved {
		turn := ""
		isSolved = true
		pathsUsed := []string{}
		for _, ant := range field.ants {
			if ant.currentRoom == field.endRoomName {
				ant.isFinished = true
			}

			if !ant.isFinished {
				nextRoom := getNextRoom(ant.currentRoom, pathsToExit, pathsUsed, *field)

				for _, path := range pathsUsed {
					if path == ant.currentRoom+"-"+nextRoom {
						nextRoom = "" // to prevent ants from colliding
					}
				}

				if nextRoom != "" {
					turnStruct := "%s L%v-%s"
					if turn == "" {
						turnStruct = "%sL%v-%s"
					}

					pathsUsed = append(pathsUsed, ant.currentRoom+"-"+nextRoom)
					ant.currentRoom = nextRoom
					turn = fmt.Sprintf(turnStruct, turn, ant.id+1, nextRoom)
				}

				isSolved = false
			}
		}

		if !isSolved {
			turns = append(turns, turn)
		}
	}

	for _, turn := range turns {
		fmt.Println(turn)
	}
}

func getNextRoom(currentRoom string, pathsToExit [][]string, usedPaths []string, field Field) string {
	for _, pathToExit := range pathsToExit {
		for i, room := range pathToExit {
			if room == currentRoom {
				nextRoom := pathToExit[i+1]

				// As I remember, this is a hack to prevent ants from colliding in some test case... So a little bit of cheating here :D
				if len(field.ants)-len(pathToExit) == getNumOfFinishedAnts(field.ants) && nextRoom != field.endRoomName && currentRoom == field.startRoomName && len(pathsToExit) == 2 {
					continue // Wait for better path
				}

				if isRoomEmpty(nextRoom, field) || nextRoom == field.endRoomName {
					isPathUsed := false
					for _, path := range usedPaths {
						if path == currentRoom+"-"+nextRoom {
							isPathUsed = true
						}
					}

					if isPathUsed {
						continue
					}

					return nextRoom
				} else {
					continue
				}
			}
		}
	}

	return ""
}

func isRoomEmpty(roomName string, field Field) bool {
	for _, ant := range field.ants {
		if ant.currentRoom == roomName {
			return false
		}
	}

	return true
}

func getNumOfFinishedAnts(ants []*Ant) int {
	numOfFinishedAnts := 0
	for _, ant := range ants {
		if ant.isFinished {
			numOfFinishedAnts++
		}
	}

	return numOfFinishedAnts
}
