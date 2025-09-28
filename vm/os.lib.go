package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"bytes"
	"io"
	"os"
	"os/exec"
)

func NewCommandObject(commandArray object.Array, vm *VM) object.Map {
	command := object.Map{Map: map[string]object.Object{}}
	command.Map["isRunning"] = object.Boolean{true}
	command.Map["onData"] = getMessagePrinterClosure("Warning: onData was called without an implementation", 1)
	command.Map["onError"] = getMessagePrinterClosure("Warning: onError was called without an implementation", 1)
	command.Map["onEnd"] = getMessagePrinterClosure("Warning: onEnd was called without an implementation", 0)

	cmdArray := []string{}
	for _, elem := range commandArray.Value {
		str, ok := elem.(object.String)
		if !ok {
			vm.runtimeError("All elements of the command array must be strings")
			return command
		}
		cmdArray = append(cmdArray, str.Value)
	}

	cmd := exec.Command(cmdArray[0], cmdArray[1:]...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	stdin, _ := cmd.StdinPipe()
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	done := make(chan struct{})

	command.Map["write"] = NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			buf := bytes.NewBuffer([]byte{})
			vm.assertArgumentToType(args[0], object.STRING, "write", 0)
			data := args[0].(object.String)
			_, err := buf.WriteString(data.Value)
			if err != nil {
				vm.FireEvent(command.Map["onError"].(object.Closure), object.NewString(err.Error()))
			}
			_, err = stdin.Write(buf.Bytes())
			if err != nil {
				vm.FireEvent(command.Map["onError"].(object.Closure), object.NewString(err.Error()))
			}
			return object.Nil{}
		},
		Arity: 1,
		Name:  "write",
	}

	readerRoutine := func(reader io.ReadCloser, functionName string) {
		buf := make([]byte, 1024)
		for {
			select {
			case <-done:
				return
			default:
				n, err := reader.Read(buf)
				if n > 0 {
					closure, ok := command.Map[functionName].(object.Closure)
					if !ok {
						vm.runtimeError(functionName + " is not a function")
					}
					vm.FireEvent(closure, object.NewString(string(buf[:n])))
				}
				if err != nil {
					break
				}
			}
		}
	}

	go readerRoutine(stdout, "onData")
	go readerRoutine(stderr, "onError")

	vm.RegisterEvent()
	go func() {
		cmd.Wait()
		vm.DetachEvent()
		stdin.Close()
		stdout.Close()
		stderr.Close()
		command.Map["isRunning"] = object.Boolean{false}
		vm.FireEvent(command.Map["onEnd"].(object.Closure))
		close(done)
	}()

	return command
}

func getOs() *object.Map {
	osLib := &object.Map{Map: map[string]object.Object{}}

	osLib.Map["exit"] = NativeFunction{
		Name:  "exit",
		Arity: 1,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.NUMBER, "exit", 0)
			code := args[0].(object.Number).Value
			os.Exit(int(code))
			return object.Nil{}
		},
	}

	osLib.Map["exec"] = NativeFunction{
		Name:  "exec",
		Arity: 1,
		Function: func(vm *VM, args ...object.Object) object.Object {
			if len(args) == 0 {
				vm.runtimeError("exec expects at least one argument")
			}
			vm.assertArgumentToType(args[0], object.ARRAY, "exec", 0)
			// commandStr := args[0].(object.Array).Elements[0].(object.String).Value
			cmd := args[0].(object.Array)
			imp := NewCommandObject(cmd, vm)
			return imp
		},
	}

	return osLib
}
