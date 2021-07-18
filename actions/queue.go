package actions

import (
	"time"
)

var FreeSlot int
var queue []string

func GetQueue(messageId string) {
	queue = append(queue, messageId)

	for {
		if FreeSlot < 1 {
			if len(queue) > 0 && queue[0] == messageId {
				queue = queue[1:]
				FreeSlot = FreeSlot + 1
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
}
