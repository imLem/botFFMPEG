package actions

import (
	"fmt"
	"time"
)

var FreeSlot int
var queue []string

func GetQueue(messageId string) {
	queue = append(queue, messageId)

	for {
		if FreeSlot < 1 {
			if len(queue) > 0 && queue[0] == messageId {
				fmt.Println(queue[0] + " пошел в обработку")
				queue = queue[1:]
				FreeSlot = FreeSlot + 1
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
}
