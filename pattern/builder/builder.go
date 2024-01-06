package main

import "fmt"

type Player struct {
	name         string
	health       int
	gunType      string
	ammo         int
	ammoToReload int
}

type PlayerBuilder struct {
	player *Player
}

type PlayerInfoBuilder struct {
	PlayerBuilder
}

type PlayerGunBuilder struct {
	PlayerBuilder
}

func NewPlayerBuilder() *PlayerBuilder {
	return &PlayerBuilder{player: &Player{}}
}

func (p *PlayerBuilder) PlayerInfo() *PlayerInfoBuilder {
	return &PlayerInfoBuilder{*p}
}

func (p *PlayerBuilder) GunInfo() *PlayerGunBuilder {
	return &PlayerGunBuilder{*p}
}

func (p *PlayerInfoBuilder) Name(name string) *PlayerInfoBuilder {
	p.player.name = name
	return p
}

func (p *PlayerInfoBuilder) Health(health int) *PlayerInfoBuilder {
	p.player.health = health
	return p
}

func (p *PlayerGunBuilder) GunType(gunType string) *PlayerGunBuilder {
	p.player.gunType = gunType
	return p
}

func (p *PlayerGunBuilder) Ammo(ammo int) *PlayerGunBuilder {
	p.player.ammo = ammo
	return p
}

func (p *PlayerGunBuilder) AmmoToReload(ammoToReload int) *PlayerGunBuilder {
	p.player.ammoToReload = ammoToReload
	return p
}

func (p *PlayerBuilder) Build() *Player {
	return p.player
}

func main() {
	b := NewPlayerBuilder()
	player := b.PlayerInfo().
		Name("Ivan").
		Health(100).
		GunInfo().
		GunType("Ak-47").
		Ammo(300).
		AmmoToReload(35).
		Build()

	fmt.Print(player)
}
