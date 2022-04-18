package main

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// processMessageOrScalar gets the reflected value as its actual type. It is intended to be used for non-container values.
func processMessageOrScalar(fd protoreflect.FieldDescriptor, v protoreflect.Value) (object interface{}) {
	if fd.Kind() == protoreflect.MessageKind { // fd.Kind() will be the element type, even if the value is a List
		object = v.Message().Interface()
	} else {
		object = v.Interface()
	}
	return object

}

// processListField takes a list of Value / reflected objects and returns a list of interfaces / actual objects. The typing is not exactly preserved due to the difficulty / impossibility of doing so for every type without generators.
func processListField(fd protoreflect.FieldDescriptor, listObject protoreflect.List) []interface{} {
	l := make([]interface{}, listObject.Len())
	for i := range l {
		l[i] = processMessageOrScalar(fd, listObject.Get(i))
	}
	return l
}

// processMapField takes a map of Value / reflected objects and returns a map of interfaces / actual objects. The typing is not exactly preserved due to the difficulty / impossibility of doing so for every type without generators.
func processMapField(fd protoreflect.FieldDescriptor, mapObject protoreflect.Map) interface{} {

	mv := fd.MapValue()

	switch fd.MapKey().Kind() {
	case protoreflect.BoolKind:
		m := make(map[bool]interface{})
		mapObject.Range(func(key protoreflect.MapKey, v protoreflect.Value) bool {
			m[key.Bool()], _ = processField(mv, v)
			return true
		})
		return m
	case protoreflect.StringKind:
		m := make(map[string]interface{})
		mapObject.Range(func(key protoreflect.MapKey, v protoreflect.Value) bool {
			m[key.String()], _ = processField(mv, v)
			return true
		})
		return m
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		m := make(map[int64]interface{})
		mapObject.Range(func(key protoreflect.MapKey, v protoreflect.Value) bool {
			m[key.Int()], _ = processField(mv, v)
			return true
		})
		return m
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		m := make(map[uint64]interface{})
		mapObject.Range(func(key protoreflect.MapKey, v protoreflect.Value) bool {
			m[key.Uint()], _ = processField(mv, v)
			return true
		})
		return m
	}

	return nil
}

// processField holds the lines of logic for processing a reflected field.
func processField(fd protoreflect.FieldDescriptor, v protoreflect.Value) (object interface{}, valid bool) {

	valid = fd.HasJSONName()

	switch {
	case !valid: // skip if none
	case fd.IsMap():
		object = processMapField(fd, v.Map())
	case fd.IsList():
		object = processListField(fd, v.List())
	default:
		object = processMessageOrScalar(fd, v)
	}

	return object, valid
}

// processMessageAsMap takes a protobuf message and turns all the marshallable members at the first level into a map entry. This means that protobuf metadata members are not recorded in this map.
func processMessageAsMap(message proto.Message) map[string]interface{} {
	fields := make(map[string]interface{})

	if message != nil {
		message.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			processed, add := processField(fd, v)
			if add {
				fields[fd.JSONName()] = processed
			}
			return true
		})
	}

	return fields
}
