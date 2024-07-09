package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player struct represents a player in the arena.
type Player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}

// IsAlive checks if the player is still alive.
func (p *Player) IsAlive() bool {
	return p.Health > 0
}

// Arena struct represents the arena where players fight.
type Arena struct {
	PlayerA *Player
	PlayerB *Player
	rand    *rand.Rand
}

// NewArena creates a new arena with two players.
func NewArena(playerA, playerB *Player) *Arena {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return &Arena{PlayerA: playerA, PlayerB: playerB, rand: r}
}

// RollDice simulates rolling a six-sided die.
func (a *Arena) RollDice() int {
	return a.rand.Intn(6) + 1
}

// Fight simulates the fight between two players until one dies.
func (a *Arena) Fight() {
	for a.PlayerA.IsAlive() && a.PlayerB.IsAlive() {
		a.singleRound()
	}
}

// singleRound simulates a single round of attack and defense between players.
func (a *Arena) singleRound() {
	if a.PlayerA.Health <= a.PlayerB.Health {
		a.attack(a.PlayerA, a.PlayerB)
		if a.PlayerB.IsAlive() {
			a.attack(a.PlayerB, a.PlayerA)
		}
	} else {
		a.attack(a.PlayerB, a.PlayerA)
		if a.PlayerA.IsAlive() {
			a.attack(a.PlayerA, a.PlayerB)
		}
	}
}

// attack simulates one player attacking another player.
func (a *Arena) attack(attacker, defender *Player) {
	attackRoll := a.RollDice()
	defendRoll := a.RollDice()
	attackDamage := attacker.Attack * attackRoll
	defendStrength := defender.Strength * defendRoll
	damage := attackDamage - defendStrength
	if damage > 0 {
		defender.Health -= damage
		if defender.Health < 0 {
			defender.Health = 0
		}
	}
	fmt.Printf("%s attacks %s: %d attack damage, %d defend strength, %d damage dealt, %s health: %d\n",
		attacker.Name, defender.Name, attackDamage, defendStrength, damage, defender.Name, defender.Health)
}

// main function to set up players and start the fight.
func main() {
	playerA := &Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
	playerB := &Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}

	arena := NewArena(playerA, playerB)
	arena.Fight()

	if playerA.IsAlive() {
		fmt.Println("Player A wins!")
	} else {
		fmt.Println("Player B wins!")
	}
}

/*
import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}

func (p *Player) IsAlive() bool {
	return p.Health > 0
}

type Arena struct {
	PlayerA *Player
	PlayerB *Player
}

func NewArena(playerA, playerB *Player) *Arena {
	return &Arena{PlayerA: playerA, PlayerB: playerB}
}

func (a *Arena) RollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func (a *Arena) Fight() {
	for a.PlayerA.IsAlive() && a.PlayerB.IsAlive() {
		a.singleRound()
	}
}

func (a *Arena) singleRound() {
	if a.PlayerA.Health <= a.PlayerB.Health {
		a.attack(a.PlayerA, a.PlayerB)
		if a.PlayerB.IsAlive() {
			a.attack(a.PlayerB, a.PlayerA)
		}
	} else {
		a.attack(a.PlayerB, a.PlayerA)
		if a.PlayerA.IsAlive() {
			a.attack(a.PlayerA, a.PlayerB)
		}
	}
}

func (a *Arena) attack(attacker, defender *Player) {
	attackRoll := a.RollDice()
	defendRoll := a.RollDice()
	attackDamage := attacker.Attack * attackRoll
	defendStrength := defender.Strength * defendRoll
	damage := attackDamage - defendStrength
	if damage > 0 {
		defender.Health -= damage
		if defender.Health < 0 {
			defender.Health = 0
		}
	}
	fmt.Printf("%s attacks %s: %d attack damage, %d defend strength, %d damage dealt, %s health: %d\n",
		attacker.Name, defender.Name, attackDamage, defendStrength, damage, defender.Name, defender.Health)
}

func main() {
	playerA := &Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
	playerB := &Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}

	arena := NewArena(playerA, playerB)
	arena.Fight()

	if playerA.IsAlive() {
		fmt.Println("Player A wins!")
	} else {
		fmt.Println("Player B wins!")
	}
}
*/
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type CountryResponse struct {
// 	Data []struct {
// 		Name    string `json:"name"`
// 		Capital string `json:"capital"`
// 	} `json:"data"`
// }

// func getCapitalCity(country string) string {
// 	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/countries?name=%s", country)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "-1"
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "-1"
// 	}

// 	var countryResponse CountryResponse
// 	err = json.Unmarshal(body, &countryResponse)
// 	if err != nil {
// 		return "-1"
// 	}

// 	if len(countryResponse.Data) == 0 {
// 		return "-1"
// 	}

// 	return countryResponse.Data[0].Capital
// }

// func main() {
// 	country := "Italy"
// 	capital := getCapitalCity(country)
// 	fmt.Println(capital)
// }
