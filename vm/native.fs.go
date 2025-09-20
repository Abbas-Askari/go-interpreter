package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"fmt"
	"log"
	"net/http"
)

func getFileSystem() NativeFunction {
	return NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			// Return current time in seconds
			closure := args[0]
			port := args[1]
			if _, ok := closure.(object.Closure); !ok {
				panic("First argument to listen must be a function")
			}
			p, ok := port.(object.Number)
			if !ok {
				panic("Second argument to listen must be a number")
			}
			eventIndex := vm.RegisterEvent(closure.(object.Closure))
			helloHandler := func(w http.ResponseWriter, r *http.Request) {
				done := make(chan struct{})

				headers := object.Map{Map: map[string]object.Object{}}
				for k, v := range r.Header {
					if len(v) > 0 {
						headers.Map[k] = object.NewString(v[0])
					}
				}
				body := NativeFunction{
					Function: func(vm *VM, args ...object.Object) object.Object {
						buf := make([]byte, r.ContentLength)
						_, err := r.Body.Read(buf)
						if err != nil {
							log.Println("Error reading body:", err)
							return object.NewString("")
						}
						return object.NewString(string(buf))
					}, Arity: 0, Name: "readBody",
				}
				req := object.Map{Map: map[string]object.Object{
					"path":    object.NewString(r.URL.Path),
					"headers": headers,
					"method":  object.NewString(r.Method),
					"body":    body,
				}}
				// fire event with req and res objects

				writeHeader := NativeFunction{
					Function: func(vm *VM, args ...object.Object) object.Object {
						w.Header().Set(args[0].String(), args[1].String())
						fmt.Println("Set header:", args[0].String(), args[1].String())
						return object.Nil{}
					}, Arity: 2, Name: "writeHeader",
				}
				flusher := w.(http.Flusher)
				flush := NativeFunction{
					Function: func(vm *VM, args ...object.Object) object.Object {
						flusher.Flush()
						return object.Nil{}
					}, Arity: 0, Name: "flush",
				}
				writeStatus := NativeFunction{
					Function: func(vm *VM, args ...object.Object) object.Object {
						statusCode, ok := args[0].(object.Number)
						if !ok {
							panic("First argument to writeStatus must be a number")
						}
						w.WriteHeader(int(statusCode.Value))
						return object.Nil{}
					}, Arity: 1, Name: "writeStatus",
				}
				writeBody := NativeFunction{
					Function: func(vm *VM, args ...object.Object) object.Object {
						body, ok := args[0].(object.String)
						if !ok {
							panic("First argument to write must be a string")
						}
						_, err := w.Write([]byte(body.Value))
						if err != nil {
							log.Println("Error writing body:", err)
						}
						return object.Nil{}
					}, Arity: 1, Name: "write",
				}
				end := NativeFunction{
					Function: func(vm *VM, args ...object.Object) object.Object {
						close(done)
						return object.Nil{}
					}, Arity: 0, Name: "end",
				}
				// create a
				// response object with writeHeader and write functions

				res := object.Map{Map: map[string]object.Object{
					"writeStatus": writeStatus,
					"writeHeader": writeHeader,
					"write":       writeBody,
					"end":         end,
					"flush":       flush,
				}}
				vm.FireEvent(eventIndex, req, res)

				<-done
			}
			http.HandleFunc("/", helloHandler)

			go func() {
				if err := http.ListenAndServe(fmt.Sprintf(":%d", int(p.Value)), nil); err != nil {
					panic(err)
				}
			}()
			return object.Nil{}
		},
		Arity: 2,
		Name:  "listen",
	}
}
