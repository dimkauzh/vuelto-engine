/*
 * Copyright (C) 2025 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package trita

import (
	"reflect"
)

const Func = reflect.Func

type TritaType struct {
	Kind         any
	NumArguments int
	ReflectValue reflect.Value
}

func NewTritaValue(Variable any) *TritaType {
	varoutput := reflect.ValueOf(Variable)
	return &TritaType{
		Kind:         varoutput.Kind(),
		NumArguments: varoutput.Type().NumIn(),
		ReflectValue: varoutput,
	}
}

func (trita *TritaType) TypeArgument(ArgumentNum int) any {
	return trita.ReflectValue.Type().In(ArgumentNum)
}

func YourType(Types any) any {
	return reflect.TypeOf(Types)
}

func (trita *TritaType) Call(args ...any) []reflect.Value {
	argumentsList := make([]reflect.Value, len(args))
	for i, arg := range args {
		argumentsList[i] = reflect.ValueOf(arg)
	}
	return trita.ReflectValue.Call(argumentsList)
}

func (trita *TritaType) HasReturn() bool {
	return trita.ReflectValue.Type().NumOut() > 0
}
