package main

type ops interface {
	create()
	read()
	update()
	delete()
}
