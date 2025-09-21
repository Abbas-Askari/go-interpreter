package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getHttp() *object.Map {
	httpLib := &object.Map{Map: map[string]object.Object{}}

	httpLib.Map["listen"] = NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			// Return current time in seconds
			closure, ok := args[0].(object.Closure)
			port := args[1]
			if !ok {
				panic("First argument to listen must be a function")
			}
			p, ok := port.(object.Number)
			if !ok {
				panic("Second argument to listen must be a number")
			}
			vm.RegisterEvent()
			helloHandler := func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("Received request:", r.Method, r.URL.Path)
				done := make(chan struct{})

				headers := object.Map{Map: map[string]object.Object{}}
				for k, v := range r.Header {
					if len(v) > 0 {
						headers.Map[k] = object.NewString(v[0])
					}
				}
				body := NativeFunction{
					Function: func(vm *VM, args ...object.Object) object.Object {
						bodyBytes, err := io.ReadAll(r.Body)
						if err != nil {
							vm.runtimeError("Error reading body: %v ", err)
							return object.NewString("")
						}
						return object.NewString(string(bodyBytes))
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
				vm.FireEvent(closure, req, res)

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

	httpLib.Map["request"] = NativeFunction{
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.STRING, "request", 0)
			vm.assertArgumentToType(args[1], object.STRING, "request", 1)
			vm.assertArgumentToType(args[2], object.MAP, "request", 2)
			vm.assertArgumentToType(args[3], object.STRING, "request", 3)
			vm.assertArgumentToType(args[4], object.CLOSURE, "request", 4)
			verb := args[0].(object.String).Value
			url := args[1].(object.String).Value
			headers := args[2].(object.Map).Map
			body := args[3].(object.String).Value
			closure := args[4].(object.Closure)
			vm.RegisterEvent()
			handler := func() {
				defer vm.DetachEvent()
				client := &http.Client{}
				req, err := http.NewRequest(verb, url, nil)
				if err != nil {
					vm.FireEvent(closure, object.Nil{}, object.NewString(err.Error()))
					return
				}
				for k, v := range headers {
					req.Header.Set(k, v.String())
				}
				if body != "" {
					req.Body = io.NopCloser(strings.NewReader(body))
				}
				resp, err := client.Do(req)
				if err != nil {
					vm.FireEvent(closure, object.Nil{}, object.NewString(err.Error()))
					return
				}
				defer resp.Body.Close()
				respHeaders := object.Map{Map: map[string]object.Object{}}
				for k, v := range resp.Header {
					if len(v) > 0 {
						respHeaders.Map[k] = object.NewString(v[0])
					}
				}
				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					vm.FireEvent(closure, object.Nil{}, object.NewString(err.Error()))
					return
				}
				responseObject := object.Map{Map: map[string]object.Object{
					"status":  object.Number{Value: float64(resp.StatusCode)},
					"headers": respHeaders,
					"body":    object.NewString(string(bodyBytes)),
				}}
				vm.FireEvent(closure, responseObject, object.Nil{})
			}
			go handler()
			return object.Nil{}

		},
		Arity: 5,
		Name:  "request",
	}
	return httpLib

}
