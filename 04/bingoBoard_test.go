package main

import (
	"reflect"
	"testing"
)

func TestBoardRows(t *testing.T) {
	arr1 := []int{0, 1, 2, 3}
	expectArr1 := [][]Cell{
		{Cell{v: 0}, Cell{v: 1}},
		{Cell{v: 2}, Cell{v: 3}},
	}
	board := NewBoard(arr1, 2)
	rows := board.rows()
	t.Log(rows)
	if !reflect.DeepEqual(expectArr1, rows) {
		t.Fatal()
	}
}

func TestBoardColumns(t *testing.T) {
	arr1 := []int{0, 1, 2, 3}
	expectArr1 := [][]Cell{
		{Cell{v: 0}, Cell{v: 2}},
		{Cell{v: 1}, Cell{v: 3}},
	}
	board := NewBoard(arr1, 2)
	rows := board.columns()
	if !reflect.DeepEqual(expectArr1, rows) {
		t.Fatal()
	}
}

func TestBoardMarkCell(t *testing.T) {
	arr := []int{0, 1, 2, 3}
	board := NewBoard(arr, 2)
	board.markCell(0)
	expectArr1 := [][]Cell{
		{Cell{v: 0, checked: true}, Cell{v: 1}},
		{Cell{v: 2}, Cell{v: 3}},
	}
	if !reflect.DeepEqual(expectArr1, board.rows()) {
		t.Fatal("markCell to 0 is broken")
	}

	board.markCell(99)
	if !reflect.DeepEqual(expectArr1, board.rows()) {
		t.Fatal("markCell to 99 is broken")
	}
}

func TestBoardBingo(t *testing.T) {
	arr := []int{0, 1, 2, 3}
	board := NewBoard(arr, 2)
	if board.isBingo() {
		t.Fatal("false bingo: 0 marks")
	}
	board.markCell(0)
	if board.isBingo() {
		t.Fatal("false bingo: 1 marks")
	}
	board.markCell(1)
	if !board.isBingo() {
		t.Fatal("missed bingo: 2 marks")
	}

	arr1 := []int{0, 1, 2, 3}
	board = NewBoard(arr1, 2)
	if board.isBingo() {
		t.Fatal("false bingo: 0 marks")
	}
	board.markCell(0)
	if board.isBingo() {
		t.Fatal("false bingo: 1 marks")
	}
	board.markCell(2)
	if !board.isBingo() {
		t.Fatal("missed bingo: 2 marks")
	}

	arr2 := []int{0, 1, 2, 3}
	board = NewBoard(arr2, 2)
	if board.isBingo() {
		t.Fatal("false bingo: 0 marks")
	}
	board.markCell(0)
	if board.isBingo() {
		t.Fatal("false bingo: 1 marks")
	}
	board.markCell(3)
	if board.isBingo() {
		t.Fatal("false bingo: 2 marks")
	}
	board.markCell(2)
	if !board.isBingo() {
		t.Fatal("missed bingo: 3 marks")
	}
}
