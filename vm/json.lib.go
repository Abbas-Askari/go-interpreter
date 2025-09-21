package vm

import (
	"Abbas-Askari/interpreter-v2/object"
	"encoding/json"
	"fmt"
)

func convertToObject(v interface{}) object.Object {
	switch val := v.(type) {
	case nil:
		return object.Nil{}
	case bool:
		return object.Boolean{Value: val}
	case float64:
		return object.Number{Value: val}
	case string:
		return object.String{Value: val}
	case []interface{}:
		arr := object.NewArray([]object.Object{})
		for _, elem := range val {
			arr.Value = append(arr.Value, convertToObject(elem))
		}
		return arr
	case map[string]interface{}:
		obj := object.Map{Map: map[string]object.Object{}}
		for k, vv := range val {
			obj.Map[k] = convertToObject(vv)
		}
		return obj
	default:
		return nil
	}
}

func stringify(v object.Object) object.Object {
	switch val := v.(type) {
	case nil:
		return object.Nil{}
	case object.Nil:
		return object.NewString("null")
	case object.Boolean:
		return object.NewString(fmt.Sprintf("%v", val.Value))
	case object.Number:
		return object.NewString(fmt.Sprintf("%v", val.Value))
	case object.String:
		str := val.Value
		// Escape special characters
		str = fmt.Sprintf("%q", str)
		return object.NewString(str)
	case object.Array:
		str := "["
		for i, elem := range val.Value {
			str += fmt.Sprint(stringify(elem))
			if i < len(val.Value)-1 {
				str += ", "
			}
		}
		str += "]"
		return object.NewString(str)
	case object.Map:
		str := "{"
		for k, v := range val.Map {
			str += fmt.Sprintf("\"%s\": %s, ", k, stringify(v))
		}
		if len(val.Map) > 0 {
			str = str[:len(str)-2] // Remove trailing comma and space
		}
		str += "}"
		return object.NewString(str)
	default:
		return nil
	}
}

func parse(str string) object.Object {
	var v interface{}
	err := json.Unmarshal([]byte(str), &v)
	if err != nil {
		fmt.Println("JSON error:", err)
		return nil
	}
	return convertToObject(v)
}

func getJson() *object.Map {
	jsonLib := &object.Map{Map: map[string]object.Object{}}

	jsonLib.Map["parse"] = NativeFunction{
		Name:  "parse",
		Arity: 1,
		Function: func(vm *VM, args ...object.Object) object.Object {
			vm.assertArgumentToType(args[0], object.STRING, "parse", 0)
			str := args[0].(object.String).Value
			obj := parse(str)
			if obj == nil {
				return object.Nil{}
			}
			return obj

		},
	}

	jsonLib.Map["stringify"] = NativeFunction{
		Name:  "stringify",
		Arity: 1,
		Function: func(vm *VM, args ...object.Object) object.Object {
			obj := args[0]
			res := stringify(obj)
			if res == nil {
				vm.runtimeError("Error stringifying JSON")
				return object.Nil{}
			}
			return res
		},
	}

	return jsonLib
}
