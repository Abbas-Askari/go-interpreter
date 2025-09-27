package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"Abbas-Askari/interpreter-v2/op"
	"fmt"
	"io"
	"net"
)

func getMessagePrinterClosure(message string) object.Closure {
	return object.NewClosure(
		object.NewFunction(0, "printMessage", "internal", []op.OpCode{
			op.OpConstant, 0,
			op.OpPrint,
			op.OpNil,
			op.OpReturn,
		}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []object.Object{
			object.NewString(message),
		}))
}

func NewSocketObject(conn net.Conn) object.Map {
	socket := object.Map{Map: map[string]object.Object{}}
	socket.Map["isOpen"] = object.Boolean{true}
	socket.Map["onData"] = getMessagePrinterClosure("Warning: onData was called without an implementation")
	socket.Map["onError"] = getMessagePrinterClosure("Warning: onError was called without an implementation")
	socket.Map["onEnd"] = getMessagePrinterClosure("Warning: onEnd was called without an implementation")

	socket.Map["write"] = NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.STRING, "write", 0)
			data := args[0].(object.String)
			if socket.Map["isOpen"] == (object.Boolean{false}) {
				vm.runtimeError("Cannot write to a closed socket")
				return object.Nil{}
			}
			_, err := conn.Write([]byte(data.Value))
			if err != nil {
				vm.FireEvent(socket.Map["onError"].(object.Closure), object.NewString(err.Error()))
			}
			return object.Nil{}
		},
		Arity: 1,
		Name:  "write",
	}

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
			vm.RegisterEvent()

			handleConnection := func(conn net.Conn) {
				vm.RegisterEvent()
				socket := NewSocketObject(conn)

				buffer := make([]byte, 1024)

				end := func() {
					// I feel like this is cheating, because isOpen is visible in turtle
					// and I am changing it here
					// Some async code might be running and see that isOpen is false
					// but YOLO!
					socket.Map["isOpen"] = object.Boolean{false}
					vm.FireEvent(socket.Map["onEnd"].(object.Closure))
					vm.DetachEvent()
					conn.Close()
				}

				readerRoutine := func() {
					for {
						n, err := conn.Read(buffer)
						if err != nil {
							if err != io.EOF {
								vm.FireEvent(socket.Map["onError"].(object.Closure), object.NewString(err.Error()))
								continue
							}
							end()
							return
						}
						// Successfully read data
						data := make([]byte, n)
						copy(data, buffer[:n])
						vm.FireEvent(socket.Map["onData"].(object.Closure), object.NewString(string(data)))
					}
				}

				go readerRoutine()
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
					// Accept a new connection
					conn, err := ln.Accept()
					if err != nil {
						fmt.Println("Accept error:", err)
						continue
					}

					// Handle connection concurrently
					go handleConnection(conn)
				}
			}

			go startServer()
			return object.Nil{}

		},
		Arity: 2,
		Name:  "server",
	}
	return TCPLib

}
