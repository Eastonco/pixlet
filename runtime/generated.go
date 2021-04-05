package runtime

// Code generated by runtime/gen. DO NOT EDIT.

import (
	"fmt"
	"sync"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/mitchellh/hashstructure"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"tidbyt.dev/pixlet/render"
)

const (
	ModuleName = "render"
)

var (
	once   sync.Once
	module starlark.StringDict
)

func LoadModule() (starlark.StringDict, error) {
	once.Do(func() {
		fnt := starlark.NewDict(len(render.Font))
		for k, _ := range render.Font {
			fnt.SetKey(starlark.String(k), starlark.String(k))
		}
		fnt.Freeze()

		module = starlark.StringDict{
			ModuleName: &starlarkstruct.Module{
				Name: ModuleName,
				Members: starlark.StringDict{
					"Root": starlark.NewBuiltin("Root", newRoot),
					"Plot": starlark.NewBuiltin("Plot", newPlot),
					"AnimatedPositioned": starlark.NewBuiltin(
						"AnimatedPositioned",
						newAnimatedPositioned,
					),
					"fonts": fnt,

					"Animation": starlark.NewBuiltin("Animation", newAnimation),

					"Box": starlark.NewBuiltin("Box", newBox),

					"Circle": starlark.NewBuiltin("Circle", newCircle),

					"Column": starlark.NewBuiltin("Column", newColumn),

					"Image": starlark.NewBuiltin("Image", newImage),

					"Marquee": starlark.NewBuiltin("Marquee", newMarquee),

					"Padding": starlark.NewBuiltin("Padding", newPadding),

					"Row": starlark.NewBuiltin("Row", newRow),

					"Stack": starlark.NewBuiltin("Stack", newStack),

					"Text": starlark.NewBuiltin("Text", newText),

					"WrappedText": starlark.NewBuiltin("WrappedText", newWrappedText),
				},
			},
		}
	})

	return module, nil
}

type Widget interface {
	AsRenderWidget() render.Widget
}
type Animation struct {
	Widget
	render.Animation

	starlarkChildren *starlark.List
}

func newAnimation(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		children *starlark.List
	)

	if err := starlark.UnpackArgs(
		"Animation",
		args, kwargs,
		"children?", &children,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Animation: %s", err)
	}

	w := &Animation{}

	var childrenVal starlark.Value
	childrenIter := children.Iterate()
	defer childrenIter.Done()
	for i := 0; childrenIter.Next(&childrenVal); {
		if _, isNone := childrenVal.(starlark.NoneType); isNone {
			continue
		}

		childrenChild, ok := childrenVal.(Widget)
		if !ok {
			return nil, fmt.Errorf(
				"expected children to be a list of Widget but found: %s (at index %d)",
				childrenVal.Type(),
				i,
			)
		}

		w.Children = append(w.Children, childrenChild.AsRenderWidget())
	}
	w.starlarkChildren = children

	return w, nil
}

func (w *Animation) AsRenderWidget() render.Widget {
	return &w.Animation
}

func (w *Animation) AttrNames() []string {
	return []string{
		"children",
	}
}

func (w *Animation) Attr(name string) (starlark.Value, error) {
	switch name {

	case "children":
		return w.starlarkChildren, nil

	default:
		return nil, nil
	}
}

func (w *Animation) String() string       { return "Animation(...)" }
func (w *Animation) Type() string         { return "Animation" }
func (w *Animation) Freeze()              {}
func (w *Animation) Truth() starlark.Bool { return true }

func (w *Animation) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

type Box struct {
	Widget
	render.Box

	starlarkChild starlark.Value
}

func newBox(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		width   starlark.Int
		height  starlark.Int
		padding starlark.Int

		color starlark.String

		child starlark.Value
	)

	if err := starlark.UnpackArgs(
		"Box",
		args, kwargs,
		"child?", &child,
		"width?", &width,
		"height?", &height,
		"padding?", &padding,
		"color?", &color,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Box: %s", err)
	}

	w := &Box{}
	w.Width = int(width.BigInt().Int64())
	w.Height = int(height.BigInt().Int64())
	w.Padding = int(padding.BigInt().Int64())

	if color.Len() > 0 {
		c, err := colorful.Hex(color.GoString())
		if err != nil {
			return nil, fmt.Errorf("color is not a valid hex string: %s", color.String())
		}
		w.Color = c
	}

	if child != nil {
		childWidget, ok := child.(Widget)
		if !ok {
			return nil, fmt.Errorf(
				"invalid type for child: %s (expected Widget)",
				child.Type(),
			)
		}
		w.Child = childWidget.AsRenderWidget()
		w.starlarkChild = child
	}

	return w, nil
}

func (w *Box) AsRenderWidget() render.Widget {
	return &w.Box
}

func (w *Box) AttrNames() []string {
	return []string{
		"child", "width", "height", "padding", "color",
	}
}

func (w *Box) Attr(name string) (starlark.Value, error) {
	switch name {

	case "width":
		return starlark.MakeInt(w.Width), nil

	case "height":
		return starlark.MakeInt(w.Height), nil

	case "padding":
		return starlark.MakeInt(w.Padding), nil

	case "color":
		if w.Color == nil {
			return nil, nil
		}
		c, ok := colorful.MakeColor(w.Color)
		if !ok {
			return nil, nil
		}
		return starlark.String(c.Hex()), nil

	case "child":
		return w.starlarkChild, nil

	default:
		return nil, nil
	}
}

func (w *Box) String() string       { return "Box(...)" }
func (w *Box) Type() string         { return "Box" }
func (w *Box) Freeze()              {}
func (w *Box) Truth() starlark.Bool { return true }

func (w *Box) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

type Circle struct {
	Widget
	render.Circle

	starlarkChild starlark.Value
}

func newCircle(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		diameter starlark.Int

		color starlark.String

		child starlark.Value
	)

	if err := starlark.UnpackArgs(
		"Circle",
		args, kwargs,
		"color", &color,
		"diameter", &diameter,
		"child?", &child,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Circle: %s", err)
	}

	w := &Circle{}
	w.Diameter = int(diameter.BigInt().Int64())

	if color.Len() > 0 {
		c, err := colorful.Hex(color.GoString())
		if err != nil {
			return nil, fmt.Errorf("color is not a valid hex string: %s", color.String())
		}
		w.Color = c
	}

	if child != nil {
		childWidget, ok := child.(Widget)
		if !ok {
			return nil, fmt.Errorf(
				"invalid type for child: %s (expected Widget)",
				child.Type(),
			)
		}
		w.Child = childWidget.AsRenderWidget()
		w.starlarkChild = child
	}

	return w, nil
}

func (w *Circle) AsRenderWidget() render.Widget {
	return &w.Circle
}

func (w *Circle) AttrNames() []string {
	return []string{
		"color", "diameter", "child",
	}
}

func (w *Circle) Attr(name string) (starlark.Value, error) {
	switch name {

	case "diameter":
		return starlark.MakeInt(w.Diameter), nil

	case "color":
		if w.Color == nil {
			return nil, nil
		}
		c, ok := colorful.MakeColor(w.Color)
		if !ok {
			return nil, nil
		}
		return starlark.String(c.Hex()), nil

	case "child":
		return w.starlarkChild, nil

	default:
		return nil, nil
	}
}

func (w *Circle) String() string       { return "Circle(...)" }
func (w *Circle) Type() string         { return "Circle" }
func (w *Circle) Freeze()              {}
func (w *Circle) Truth() starlark.Bool { return true }

func (w *Circle) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

type Column struct {
	Widget
	render.Column

	starlarkChildren *starlark.List
}

func newColumn(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		main_align  starlark.String
		cross_align starlark.String

		expanded starlark.Bool

		children *starlark.List
	)

	if err := starlark.UnpackArgs(
		"Column",
		args, kwargs,
		"children", &children,
		"main_align?", &main_align,
		"cross_align?", &cross_align,
		"expanded?", &expanded,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Column: %s", err)
	}

	w := &Column{}
	w.MainAlign = main_align.GoString()
	w.CrossAlign = cross_align.GoString()
	w.Expanded = bool(expanded)

	var childrenVal starlark.Value
	childrenIter := children.Iterate()
	defer childrenIter.Done()
	for i := 0; childrenIter.Next(&childrenVal); {
		if _, isNone := childrenVal.(starlark.NoneType); isNone {
			continue
		}

		childrenChild, ok := childrenVal.(Widget)
		if !ok {
			return nil, fmt.Errorf(
				"expected children to be a list of Widget but found: %s (at index %d)",
				childrenVal.Type(),
				i,
			)
		}

		w.Children = append(w.Children, childrenChild.AsRenderWidget())
	}
	w.starlarkChildren = children

	return w, nil
}

func (w *Column) AsRenderWidget() render.Widget {
	return &w.Column
}

func (w *Column) AttrNames() []string {
	return []string{
		"children", "main_align", "cross_align", "expanded",
	}
}

func (w *Column) Attr(name string) (starlark.Value, error) {
	switch name {

	case "main_align":
		return starlark.String(w.MainAlign), nil

	case "cross_align":
		return starlark.String(w.CrossAlign), nil

	case "expanded":
		return starlark.Bool(w.Expanded), nil

	case "children":
		return w.starlarkChildren, nil

	default:
		return nil, nil
	}
}

func (w *Column) String() string       { return "Column(...)" }
func (w *Column) Type() string         { return "Column" }
func (w *Column) Freeze()              {}
func (w *Column) Truth() starlark.Bool { return true }

func (w *Column) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

type Image struct {
	Widget
	render.Image

	size *starlark.Builtin
}

func newImage(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		src starlark.String

		width  starlark.Int
		height starlark.Int
	)

	if err := starlark.UnpackArgs(
		"Image",
		args, kwargs,
		"src", &src,
		"width?", &width,
		"height?", &height,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Image: %s", err)
	}

	w := &Image{}
	w.Src = src.GoString()
	w.Width = int(width.BigInt().Int64())
	w.Height = int(height.BigInt().Int64())

	w.size = starlark.NewBuiltin("size", imageSize)

	return w, nil
}

func (w *Image) AsRenderWidget() render.Widget {
	return &w.Image
}

func (w *Image) AttrNames() []string {
	return []string{
		"src", "width", "height",
	}
}

func (w *Image) Attr(name string) (starlark.Value, error) {
	switch name {

	case "src":
		return starlark.String(w.Src), nil

	case "width":
		return starlark.MakeInt(w.Width), nil

	case "height":
		return starlark.MakeInt(w.Height), nil

	case "size":
		return w.size.BindReceiver(w), nil

	default:
		return nil, nil
	}
}

func (w *Image) String() string       { return "Image(...)" }
func (w *Image) Type() string         { return "Image" }
func (w *Image) Freeze()              {}
func (w *Image) Truth() starlark.Bool { return true }

func (w *Image) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

func imageSize(
	thread *starlark.Thread,
	b *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	w := b.Receiver().(*Image)
	width, height := w.Size()

	return starlark.Tuple([]starlark.Value{
		starlark.MakeInt(width),
		starlark.MakeInt(height),
	}), nil
}

type Marquee struct {
	Widget
	render.Marquee

	starlarkChild starlark.Value
}

func newMarquee(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		width        starlark.Int
		offset_start starlark.Int
		offset_end   starlark.Int

		child starlark.Value
	)

	if err := starlark.UnpackArgs(
		"Marquee",
		args, kwargs,
		"child", &child,
		"width", &width,
		"offset_start?", &offset_start,
		"offset_end?", &offset_end,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Marquee: %s", err)
	}

	w := &Marquee{}
	w.Width = int(width.BigInt().Int64())
	w.OffsetStart = int(offset_start.BigInt().Int64())
	w.OffsetEnd = int(offset_end.BigInt().Int64())

	if child != nil {
		childWidget, ok := child.(Widget)
		if !ok {
			return nil, fmt.Errorf(
				"invalid type for child: %s (expected Widget)",
				child.Type(),
			)
		}
		w.Child = childWidget.AsRenderWidget()
		w.starlarkChild = child
	}

	return w, nil
}

func (w *Marquee) AsRenderWidget() render.Widget {
	return &w.Marquee
}

func (w *Marquee) AttrNames() []string {
	return []string{
		"child", "width", "offset_start", "offset_end",
	}
}

func (w *Marquee) Attr(name string) (starlark.Value, error) {
	switch name {

	case "width":
		return starlark.MakeInt(w.Width), nil

	case "offset_start":
		return starlark.MakeInt(w.OffsetStart), nil

	case "offset_end":
		return starlark.MakeInt(w.OffsetEnd), nil

	case "child":
		return w.starlarkChild, nil

	default:
		return nil, nil
	}
}

func (w *Marquee) String() string       { return "Marquee(...)" }
func (w *Marquee) Type() string         { return "Marquee" }
func (w *Marquee) Freeze()              {}
func (w *Marquee) Truth() starlark.Bool { return true }

func (w *Marquee) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

type Padding struct {
	Widget
	render.Padding

	starlarkChild starlark.Value

	starlarkPad starlark.Value
}

func newPadding(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		expanded starlark.Bool

		child starlark.Value

		pad starlark.Value
	)

	if err := starlark.UnpackArgs(
		"Padding",
		args, kwargs,
		"child", &child,
		"pad?", &pad,
		"expanded?", &expanded,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Padding: %s", err)
	}

	w := &Padding{}
	w.Expanded = bool(expanded)

	if child != nil {
		childWidget, ok := child.(Widget)
		if !ok {
			return nil, fmt.Errorf(
				"invalid type for child: %s (expected Widget)",
				child.Type(),
			)
		}
		w.Child = childWidget.AsRenderWidget()
		w.starlarkChild = child
	}

	w.starlarkPad = pad
	switch padVal := pad.(type) {
	case starlark.Int:
		padInt := int(padVal.BigInt().Int64())
		w.Pad.Left = padInt
		w.Pad.Top = padInt
		w.Pad.Right = padInt
		w.Pad.Bottom = padInt
	case starlark.Tuple:
		padList := []starlark.Value(padVal)
		if len(padList) != 4 {
			return nil, fmt.Errorf(
				"pad tuple must hold 4 elements (left, top, right, bottom), found %d",
				len(padList),
			)
		}
		padListInt := make([]starlark.Int, 4)
		for i := 0; i < 4; i++ {
			pi, ok := padList[i].(starlark.Int)
			if !ok {
				return nil, fmt.Errorf("pad element %d is not int", i)
			}
			padListInt[i] = pi
		}
		w.Pad.Left = int(padListInt[0].BigInt().Int64())
		w.Pad.Top = int(padListInt[1].BigInt().Int64())
		w.Pad.Right = int(padListInt[2].BigInt().Int64())
		w.Pad.Bottom = int(padListInt[3].BigInt().Int64())
	default:
		return nil, fmt.Errorf("pad must be int or 4-tuple of int")
	}

	return w, nil
}

func (w *Padding) AsRenderWidget() render.Widget {
	return &w.Padding
}

func (w *Padding) AttrNames() []string {
	return []string{
		"child", "pad", "expanded",
	}
}

func (w *Padding) Attr(name string) (starlark.Value, error) {
	switch name {

	case "expanded":
		return starlark.Bool(w.Expanded), nil

	case "child":
		return w.starlarkChild, nil

	case "pad":
		return w.starlarkPad, nil

	default:
		return nil, nil
	}
}

func (w *Padding) String() string       { return "Padding(...)" }
func (w *Padding) Type() string         { return "Padding" }
func (w *Padding) Freeze()              {}
func (w *Padding) Truth() starlark.Bool { return true }

func (w *Padding) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

type Row struct {
	Widget
	render.Row

	starlarkChildren *starlark.List
}

func newRow(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		main_align  starlark.String
		cross_align starlark.String

		expanded starlark.Bool

		children *starlark.List
	)

	if err := starlark.UnpackArgs(
		"Row",
		args, kwargs,
		"children", &children,
		"main_align?", &main_align,
		"cross_align?", &cross_align,
		"expanded?", &expanded,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Row: %s", err)
	}

	w := &Row{}
	w.MainAlign = main_align.GoString()
	w.CrossAlign = cross_align.GoString()
	w.Expanded = bool(expanded)

	var childrenVal starlark.Value
	childrenIter := children.Iterate()
	defer childrenIter.Done()
	for i := 0; childrenIter.Next(&childrenVal); {
		if _, isNone := childrenVal.(starlark.NoneType); isNone {
			continue
		}

		childrenChild, ok := childrenVal.(Widget)
		if !ok {
			return nil, fmt.Errorf(
				"expected children to be a list of Widget but found: %s (at index %d)",
				childrenVal.Type(),
				i,
			)
		}

		w.Children = append(w.Children, childrenChild.AsRenderWidget())
	}
	w.starlarkChildren = children

	return w, nil
}

func (w *Row) AsRenderWidget() render.Widget {
	return &w.Row
}

func (w *Row) AttrNames() []string {
	return []string{
		"children", "main_align", "cross_align", "expanded",
	}
}

func (w *Row) Attr(name string) (starlark.Value, error) {
	switch name {

	case "main_align":
		return starlark.String(w.MainAlign), nil

	case "cross_align":
		return starlark.String(w.CrossAlign), nil

	case "expanded":
		return starlark.Bool(w.Expanded), nil

	case "children":
		return w.starlarkChildren, nil

	default:
		return nil, nil
	}
}

func (w *Row) String() string       { return "Row(...)" }
func (w *Row) Type() string         { return "Row" }
func (w *Row) Freeze()              {}
func (w *Row) Truth() starlark.Bool { return true }

func (w *Row) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

type Stack struct {
	Widget
	render.Stack

	starlarkChildren *starlark.List
}

func newStack(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		children *starlark.List
	)

	if err := starlark.UnpackArgs(
		"Stack",
		args, kwargs,
		"children", &children,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Stack: %s", err)
	}

	w := &Stack{}

	var childrenVal starlark.Value
	childrenIter := children.Iterate()
	defer childrenIter.Done()
	for i := 0; childrenIter.Next(&childrenVal); {
		if _, isNone := childrenVal.(starlark.NoneType); isNone {
			continue
		}

		childrenChild, ok := childrenVal.(Widget)
		if !ok {
			return nil, fmt.Errorf(
				"expected children to be a list of Widget but found: %s (at index %d)",
				childrenVal.Type(),
				i,
			)
		}

		w.Children = append(w.Children, childrenChild.AsRenderWidget())
	}
	w.starlarkChildren = children

	return w, nil
}

func (w *Stack) AsRenderWidget() render.Widget {
	return &w.Stack
}

func (w *Stack) AttrNames() []string {
	return []string{
		"children",
	}
}

func (w *Stack) Attr(name string) (starlark.Value, error) {
	switch name {

	case "children":
		return w.starlarkChildren, nil

	default:
		return nil, nil
	}
}

func (w *Stack) String() string       { return "Stack(...)" }
func (w *Stack) Type() string         { return "Stack" }
func (w *Stack) Freeze()              {}
func (w *Stack) Truth() starlark.Bool { return true }

func (w *Stack) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

type Text struct {
	Widget
	render.Text

	size *starlark.Builtin
}

func newText(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		content starlark.String
		font    starlark.String

		height starlark.Int
		offset starlark.Int

		color starlark.String
	)

	if err := starlark.UnpackArgs(
		"Text",
		args, kwargs,
		"content", &content,
		"font?", &font,
		"height?", &height,
		"offset?", &offset,
		"color?", &color,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for Text: %s", err)
	}

	w := &Text{}
	w.Content = content.GoString()
	w.Font = font.GoString()
	w.Height = int(height.BigInt().Int64())
	w.Offset = int(offset.BigInt().Int64())

	if color.Len() > 0 {
		c, err := colorful.Hex(color.GoString())
		if err != nil {
			return nil, fmt.Errorf("color is not a valid hex string: %s", color.String())
		}
		w.Color = c
	}

	w.size = starlark.NewBuiltin("size", textSize)

	return w, nil
}

func (w *Text) AsRenderWidget() render.Widget {
	return &w.Text
}

func (w *Text) AttrNames() []string {
	return []string{
		"content", "font", "height", "offset", "color",
	}
}

func (w *Text) Attr(name string) (starlark.Value, error) {
	switch name {

	case "content":
		return starlark.String(w.Content), nil

	case "font":
		return starlark.String(w.Font), nil

	case "height":
		return starlark.MakeInt(w.Height), nil

	case "offset":
		return starlark.MakeInt(w.Offset), nil

	case "color":
		if w.Color == nil {
			return nil, nil
		}
		c, ok := colorful.MakeColor(w.Color)
		if !ok {
			return nil, nil
		}
		return starlark.String(c.Hex()), nil

	case "size":
		return w.size.BindReceiver(w), nil

	default:
		return nil, nil
	}
}

func (w *Text) String() string       { return "Text(...)" }
func (w *Text) Type() string         { return "Text" }
func (w *Text) Freeze()              {}
func (w *Text) Truth() starlark.Bool { return true }

func (w *Text) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}

func textSize(
	thread *starlark.Thread,
	b *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	w := b.Receiver().(*Text)
	width, height := w.Size()

	return starlark.Tuple([]starlark.Value{
		starlark.MakeInt(width),
		starlark.MakeInt(height),
	}), nil
}

type WrappedText struct {
	Widget
	render.WrappedText
}

func newWrappedText(
	thread *starlark.Thread,
	_ *starlark.Builtin,
	args starlark.Tuple,
	kwargs []starlark.Tuple,
) (starlark.Value, error) {

	var (
		content starlark.String
		font    starlark.String

		height      starlark.Int
		width       starlark.Int
		linespacing starlark.Int

		color starlark.String
	)

	if err := starlark.UnpackArgs(
		"WrappedText",
		args, kwargs,
		"content", &content,
		"font?", &font,
		"height?", &height,
		"width?", &width,
		"linespacing?", &linespacing,
		"color?", &color,
	); err != nil {
		return nil, fmt.Errorf("unpacking arguments for WrappedText: %s", err)
	}

	w := &WrappedText{}
	w.Content = content.GoString()
	w.Font = font.GoString()
	w.Height = int(height.BigInt().Int64())
	w.Width = int(width.BigInt().Int64())
	w.LineSpacing = int(linespacing.BigInt().Int64())

	if color.Len() > 0 {
		c, err := colorful.Hex(color.GoString())
		if err != nil {
			return nil, fmt.Errorf("color is not a valid hex string: %s", color.String())
		}
		w.Color = c
	}

	return w, nil
}

func (w *WrappedText) AsRenderWidget() render.Widget {
	return &w.WrappedText
}

func (w *WrappedText) AttrNames() []string {
	return []string{
		"content", "font", "height", "width", "linespacing", "color",
	}
}

func (w *WrappedText) Attr(name string) (starlark.Value, error) {
	switch name {

	case "content":
		return starlark.String(w.Content), nil

	case "font":
		return starlark.String(w.Font), nil

	case "height":
		return starlark.MakeInt(w.Height), nil

	case "width":
		return starlark.MakeInt(w.Width), nil

	case "linespacing":
		return starlark.MakeInt(w.LineSpacing), nil

	case "color":
		if w.Color == nil {
			return nil, nil
		}
		c, ok := colorful.MakeColor(w.Color)
		if !ok {
			return nil, nil
		}
		return starlark.String(c.Hex()), nil

	default:
		return nil, nil
	}
}

func (w *WrappedText) String() string       { return "WrappedText(...)" }
func (w *WrappedText) Type() string         { return "WrappedText" }
func (w *WrappedText) Freeze()              {}
func (w *WrappedText) Truth() starlark.Bool { return true }

func (w *WrappedText) Hash() (uint32, error) {
	sum, err := hashstructure.Hash(w, nil)
	return uint32(sum), err
}
