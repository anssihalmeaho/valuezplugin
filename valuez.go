package main

import (
	"github.com/anssihalmeaho/fuvaluez/fuvaluez"

	"github.com/anssihalmeaho/funl/funl"
)

type proceInfo struct {
	Name   string
	Getter func(string) fuvaluez.FZProc
}

func Setup(napi funl.FNIApi) (err error) {
	vzProcs := []proceInfo{
		{
			Name:   "open",
			Getter: fuvaluez.GetVZOpen,
		},
		{
			Name:   "new-col",
			Getter: fuvaluez.GetVZNewCol,
		},
		{
			Name:   "get-col",
			Getter: fuvaluez.GetVZGetCol,
		},
		{
			Name:   "get-col-names",
			Getter: fuvaluez.GetVZGetColNames,
		},
		{
			Name:   "put-value",
			Getter: fuvaluez.GetVZPutValue,
		},
		{
			Name:   "get-values",
			Getter: fuvaluez.GetVZGetValues,
		},
		{
			Name:   "take-values",
			Getter: fuvaluez.GetVZTakeValues,
		},
		{
			Name:   "update",
			Getter: fuvaluez.GetVZUpdate,
		},
		{
			Name:   "trans",
			Getter: fuvaluez.GetVZTrans,
		},
		{
			Name:   "view",
			Getter: fuvaluez.GetVZView,
		},
		{
			Name:   "del-col",
			Getter: fuvaluez.GetVZDelCol,
		},
		{
			Name:   "close",
			Getter: fuvaluez.GetVZClose,
		},
	}

	for _, vzProc := range vzProcs {
		err = napi.RegExtProc(funl.ExtProcType{Impl: vzProc.Getter(vzProc.Name)}, vzProc.Name)
		if err != nil {
			return
		}
	}
	return
}
