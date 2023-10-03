package main

type IGun interface {
	setCategory(category string)
	getCategory() string
	setName(name string)
	getName() string
	setMagazineBullets(bullets uint8)
	getMagazineBullets() uint8
	setStockBullets(bullets uint8)
	getStockBullets() uint8
}

type Gun struct {
	category        string
	name            string
	magazineBullets uint8
	stockBullets    uint8
}

func (g *Gun) setCategory(name string) {
	g.category = name
}

func (g *Gun) getCategory() string {
	return g.category
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setMagazineBullets(bullets uint8) {
	g.magazineBullets = bullets
}

func (g *Gun) getMagazineBullets() uint8 {
	return g.magazineBullets
}

func (g *Gun) setStockBullets(bullets uint8) {
	g.stockBullets = bullets
}

func (g *Gun) getStockBullets() uint8 {
	return g.stockBullets
}
