package main

import (
	"fmt"
)

type VisitorFunc func(*Info) error

type Info struct {
	User string
	Num  int
}

type Visitor interface {
	Visit(VisitorFunc) error
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info)
}

type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info) error {
		println("log before")
		err := fn(info)
		if err == nil {
			fmt.Printf("info.User:%s,info.Num:%d\n", info.User, info.Num)
		}
		println("log after")

		return err
	})
}

type FormatVisitor struct {
	visitor Visitor
}

func (v FormatVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info) error {
		println("format before")
		err := fn(info)
		if err == nil {
			info.User += "_suffix"
			info.Num += 1
		}
		println("format after")

		return err
	})
}

func main() {
	var v Visitor = &Info{}
	v = LogVisitor{v}
	v = FormatVisitor{v}

	load := func(info *Info) error {
		info.User = "icepigss"
		info.Num = 1
		return nil
	}
	err := v.Visit(load)
	if err != nil {
		fmt.Printf("error:%+v\n", err)
	}
}
