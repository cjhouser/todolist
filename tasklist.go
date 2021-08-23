package main

import "container/list"

type ops interface {
	create()
	read()
	update()
	delete()
}

type taskList struct {
	tasks list.List
}
