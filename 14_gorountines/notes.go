package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	// =-=-=-=-=-==-=-=-=-=-= goroutines =-=-=-=-=-==-=-=-=-=-=
	for i := 0; i < 10; i++ {
		go showMessage(strconv.Itoa(i)) // roda de forma assync (assincrona)
	}
	/*
		nesse caso foi foi imprimido todos os numeros pq a go routine main foi finalizada antes da go rountine showMessage. Como ela é a rotina principal quando ela é finalizada finaliza as outras tb
	*/

	time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))

	// =-=-=-=-=-==-=-=-=-=-= sync.WaitGroup =-=-=-=-=-==-=-=-=-=-=
	/*
		serve para dizer quando os processos em paralelo foram finalizados
	*/
	var wg sync.WaitGroup
	wg.Add(3)

	go callDatabase(&wg)
	go callAPI(&wg)
	go proccesInternal(&wg)

	wg.Wait() // funciona como o sleep só que esse espera os processos terminarem
	fmt.Println()

	// =-=-=-=-=-==-=-=-=-=-= sync.Mutex =-=-=-=-=-==-=-=-=-=-=
	/*
		server para travar uma variavel para que os dados continuem sincronizados (seguindo uma ordem)
	*/

	var m sync.Mutex
	i := 0
	for x := 0; x < 10000; x++ {
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(i)
	println()

	// =-=-=-=-=-==-=-=-=-=-= channel =-=-=-=-=-==-=-=-=-=-=
	channel := make(chan int)
	go setList(channel)

	// valor := <- channel (outra forma de receber o valor de uma channel) <-
	for v := range channel {
		fmt.Println("Recebendo: ", v)
		time.Sleep(time.Second)
	}
	println()

	// =-=-=-=-=-==-=-=-=-=-= channel buffer =-=-=-=-=-==-=-=-=-=-=
	/*
		Channel bufferizado é um channel com um tamanho definido e ele espera o channel encher para depois enviar (geralmente um channel recebe e depois envia, ai segue para o proximo)
	*/
	channel2 := make(chan int, 100)
	go setList(channel2)

	// valor := <- channel (outra forma de receber o valor de uma channel) <-
	for v := range channel2 {
		fmt.Println("Recebendo: ", v)
		time.Sleep(time.Second)
	}
	println()

}

func showMessage(msg string) { // goroutine
	fmt.Println(msg)
}

func callDatabase(wg *sync.WaitGroup) { //sync.WaitGroup
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizado callDatabase")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) { //sync.WaitGroup
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado callAPI")
	wg.Done()
}

func proccesInternal(wg *sync.WaitGroup) { //sync.WaitGroup
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizado proccesInternal")
	wg.Done()
}

func setList(channel chan<- int) { // channel
	for i := 0; i < 10; i++ {
		channel <- i // atribuição para channels <-
		fmt.Println("Enviando: ", i)
	}
	close(channel) // é necessario fechar o canal
}
