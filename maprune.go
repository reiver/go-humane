package humane

// mapRune is used to change the numerical order of runes to a more human-friendly order.
//
// In particular, instead of having this order:
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
// It changes it to this (more human-friendly) order.
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
//	[
//	\
//	]
//	^
//	_
//	`
//
// All other runes stay where they are (in the ordering).
func mapRune(r rune) rune {

	switch {
	case r < ':':
		return r
	case (':' <= r) && (r <= '@'):
		return r - ':' + ('z' + 1) - ('@' - ':' + 1) - ('`' - '[' + 1)
	case ('A' <= r) && (r <= 'Z'):
		return ((r - 'A') * 2) + ':'
	case ('[' <= r) && (r <= '`'):
		return r - '[' + ('z' + 1) - ('@' - ':')
	case ('a' <= r) && (r <= 'z'):
		return ((r - 'a') * 2) + (':' + 1)
	case 'z' < r:
		return r;
	default:
		return r
	}
}
