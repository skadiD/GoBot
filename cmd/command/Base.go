package command

import (
	"github.com/fexli/logger"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"gobot/cmd/structs"
)

// Send 发送请求
func Send(CommandName string, data interface{}, conn *websocket.Conn) {
	logger.RootLogger.Debug(logger.WithContent(CommandName, data))
	_base := structs.SendBase{
		SyncId:     -1,
		Command:    CommandName,
		SubCommand: nil,
		Content:    data,
	}
	go func() {
		//conn.WriteJSON(_base)
		w, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			//return err
		}
		_ = json.NewEncoder(w).Encode(_base)
		_ = w.Close()
		//payload, _ := json.Marshal(_base)
		//conn.WriteMessage(websocket.TextMessage, payload)
	}()
}
