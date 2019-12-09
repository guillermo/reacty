package document

import (
	"github.com/guillermo/reacty/commands"
	"image/color"
)

// Color holds the RGB information
type Color struct {
	R, G, B uint8
}

// RGBA implement image/color interface
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	g = uint32(c.G)
	b = uint32(c.B)

	padding := uint32(8)
	r = r<<padding + r
	g = g<<padding + g
	b = b<<padding + b
	a = 0xffff
	return
}

func FromColor(c color.Color) Color {
	r, g, b, _ := c.RGBA()
	padding := uint32(8)
	return Color{
		uint8(r >> padding),
		uint8(g >> padding),
		uint8(b >> padding),
	}
}

func NewColor(R, G, B uint8) Color {
	return Color{R, G, B}
}

// Background color
type Background Color

// sequence returns the sequence to output the specify color as background
func (b Background) sequence() []byte {
	return commands.Sequence("FGCOLOR", b.R, b.G, b.B)
}

// Foreground color
type Foreground Color

// sequence returns the sequence to output the specify color as foreground
func (f Foreground) sequence() []byte {
	return commands.Sequence("FGCOLOR", f.R, f.G, f.B)
}

var (
	// Maroon color
	Maroon = Color{128, 0, 0}
	// DarkRed color
	DarkRed = Color{139, 0, 0}
	// Brown color
	Brown = Color{165, 42, 42}
	// Firebrick color
	Firebrick = Color{178, 34, 34}
	// Crimson color
	Crimson = Color{220, 20, 60}
	// Red color
	Red = Color{255, 0, 0}
	// Tomato color
	Tomato = Color{255, 99, 71}
	// Coral color
	Coral = Color{255, 127, 80}
	// IndianRed color
	IndianRed = Color{205, 92, 92}
	// LightCoral color
	LightCoral = Color{240, 128, 128}
	// DarkSalmon color
	DarkSalmon = Color{233, 150, 122}
	// Salmon color
	Salmon = Color{250, 128, 114}
	// LightSalmon color
	LightSalmon = Color{255, 160, 122}
	// OrangeRed color
	OrangeRed = Color{255, 69, 0}
	// DarkOrange color
	DarkOrange = Color{255, 140, 0}
	// Orange color
	Orange = Color{255, 165, 0}
	// Gold color
	Gold = Color{255, 215, 0}
	// DarkGoldenRod color
	DarkGoldenRod = Color{184, 134, 11}
	// GoldenRod color
	GoldenRod = Color{218, 165, 32}
	// PaleGoldenRod color
	PaleGoldenRod = Color{238, 232, 170}
	// DarkKhaki color
	DarkKhaki = Color{189, 183, 107}
	// Khaki color
	Khaki = Color{240, 230, 140}
	// Olive color
	Olive = Color{128, 128, 0}
	// Yellow color
	Yellow = Color{255, 255, 0}
	// YellowGreen color
	YellowGreen = Color{154, 205, 50}
	// DarkOliveGreen color
	DarkOliveGreen = Color{85, 107, 47}
	// OliveDrab color
	OliveDrab = Color{107, 142, 35}
	// LawnGreen color
	LawnGreen = Color{124, 252, 0}
	// ChartReuse color
	ChartReuse = Color{127, 255, 0}
	// GreenYellow color
	GreenYellow = Color{173, 255, 47}
	// DarkGreen color
	DarkGreen = Color{0, 100, 0}
	// Green color
	Green = Color{0, 128, 0}
	// ForestGreen color
	ForestGreen = Color{34, 139, 34}
	// Lime color
	Lime = Color{0, 255, 0}
	// LimeGreen color
	LimeGreen = Color{50, 205, 50}
	// LightGreen color
	LightGreen = Color{144, 238, 144}
	// PaleGreen color
	PaleGreen = Color{152, 251, 152}
	// DarkSeaGreen color
	DarkSeaGreen = Color{143, 188, 143}
	// MediumSpringGreen color
	MediumSpringGreen = Color{0, 250, 154}
	// SpringGreen color
	SpringGreen = Color{0, 255, 127}
	// SeaGreen color
	SeaGreen = Color{46, 139, 87}
	// MediumAquaMarine color
	MediumAquaMarine = Color{102, 205, 170}
	// MediumSeaGreen color
	MediumSeaGreen = Color{60, 179, 113}
	// LightSeaGreen color
	LightSeaGreen = Color{32, 178, 170}
	// DarkSlateGray color
	DarkSlateGray = Color{47, 79, 79}
	// Teal color
	Teal = Color{0, 128, 128}
	// DarkCyan color
	DarkCyan = Color{0, 139, 139}
	// Aqua color
	Aqua = Color{0, 255, 255}
	// Cyan color
	Cyan = Color{0, 255, 255}
	// LightCyan color
	LightCyan = Color{224, 255, 255}
	// DarkTurquoise color
	DarkTurquoise = Color{0, 206, 209}
	// Turquoise color
	Turquoise = Color{64, 224, 208}
	// MediumTurquoise color
	MediumTurquoise = Color{72, 209, 204}
	// PaleTurquoise color
	PaleTurquoise = Color{175, 238, 238}
	// AquaMarine color
	AquaMarine = Color{127, 255, 212}
	// PowderBlue color
	PowderBlue = Color{176, 224, 230}
	// CadetBlue color
	CadetBlue = Color{95, 158, 160}
	// SteelBlue color
	SteelBlue = Color{70, 130, 180}
	// CornFlowerBlue color
	CornFlowerBlue = Color{100, 149, 237}
	// DeepSkyBlue color
	DeepSkyBlue = Color{0, 191, 255}
	// DodgerBlue color
	DodgerBlue = Color{30, 144, 255}
	// LightBlue color
	LightBlue = Color{173, 216, 230}
	// SkyBlue color
	SkyBlue = Color{135, 206, 235}
	// LightSkyBlue color
	LightSkyBlue = Color{135, 206, 250}
	// MidnightBlue color
	MidnightBlue = Color{25, 25, 112}
	// Navy color
	Navy = Color{0, 0, 128}
	// DarkBlue color
	DarkBlue = Color{0, 0, 139}
	// MediumBlue color
	MediumBlue = Color{0, 0, 205}
	// Blue color
	Blue = Color{0, 0, 255}
	// RoyalBlue color
	RoyalBlue = Color{65, 105, 225}
	// BlueViolet color
	BlueViolet = Color{138, 43, 226}
	// Indigo color
	Indigo = Color{75, 0, 130}
	// DarkSlateBlue color
	DarkSlateBlue = Color{72, 61, 139}
	// SlateBlue color
	SlateBlue = Color{106, 90, 205}
	// MediumSlateBlue color
	MediumSlateBlue = Color{123, 104, 238}
	// MediumPurple color
	MediumPurple = Color{147, 112, 219}
	// DarkMagenta color
	DarkMagenta = Color{139, 0, 139}
	// DarkViolet color
	DarkViolet = Color{148, 0, 211}
	// DarkOrchid color
	DarkOrchid = Color{153, 50, 204}
	// MediumOrchid color
	MediumOrchid = Color{186, 85, 211}
	// Purple color
	Purple = Color{128, 0, 128}
	// Thistle color
	Thistle = Color{216, 191, 216}
	// Plum color
	Plum = Color{221, 160, 221}
	// Violet color
	Violet = Color{238, 130, 238}
	// Magenta color
	Magenta = Color{255, 0, 255}
	// Fuchsia color
	Fuchsia = Color{255, 0, 255}
	// Orchid color
	Orchid = Color{218, 112, 214}
	// MediumVioletRed color
	MediumVioletRed = Color{199, 21, 133}
	// PaleVioletRed color
	PaleVioletRed = Color{219, 112, 147}
	// DeepPink color
	DeepPink = Color{255, 20, 147}
	// HotPink color
	HotPink = Color{255, 105, 180}
	// LightPink color
	LightPink = Color{255, 182, 193}
	// Pink color
	Pink = Color{255, 192, 203}
	// AntiqueWhite color
	AntiqueWhite = Color{250, 235, 215}
	// Beige color
	Beige = Color{245, 245, 220}
	// Bisque color
	Bisque = Color{255, 228, 196}
	// BlanchedAlmond color
	BlanchedAlmond = Color{255, 235, 205}
	// Wheat color
	Wheat = Color{245, 222, 179}
	// CornSilk color
	CornSilk = Color{255, 248, 220}
	// LemonChiffon color
	LemonChiffon = Color{255, 250, 205}
	// LightGoldenRodYellow color
	LightGoldenRodYellow = Color{250, 250, 210}
	// LightYellow color
	LightYellow = Color{255, 255, 224}
	// SaddleBrown color
	SaddleBrown = Color{139, 69, 19}
	// Sienna color
	Sienna = Color{160, 82, 45}
	// Chocolate color
	Chocolate = Color{210, 105, 30}
	// Peru color
	Peru = Color{205, 133, 63}
	// SandyBrown color
	SandyBrown = Color{244, 164, 96}
	// BurlyWood color
	BurlyWood = Color{222, 184, 135}
	// Tan color
	Tan = Color{210, 180, 140}
	// RosyBrown color
	RosyBrown = Color{188, 143, 143}
	// Moccasin color
	Moccasin = Color{255, 228, 181}
	// NavajoWhite color
	NavajoWhite = Color{255, 222, 173}
	// PeachPuff color
	PeachPuff = Color{255, 218, 185}
	// MistyRose color
	MistyRose = Color{255, 228, 225}
	// LavenderBlush color
	LavenderBlush = Color{255, 240, 245}
	// Linen color
	Linen = Color{250, 240, 230}
	// OldLace color
	OldLace = Color{253, 245, 230}
	// PapayaWhip color
	PapayaWhip = Color{255, 239, 213}
	// SeaShell color
	SeaShell = Color{255, 245, 238}
	// MintCream color
	MintCream = Color{245, 255, 250}
	// SlateGray color
	SlateGray = Color{112, 128, 144}
	// LightSlateGray color
	LightSlateGray = Color{119, 136, 153}
	// LightSteelBlue color
	LightSteelBlue = Color{176, 196, 222}
	// Lavender color
	Lavender = Color{230, 230, 250}
	// FloralWhite color
	FloralWhite = Color{255, 250, 240}
	// AliceBlue color
	AliceBlue = Color{240, 248, 255}
	// GhostWhite color
	GhostWhite = Color{248, 248, 255}
	// Honeydew color
	Honeydew = Color{240, 255, 240}
	// Ivory color
	Ivory = Color{255, 255, 240}
	// Azure color
	Azure = Color{240, 255, 255}
	// Snow color
	Snow = Color{255, 250, 250}
	// Black color
	Black = Color{0, 0, 0}
	// DimGray color
	DimGray = Color{105, 105, 105}
	// DimGrey color
	DimGrey = Color{105, 105, 105}
	// Gray color
	Gray = Color{128, 128, 128}
	// Grey color
	Grey = Color{128, 128, 128}
	// DarkGray color
	DarkGray = Color{169, 169, 169}
	// DarkGrey color
	DarkGrey = Color{169, 169, 169}
	// Silver color
	Silver = Color{192, 192, 192}
	// LightGray color
	LightGray = Color{211, 211, 211}
	// LightGrey color
	LightGrey = Color{211, 211, 211}
	// Gainsboro color
	Gainsboro = Color{220, 220, 220}
	// WhiteSmoke color
	WhiteSmoke = Color{245, 245, 245}
	// White color
	White = Color{255, 255, 255}
)
