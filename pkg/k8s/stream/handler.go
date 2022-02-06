package stream

import (
	"encoding/json"
	"fmt"
	"github.com/gitctl-pro/gitctl/pkg/k8s"
	"github.com/gorilla/websocket"
	"io"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

const END_OF_TRANSMISSION = "\u0004"

// TerminalMessage is the messaging protocol between ShellController and TerminalSession.
//
// OP      DIRECTION  FIELD(S) USED  DESCRIPTION
// ---------------------------------------------------------------------
// bind    fe->be     SessionID      Id sent back from TerminalResponse
// stdin   fe->be     Data           Keystrokes/paste buffer
// resize  fe->be     Rows, Cols     New terminal size
// stdout  be->fe     Data           Output from the process
// toast   be->fe     Data           OOB message to be shown to the user
type TerminalMessage struct {
	Op, Data, SessionID string
	Rows, Cols          uint16
}

type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

type StreamHandler struct {
	WsConn     *websocket.Conn
	ResizeChan chan remotecommand.TerminalSize
	doneChan   chan struct{}
	inputCh    chan []byte
	outputCh   chan []byte
	width      uint16
	height     uint16
}

func BindStreamHandler(cfg *rest.Config, namespace, name, container, command string) {
	executor, err := ExecCommad(cfg, namespace, name, container, command)
	if err != nil {

	}
	handler := &StreamHandler{
		ResizeChan: make(chan remotecommand.TerminalSize),
	}
	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		Tty:               true,
		TerminalSizeQueue: handler,
	})
	if err != nil {

	}
}

func ExecCommad(cfg *rest.Config, namespace, name, container, command string) (remotecommand.Executor, error) {
	client, resource, _ := k8s.KindForResource(cfg, &schema.GroupVersionKind{
		Kind:    "pod",
		Version: "v1",
	})
	req := client.Post().
		Resource(resource).
		Name(name).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Container: container,
		Command:   []string{command},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(cfg, "POST", req.URL())
	return exec, err
}

func (handler *StreamHandler) Write(p []byte) (size int, err error) {
	msg, err := json.Marshal(TerminalMessage{
		Op:   "stdout",
		Data: string(p),
	})
	if err != nil {
		return 0, err
	}
	err = handler.WsConn.WriteMessage(websocket.TextMessage, msg)
	return len(p), nil
}

func (handler *StreamHandler) Next() (size *remotecommand.TerminalSize) {
	select {
	case size := <-handler.ResizeChan:
		return &size
	case <-handler.doneChan:
		return nil
	}
}

func (handler *StreamHandler) Read(p []byte) (size int, err error) {
	_, buf, err := handler.WsConn.ReadMessage()
	var msg TerminalMessage
	if err := json.Unmarshal([]byte(buf), &msg); err != nil {
		return copy(p, END_OF_TRANSMISSION), err
	}

	switch msg.Op {
	case "bind":
		//log.Println(msg)
	case "stdin":
		return copy(p, msg.Data), nil
	case "resize":
		handler.ResizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	default:
		return copy(p, END_OF_TRANSMISSION), fmt.Errorf("unknown message type '%s'", msg.Op)
	}
	return
}

func (handler *StreamHandler) Close() error {
	close(handler.outputCh)
	close(handler.inputCh)
	return nil
}
