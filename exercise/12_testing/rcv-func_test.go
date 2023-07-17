package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}
}

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.addHealth(999)
	if player.health > player.maxHealth {
		t.Fatalf("Health went over limit: %v, want: %v", player.health, player.maxHealth)
	}
	player.applyDamage(player.maxHealth + 1)
	if player.health < 0 {
		t.Fatalf("Health: %v. Minimum 0", player.health)
	}
	if player.health > player.maxHealth {
		t.Fatalf("Health: %v. Maximum: %v", player.health, player.maxHealth)
	}
}

func TestEnergy(t *testing.T) {
	player := newPlayer()
	player.addEnergy(999)
	if player.energy > player.maxEnergy {
		t.Fatalf("Energy went over limit: %v, want: %v", player.energy, player.maxEnergy)
	}
	player.applyDamage(player.maxEnergy + 1)
	if player.energy < 0 {
		t.Fatalf("Energy: %v. Minimum 0", player.energy)
	}
	if player.energy > player.maxEnergy {
		t.Fatalf("Energy: %v. Maximum: %v", player.energy, player.maxEnergy)
	}
}
