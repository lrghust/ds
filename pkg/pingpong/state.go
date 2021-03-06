package pingpong

import (
	"project/pkg/base"
)

func IsFinal(state_ *base.State) bool {
	for _, node := range state_.Nodes() {
		client, ok := node.(*Client)
		if !ok {
			continue
		}

		count := client.Attribute().(int)
		return count == 5
	}
	return false
}
