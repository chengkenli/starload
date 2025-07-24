package application

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math"
	"starload/util"
)

// WatermarkContainer 带平铺水印的容器
type WatermarkContainer struct {
	widget.BaseWidget
	content    fyne.CanvasObject
	text       string
	angle      float32
	spacing    float32
	fontSize   float32
	opacity    uint8
	watermarks []fyne.CanvasObject
}

func NewWatermarkContainer(content fyne.CanvasObject) *WatermarkContainer {
	w := &WatermarkContainer{
		content:  content,
		text:     fmt.Sprintf("Starload©v%0.1f", util.Segment),
		angle:    30,
		spacing:  200,
		fontSize: 25, /*字体大小*/
		opacity:  40, /*透明度*/
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *WatermarkContainer) CreateRenderer() fyne.WidgetRenderer {
	return &watermarkRenderer{w: w}
}

type watermarkRenderer struct {
	w *WatermarkContainer
}

func (r *watermarkRenderer) Layout(size fyne.Size) {
	// 布局主内容
	r.w.content.Resize(size)

	// 计算水印布局参数
	rad := float64(r.w.angle) * math.Pi / 180
	colSpacing := float32(float64(r.w.spacing) / math.Cos(rad))
	// 水印行间距50
	rowSpacing := float32(float64(50) / math.Sin(rad))

	cols := int(size.Width/colSpacing) + 3
	rows := int(size.Height/rowSpacing) + 3

	// 动态创建水印对象
	if len(r.w.watermarks) < cols*rows {
		for i := len(r.w.watermarks); i < cols*rows; i++ {
			text := canvas.NewText(r.w.text, color.NRGBA{R: 200, G: 200, B: 200, A: r.w.opacity})
			text.TextSize = r.w.fontSize
			text.TextStyle = fyne.TextStyle{Italic: true}
			r.w.watermarks = append(r.w.watermarks, text)
		}
	}

	// 定位水印
	for y := -1; y < rows; y++ {
		for x := -1; x < cols; x++ {
			idx := (y+1)*cols + (x + 1)
			if idx >= len(r.w.watermarks) {
				continue
			}

			posX := float32(x)*colSpacing - colSpacing/2
			posY := float32(y)*rowSpacing - rowSpacing/2

			text := r.w.watermarks[idx].(*canvas.Text)
			text.Move(fyne.NewPos(posX, posY))
		}
	}
}

func (r *watermarkRenderer) MinSize() fyne.Size {
	return r.w.content.MinSize()
}

func (r *watermarkRenderer) Refresh() {
	r.w.content.Refresh()
	for _, wm := range r.w.watermarks {
		wm.Refresh()
	}
}

func (r *watermarkRenderer) Objects() []fyne.CanvasObject {
	return append([]fyne.CanvasObject{r.w.content}, r.w.watermarks...)
}

func (r *watermarkRenderer) Destroy() {}
