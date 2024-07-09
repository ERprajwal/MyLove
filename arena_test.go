package main

import (
	"testing"
)

func TestPlayerSurvival(t *testing.T) {
	player := &Player{Name: "Test Player", Health: 10, Strength: 5, Attack: 5}
	if !player.IsAlive() {
		t.Errorf("Expected player to be alive, but they are not")
	}
	player.Health = 0
	if player.IsAlive() {
		t.Errorf("Expected player to be dead, but they are alive")
	}
}

func TestRollDice(t *testing.T) {
	arena := NewArena(&Player{}, &Player{})
	for i := 0; i < 100; i++ {
		roll := arena.RollDice()
		if roll < 1 || roll > 6 {
			t.Errorf("Expected dice roll to be between 1 and 6, but got %d", roll)
		}
	}
}

func TestAttackLogic(t *testing.T) {
	playerA := &Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
	playerB := &Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}
	arena := NewArena(playerA, playerB)
	arena.attack(playerA, playerB)
	if playerB.Health == 100 {
		t.Errorf("Expected Player B's health to decrease, but it did not")
	}
}
