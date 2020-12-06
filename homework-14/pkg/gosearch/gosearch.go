// gosearch - основной пакет приложения
// Функции пакета:
// 1. Создание экзепляра поискового робота со всеми зависимостями
// 2. Инициализация - старт сканирования и индексирования
package gosearch

import (
	"fmt"
	"go-core-own/homework-14/pkg/engine"
	"go-core-own/homework-14/pkg/index"
	"go-core-own/homework-14/pkg/scanner"
	"go-core-own/homework-14/pkg/scanner/spider"
	"go-core-own/homework-14/pkg/storage/btree"
	"sync"
)

// Поисковик GoSearch
type Service struct {
	scanner *spider.Service
	index   *index.Service
	storage *btree.Tree
	Engine  *engine.Service

	sites []string
	depth int
}

// Конструктор поисковика
func New(sites []string, depth int) *Service {
	gs := Service{
		sites: sites,
		depth: depth,

		// Определяются зависимости сканер сайтов, служба индексирования
		scanner: spider.New(&spider.Net{}),
		index:   index.New(),
		storage: btree.New(),
	}

	// Поисковый движок
	gs.Engine = engine.New(gs.index, gs.storage)

	return &gs
}

// запуска скинирования страниц, указанных в конструкторе поисковика
func (gs *Service) Init() {
	// пока идет сканирование движок выдает результаты из загруженных из долговременного хранилища
	go func() {
		const W = 10                         // кол-во "рабочих"
		sites := make(chan string)           // канал для заданий на сканирование
		res := make(chan []scanner.Document) // канал результатов сканирования
		defer close(res)
		err := make(chan error) // канал с ошибками
		defer close(err)

		var wg sync.WaitGroup
		wg.Add(W)

		for i := 0; i < W; i++ {
			go func(sites <-chan string, results chan<- []scanner.Document, err chan<- error) {
				defer wg.Done()
				for s := range sites {
					data, e := gs.scanner.Scan(s, gs.depth)
					if e != nil {
						err <- e
						continue
					}
					results <- data
				}

			}(sites, res, err)
		}
		// поток для записи результатов сканирования
		go func(results <-chan []scanner.Document) {
			for data := range results {
				gs.Engine.Store(data)
			}
		}(res)
		// поток для чтения ошибок
		go func(err <-chan error) {
			for e := range err {
				fmt.Printf("Ошибка при сканировании сайта: %v", e)
			}
		}(err)

		for _, s := range gs.sites {
			sites <- s
		}
		close(sites)
		wg.Wait()
	}()
}
