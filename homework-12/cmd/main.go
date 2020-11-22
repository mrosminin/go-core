package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
С помощью каналов и потоков смоделируем условную игру в пинг-понг.
Игра Пинг-Понг с помощью потоков, каналов и групп ожидания.
Нужно использовать канал для перебрасывания шарика между
половинами стола. Каждому игроку соответствует поток.
Подача начинается командой «begin» в канал. Дальше произвольный
игрок подаёт и пишет в канал "ping". Соперник отвечает "pong".
Эти надписи выводятся на экран вместе с именем игрока,
совершившего удар.
В 20% случаев любому игроку во время удара может повезти и он
мощно загасит шарик в самый угол и выиграет очко.
В этом случае в канал подаётся сигнал "stop", означающий завершение
подачи.
В конце партии на экран печатается счёт
*/

type player struct {
	name  string
	score int
	sync  chan struct{}
}

func (p *player) play(g *game, opponent *player) {
	for range p.sync {
		g.mux.Lock()
		if g.state == "" || g.state == "pong" {
			g.state = "ping"
		} else {
			g.state = "pong"
		}
		fmt.Println(p.name + " " + g.state)
		g.mux.Unlock()

		if accuracy := rand.Intn(10); accuracy > 7 {
			fmt.Println("Score!!!")
			p.score++
			if p.score > 10 {
				g.end <- struct{}{}
				return
			}
			g.ch <- "stop"
		} else {
			opponent.sync <- struct{}{}
		}
	}
}

type game struct {
	state string
	mux   sync.Mutex

	player1 *player
	player2 *player
	ch      chan string
	end     chan struct{}
}

func (g *game) start() {
	for {
		val := <-g.ch
		if val == "begin" {
			go g.player1.play(g, g.player2)
			go g.player2.play(g, g.player1)
			g.serve()

		}
		if val == "stop" {
			g.state = ""
			g.serve()
		}
	}
}

func (g *game) serve() {
	if order := rand.Intn(2); order == 0 {
		g.player1.sync <- struct{}{}
	} else {
		g.player2.sync <- struct{}{}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	g := game{
		player1: &player{
			name: "Player1",
			sync: make(chan struct{}),
		},
		player2: &player{
			name: "Player2",
			sync: make(chan struct{}),
		},
		end: make(chan struct{}),
		ch:  make(chan string),
	}
	go g.start()
	g.ch <- "begin"

	<-g.end
	fmt.Printf("%s %d : %d %s", g.player1.name, g.player1.score, g.player2.score, g.player2.name)

}
