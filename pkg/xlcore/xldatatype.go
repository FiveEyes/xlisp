package core

import (
	"fmt"
	"strconv"
)

const (
	DT_Nil int = 0
	DT_Int int = 1
	DT_Float int = 2
	DT_String int = 3
	DT_Symbol int = 4
	DT_Function int = 5
	DT_NativeFunction int = 6
	DT_Pair int = 7
	DT_Map int = 8
	DT_Array int = 9
	DT_Lazy int = 10
)

type XLObj interface {
	XLObjType() int
}

type XLNil struct {
}

func (d *XLNil) XLObjType() int {
	return DT_Nil
}

var Nil XLObj = &XLNil{}
var Zero XLObj = NewXLInt(0)
var One XLObj = NewXLInt(1)

type XLInt struct {
	Value int
}

func (d *XLInt) XLObjType() int {
	return DT_Int
}

func NewXLInt(v int) *XLInt {
	return &XLInt{Value : v}
}

type XLFloat struct {
	Value float64
}

func (d *XLFloat) XLObjType() int {
	return DT_Float
}

func NewXLFloat(v float64) *XLFloat {
	return &XLFloat{Value : v}
}

type XLString struct {
	Value string
}

func (d *XLString) XLObjType() int {
	return DT_String
}

func NewXLString(v string) *XLString {
	return &XLString{Value : v}
}

type XLSymbol struct {
	Value string
}

func (d *XLSymbol) XLObjType() int {
	return DT_Symbol
}

func NewXLSymbol(v string) *XLSymbol {
	return &XLSymbol{Value : v}
}


type XLFunction struct {
	Params []string
	Body XLObj
	Env XLEnv
	Lazy bool
}

func (d XLFunction) XLObjType() int {
	return DT_Function
}

func NewXLFunction(params []string, body XLObj, env XLEnv, lazy bool) *XLFunction {
	return &XLFunction{Params:params, Body:body, Env:env, Lazy:lazy}
}

type XLNativeFunctionType func(XLObj, XLEnv) (XLObj, bool)

type XLNativeFunction struct {
//	Params []string
	Func XLNativeFunctionType
	Lazy bool
}

func (d XLNativeFunction) XLObjType() int {
	return DT_NativeFunction
}

func NewXLNativeFunction(f XLNativeFunctionType, lazy bool) *XLNativeFunction {
	return &XLNativeFunction{Func:f, Lazy:lazy}
}


type XLPair struct {
	Fst XLObj
	Snd XLObj
}

func (d *XLPair) XLObjType() int {
	return DT_Pair
}

func NewXLPair(fst XLObj, snd XLObj) *XLPair {
	return &XLPair{Fst:fst, Snd:snd}
}

type XLMap struct {
	Map map[string]XLObj
}

func (d *XLMap) XLObjType() int {
	return DT_Map
}

func NewXLMap() *XLMap {
	return &XLMap{Map:make(map[string]XLObj)}
}

type XLArr struct {
	Array []XLObj
}

func (d *XLArr) XLObjType() int {
	return DT_Array
}


type XLLazy struct {
	Value XLObj
	Env XLEnv
	Done bool

}

func (d *XLLazy) XLObjType() int {
	return DT_Lazy
}

func NewXLLazy(v XLObj, env XLEnv) *XLLazy {
	return &XLLazy{Value:v, Env:env, Done:false}
}

func PrettyPrint(v XLObj) string {
	//fmt.Println("PP", v)
	switch v.XLObjType() {
	case DT_Nil:
		return "(_)"
	case DT_Symbol:
		return v.(*XLSymbol).Value
	case DT_String:
		return "\"" + v.(*XLString).Value + "\""
	case DT_Int:
		return strconv.Itoa(v.(*XLInt).Value)
	case DT_Float:
		return fmt.Sprintf("%f", v.(*XLFloat).Value)
	case DT_Pair:
		fst := PrettyPrint(v.(*XLPair).Fst)
		rest := PrettyPrint(v.(*XLPair).Snd)
		return "(" + fst + " " + rest[1:]
	case DT_Lazy:
		return "(lazy: " + PrettyPrint(v.(*XLLazy).Value) + ")"
	default:
		return "???"
	}
}


func GetValue(obj XLObj) XLObj {
	switch obj.XLObjType() {
	case DT_Lazy:
		l := obj.(*XLLazy)
		//fmt.Println(&l, l)
		if !l.Done {
			l.Value, _ = ExpEval(l.Value, l.Env)
			l.Done = true
		}
		return l.Value
	default:
		return obj
	}
	return Nil
}

/*
const (
	EXP_Value int = 21
	EXP_App int = 22
)

type XLExp interface {
	XLExpType() int
}

type XLValue struct {
	Value XLObj
}

func (d XLValue) XLExpType() int {
	return EXP_Value
}

type XLApp struct {
	Exps []XLExp
}

func (d XLApp) XLExpType() int {
	return EXP_App
}
*/