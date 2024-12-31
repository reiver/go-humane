package humane

import (
	"slices"
)

// SortStrings sorts strings in a more human-friendly ordering.
//
// It mostly uses UNICODE ordering, with an exception.
//
// The exception is that, the ordering of these:
//
//	0
//	1
//	2
//	3
//	4
//	5
//	6
//	7
//	8
//	9
//	:
//	;
//	<
//	=
//	>
//	?
//	@
//	A
//	B
//	C
//	D
//	E
//	F
//	G
//	H
//	I
//	J
//	K
//	L
//	M
//	N
//	O
//	P
//	Q
//	R
//	S
//	T
//	U
//	V
//	W
//	X
//	Y
//	Z
//	[
//	\
//	]
//	^
//	_
//	`
//	a
//	b
//	c
//	d
//	e
//	f
//	g
//	h
//	i
//	j
//	k
//	l
//	m
//	n
//	o
//	p
//	q
//	r
//	s
//	t
//	u
//	v
//	w
//	x
//	y
//	z
//
// Are changed to:
//
//	0
//	1
//	2
//	3
//	4
//	5
//	6
//	7
//	8
//	9
//	A
//	a
//	B
//	b
//	C
//	c
//	D
//	d
//	E
//	e
//	F
//	f
//	G
//	g
//	H
//	h
//	I
//	i
//	J
//	j
//	K
//	k
//	L
//	l
//	M
//	m
//	N
//	n
//	O
//	o
//	P
//	p
//	Q
//	q
//	R
//	r
//	S
//	s
//	T
//	t
//	U
//	u
//	V
//	v
//	W
//	w
//	X
//	x
//	Y
//	y
//	Z
//	z
//	:
//	;
//	<
//	=
//	>
//	?
//	@
//	[
//	\
//	]
//	^
//	_
//	`
func SortStrings(strings []string) {
	slices.SortFunc(strings, cmpStrings)
}
