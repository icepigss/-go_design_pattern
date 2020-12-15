package main

import ()

type Undo []func()

func (undo *Undo) Add(f func()) {
	*undo = append(*undo, f)
}

func (undo *Undo) Undo() {
	if len(*undo) == 0 {
		return
	}

	for _, f := range *undo {
		f()
	}
}

type Deal struct {
	data map[int]bool
	undo Undo
}

func (d *Deal) Add(i int) {
	if d.Contain(i) {
		return
	}

	d.data[i] = true
	d.undo.Add(func() {
		d.Delete(i)
	})
}

func (d *Deal) Delete(i int) {
	if !d.Contain(i) {
		return
	}

	delete(d.data, i)
	d.undo.Add(func() {
		d.Add(i)
	})
}

func (d *Deal) Contain(i int) bool {
	return d.data[i]
}

func (d *Deal) Undo() {
	d.undo.Undo()
}

func main() {
	d := &Deal{
		data: make(map[int]bool),
	}
	d.Add(1)
	d.Add(2)
	d.Add(3)
	d.Delete(3)

	d.Undo()

	pr(d.data)
}

func pr(m map[int]bool) {
	for k, _ := range m {
		println(k)
	}
}
