package websocket

import (
	"errors"
)

func Send(msg map[string]interface{}) (error) {

	select{
	case Hub.Broadcast <- msg:
		break
	default:
		return errors.New("unable to send message")
	}
	return nil
}