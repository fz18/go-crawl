package persist

import "log"

func ItemSave() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item save Got $%d, %v", itemCount, item)
			itemCount++
		}
	}()

	return out
}
