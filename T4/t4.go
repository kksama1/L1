package main

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и выводят в
// stdout. Необходима возможность выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Воркер выводящий значение, полученно из канала в stdout.
// При получении данных с канала done выходит функции тем самым закрывая горутину.
func worker(id int) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case data, ok := <-dataChannel:
			if !ok {
				return
			}
			time.Sleep(1 * time.Second)
			fmt.Printf("Воркер %d: %s\n", id, data)
		}
	}
}

var (
	numWorkers  int
	dataChannel chan string
	done        chan struct{}
	wg          sync.WaitGroup
)

func main() {
	// Задаем количество воркеров
	fmt.Printf("Введите количесво воркеров: ")
	fmt.Scan(&numWorkers)

	// Создаем канал для обмена данными
	dataChannel = make(chan string)

	// Создаем канал для завершения программы
	done = make(chan struct{})

	// Запускаем воркеры
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Обработка сигнала Ctrl+C для завершения программы
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\nПолучен сигнал завершения")
		close(done)
	}()

	// Главный поток записывает данные в канал
	for {
		select {
		case <-done:
			// Программа завершена, закрываем канал
			close(dataChannel)
			close(sigCh)
			wg.Wait()
			fmt.Println("\nБезопасный выход")
			return
		default:
			i := rand.Intn(500)
			dataChannel <- fmt.Sprintf("Получил данные %d", i)
		}
	}
}
