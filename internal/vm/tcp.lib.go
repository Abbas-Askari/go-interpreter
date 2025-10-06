package vm

import (
	"Abbas-Askari/interpreter-v2/internal/object"
	"Abbas-Askari/interpreter-v2/internal/op"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func getMessagePrinterClosure(message string, arity int) object.Closure {
	return object.NewClosure(
		object.NewFunction(arity, "printMessage", "internal", []op.OpCode{
			op.OpConstant, 0,
			op.OpPrint,
			op.OpNil,
			op.OpReturn,
		}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []object.Object{
			object.NewString(message),
		}))
}

func NewSocketObject(conn net.Conn, vm *VM) object.Map {
	vm.RegisterEvent()

	socket := object.Map{Map: map[string]object.Object{}}
	socket.Map["isOpen"] = object.Boolean{true}
	socket.Map["onData"] = getMessagePrinterClosure("Warning: onData was called without an implementation", 1)
	socket.Map["onError"] = getMessagePrinterClosure("Warning: onError was called without an implementation", 1)
	socket.Map["onEnd"] = getMessagePrinterClosure("Warning: onEnd was called without an implementation", 0)

	done := make(chan struct{})

	closeSocket := func() {
		close(done)
		socket.Map["isOpen"] = object.Boolean{false}
		vm.FireEvent(socket.Map["onEnd"].(object.Closure))
		vm.DetachEvent()
		conn.Close()
	}

	socket.Map["close"] = NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			if socket.Map["isOpen"] == (object.Boolean{false}) {
				vm.runtimeError("Socket is already closed")
				return object.Nil{}
			}
			closeSocket()
			return object.Nil{}
		},
		Arity: 0,
		Name:  "close",
	}

	socket.Map["write"] = NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			buf := bytes.NewBuffer([]byte{})
			data, ok := args[0].(object.String)
			if ok {
				buf.WriteString(data.Value)
			} else if data, ok := args[0].(object.Number); ok {
				binary.Write(buf, binary.BigEndian, int32(data.Value))
			} else {
				vm.assertArgumentToType(args[0], object.STRING, "write", 0)
			}
			// go func() {
			_, err := conn.Write(buf.Bytes())
			if err != nil {
				vm.FireEvent(socket.Map["onError"].(object.Closure), object.NewString(err.Error()))
			}
			// }()
			return object.Nil{}
		},
		Arity: 1,
		Name:  "write",
	}

	buffer := make([]byte, 1024)

	readerRoutine := func() {
		for {
			select {
			case <-done:
				return
			default:
				n, err := conn.Read(buffer)
				if err != nil {
					if err != io.EOF {
						vm.FireEvent(socket.Map["onError"].(object.Closure), object.NewString(err.Error()))
						continue
					}
					closeSocket()
					return
				}
				data := make([]byte, n)
				copy(data, buffer[:n])
				// arr := make([]object.Object, n)
				// for i := 0; i < n; i++ {
				// 	arr[i] = object.Number{float64(data[i])}
				// }
				// obj := object.NewArray(arr)
				obj := object.NewBuffer(data)
				vm.FireEvent(socket.Map["onData"].(object.Closure), obj)
			}
		}
	}

	go readerRoutine()

	return socket
}

func getTCP() *object.Map {
	TCPLib := &object.Map{Map: map[string]object.Object{}}

	TCPLib.Map["server"] = NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.CLOSURE, "server", 0)
			vm.assertArgumentToType(args[1], object.NUMBER, "server", 1)
			closure := args[0].(object.Closure)
			port := args[1].(object.Number)
			// TODO: Determine when to close the server
			vm.RegisterEvent() // Register the server event

			handleConnection := func(conn net.Conn) {
				socket := NewSocketObject(conn, vm)
				vm.FireEvent(closure, socket)
			}

			startServer := func() {
				ln, err := net.Listen("tcp", fmt.Sprintf(":%d", int(port.Value)))
				if err != nil {
					vm.runtimeError("Error starting server: %v", err)
					return
				}
				defer ln.Close()

				for {
					conn, err := ln.Accept()
					if err != nil {
						fmt.Println("Accept error:", err)
						continue
					}
					go handleConnection(conn)
				}
			}

			go startServer() // Against the server event registered above
			return object.Nil{}
		},
		Arity: 2,
		Name:  "server",
	}

	TCPLib.Map["connect"] = NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.STRING, "connect", 0)
			vm.assertArgumentToType(args[1], object.CLOSURE, "connect", 1)
			address := args[0].(object.String)
			closure := args[1].(object.Closure)

			conn, err := net.Dial("tcp", address.Value)
			if err != nil {
				vm.FireEvent(closure,
					object.Nil{}, object.NewString(err.Error()))
				return object.Nil{}
			}

			socket := NewSocketObject(conn, vm)
			vm.FireEvent(closure, socket, object.Nil{})
			return object.Nil{}
		},
		Arity: 2,
		Name:  "connect",
	}

	return TCPLib

}
