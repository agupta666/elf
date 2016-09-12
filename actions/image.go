package actions

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"
	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// ImageAction action represents actions which responds with an image
type ImageAction struct {
	Width  int
	Height int
	Color  color.Color
	Type   string
	Name   string
}

// HasName returns true if the image action has a name
func (ia *ImageAction) HasName() bool {
	return len(ia.Name) != 0
}

// NewImageActionFromExpr creates a new image action from an expression
func NewImageActionFromExpr(p string) (*ImageAction, error) {
	return parseImageExpr(p)
}

// DefaultImageAction creates a new image action from width and height
func DefaultImageAction(width string, height string) (*ImageAction, error) {

	nWidth, err := strconv.Atoi(width)
	if err != nil {
		return nil, errors.New("width must be integer")
	}

	nHeight, err := strconv.Atoi(height)
	if err != nil {
		return nil, errors.New("width must be integer")
	}

	return &ImageAction{
		Width:  nWidth,
		Height: nHeight,
		Color:  colorful.HappyColor(),
		Type:   ".png",
	}, nil
}

// NewImageAction creates a new image action from width, height and color
func NewImageAction(width string, height string, color string) (*ImageAction, error) {
	ia, err := DefaultImageAction(width, height)

	if err != nil {
		return nil, err
	}

	c, err := colorful.Hex(strings.TrimSpace(color))

	if err != nil {
		return nil, err
	}

	ia.Color = c

	return ia, nil
}

// NewImageActionWithType creates a new image action from width, height, color and type
func NewImageActionWithType(width string, height string, color string, ext string) (*ImageAction, error) {
	ia, err := NewImageAction(width, height, color)

	if err != nil {
		return nil, err
	}

	if ext == "" {
		return nil, errors.New("invalid image extension")
	}

	ext = strings.TrimSpace(ext)

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif":
		ia.Type = ext
		return ia, nil
	default:
		return nil, errors.New("unsupported image type")
	}

}

// NewImageActionWithName creates a new image action from width, height, color, type and name
func NewImageActionWithName(width string, height string, color string, ext string, name string) (*ImageAction, error) {
	ia, err := NewImageActionWithType(width, height, color, ext)

	if err != nil {
		return nil, err
	}

	ia.Name = strings.TrimSpace(name)
	return ia, nil

}

func parseImageExpr(s string) (*ImageAction, error) {
	s = strings.TrimPrefix(s, "image[")
	s = strings.TrimSuffix(s, "]")
	args := strings.Split(s, ",")

	switch len(args) {
	case 0, 1:
		return nil, errors.New("invalid expression")
	case 2:
		return DefaultImageAction(args[0], args[1])
	case 3:
		return NewImageAction(args[0], args[1], args[2])
	case 4:
		return NewImageActionWithType(args[0], args[1], args[2], args[3])
	case 5:
		return NewImageActionWithName(args[0], args[1], args[2], args[3], args[4])
	default:
		return nil, errors.New("invalid expression")

	}
}

// Exec executes a image action
func (ia *ImageAction) Exec(w http.ResponseWriter, r *http.Request) error {

	if ia.HasName() {
		writeFileName(ia.Name, ia.Type, w)
	}

	writeMimeType(ia.Type, w)

	m := image.NewRGBA(image.Rect(0, 0, ia.Width, ia.Height))

	draw.Draw(m, m.Bounds(), &image.Uniform{ia.Color}, image.ZP, draw.Src)

	switch ia.Type {
	case ".png":
		png.Encode(w, m)
	case ".gif":
		gif.Encode(w, m, nil)
	case ",jpg", ".jpeg":
		jpeg.Encode(w, m, nil)
	default:
		return errors.New("unsupported image format")
	}

	return nil
}
