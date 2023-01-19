package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 * Q mokymasis programa GO kalba
 * Pradineje matricoje naudojama -1 nurodyti sienoms
 * Mokymosi metu, kitam zingsniui teikiama pirmenybe langeliui, kuris buvo maziausia kartu lankytas
 */
var visited = [9][12]int{
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	{-1, 0, -1, 0, 0, 0, -1, 0, -1, 0, 0, -1},
	{-1, 0 - 1, 0, 0, -1, -1, 0, -1, 0, 0, -1},
	{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1},
	{-1, 0, -1, -1, -1, -1, -1, 0, -1, -1, 0, -1},
	{-1, 0, 0, 0, -1, 0, 0, 0, -1, 0, 0, -1},
	{-1, 0, 0, 0, -1, 0, -1, -1, -1, -1, 0, -1},
	{-1, 0, 0, 0, -1, 0, 0, 0, 0, 0, 0, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
}

func nextStep(R [9][12]int, cx int, cy int) (int, int) {
	// maziausiai aplankytam langeliui teikiama pirmenybe mokymosi metu
	visits := 99999
	x, y := cx, cy

	if R[cy][cx+1] != -1 {
		if visited[cy][cx+1] < visits {
			visits = visited[cy][cx+1]
			x, y = cx+1, cy
		}
	}
	if R[cy+1][cx] != -1 {
		if visited[cy+1][cx] < visits {
			visits = visited[cy+1][cx]
			x, y = cx, cy+1
		}
	}
	if R[cy][cx-1] != -1 {
		if visited[cy][cx-1] < visits {
			visits = visited[cy][cx-1]
			x, y = cx-1, cy
		}
	}
	if R[cy-1][cx] != -1 {
		if visited[cy-1][cx] < visits {
			x, y = cx, cy-1
		}
	}

	visited[y][x]++
	return x, y
}

// randamas kaimyninis langelis su auksciausia Q reiksme
func maxQValue(Q [9][12]float32, x int, y int) float32 {
	var max float32 = 0
	if max < Q[y][x] {
		max = Q[y][x]
	}
	if max < Q[y+1][x] {
		max = Q[y+1][x]
	}
	if max < Q[y][x+1] {
		max = Q[y][x+1]
	}
	if max < Q[y][x-1] {
		max = Q[y][x-1]
	}
	if max < Q[y-1][x] {
		max = Q[y-1][x]
	}
	return max
}

// apskaiciuojama Q reiksme
func calcQ(R [9][12]int, Q [9][12]float32, cx int, cy int, x int, y int) float32 {
	return float32(R[cy][cx]) + 0.7*maxQValue(Q, x, y)
}

// Q-learning lenteles Q uzpildymas
func learn(R [9][12]int, Q [9][12]float32, x int, y int, tx int, ty int, steps int) [9][12]float32 {
	steps++
	next_x, next_y := nextStep(R, x, y)
	Q[y][x] = calcQ(R, Q, x, y, next_x, next_y)
	if x == tx && y == ty {
		fmt.Printf("Pasiektas tikslas per %5d zingsniu\n", steps)
		return Q
	}
	return learn(R, Q, next_x, next_y, tx, ty, steps)
}

// parenkama geriausia kryptis (kryptis su auksciausia Q reiksme)
func bestDirection(Q [9][12]float32, x int, y int) (int, int) {
	var max float32 = 0
	bx, by := x, y
	if max <= Q[y+1][x] {
		max = Q[y+1][x]
		bx, by = x, y+1
	}
	if max <= Q[y][x+1] {
		max = Q[y][x+1]
		bx, by = x+1, y
	}
	if max <= Q[y][x-1] {
		max = Q[y][x-1]
		bx, by = x-1, y
	}
	if max <= Q[y-1][x] {
		bx, by = x, y-1
	}
	return bx, by
}

// randamas greiciausias kelias
func shortestPath(Q [9][12]float32, sx int, sy int, tx int, ty int, steps int, verbose bool) int {
	if sx == tx && sy == ty {
		return steps
	}
	steps++
	x, y := bestDirection(Q, sx, sy)
	if verbose {
		fmt.Printf("%5d => [ %2d; %2d]\n", steps, x, y)
	}
	return shortestPath(Q, x, y, tx, ty, steps, verbose)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var R = [9][12]int{
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, 0, -1, 0, 0, 0, -1, 0, -1, 0, 0, -1},
		{-1, 0, -1, 0, 0, -1, -1, 0, -1, 0, 0, -1},
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -1},
		{-1, 0, -1, -1, -1, -1, -1, 10, -1, -1, 10, -1},
		{-1, 0, 0, 0, -1, 10, 0, 0, -1, 0, 0, -1},
		{-1, 0, 0, 0, -1, 25, -1, -1, -1, -1, 10, -1},
		{-1, 0, 0, 0, -1, 50, 75, 100, 75, 50, 25, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	}
	tx, ty := 7, 7
	var Q [9][12]float32
	fmt.Printf("Tikslo langelis: [ %3d ; %3d ]\n", tx, ty)

	for x := 1; x < 11; x++ {
		for y := 1; y < 8; y++ {
			if R[y][x] != -1 {
				fmt.Printf("Pradinis langelis: [ %3d ; %3d ]\n", x, y)
				Q = learn(R, Q, x, y, tx, ty, 0)
			}
		}
	}

	fmt.Printf("\n<--- Galutine Q lentele --->\n")
	for i := 0; i < 9; i++ {
		for j := 0; j < 12; j++ {
			if Q[i][j] > 0 {
				fmt.Printf(" %4.0f ", Q[i][j])
			} else {
				fmt.Printf(" %4s ", "")
			}
		}
		fmt.Printf("\n\n")
	}

	for x := 1; x < 11; x++ {
		for y := 1; y < 8; y++ {
			if R[y][x] != -1 {
				fmt.Printf("Trumpiausias kelias is [ %3d ; %3d ] i [ %3d ; %3d ] = %3d zingsniu\n", x, y, tx, ty, shortestPath(Q, x, y, tx, ty, 0, false))
			}
		}
	}
}
