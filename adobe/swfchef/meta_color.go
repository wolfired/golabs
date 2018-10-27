package swfchef

type rgb struct {
	red, green, blue ui8
}

type rgba struct {
	rgb
	alpha ui8
}

type argb rgba
