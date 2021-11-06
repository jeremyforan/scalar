package Scalar

type Scaler struct {

	Height 	float32
	Width 	float32

	Rows 	float32
	Columns float32
}

func (s *Scaler) ColumnWidth() float32 {
	return s.Width / s.Columns
}

func (s *Scaler) RowHeight() float32 {
	return s.Height / s.Rows
}

func (s *Scaler) MakePixelRelative(v Point) Point {

	bufferedWidth 	:=  (s.ColumnWidth() * 2 ) + s.Width
	bufferedHeight 	:=  (s.RowHeight() * 2 ) + s.Height

	widthRatio 		:= bufferedWidth  / s.Width
	heightRatio 	:= bufferedHeight / s.Height

	widthPercent := v.X 	/  bufferedWidth
	heightPercent := v.Y 	/ bufferedHeight

	widthScaled 	:= widthPercent * (widthRatio * 2)
	heightScaled 	:= heightPercent * (heightRatio * 2)

	widthAdjusted 	:= widthScaled - widthRatio
	heightAdjusted 	:= heightScaled - heightRatio

	vector := Point{
		X: widthAdjusted ,
		Y: heightAdjusted,
	}

	return vector
}

type Point struct {
	X,Y float32
}

func (v Point) Subtract(p Point) Point {
	output := Point{
		X: v.X - p.X,
		Y: v.Y - p.Y,
	}
	return output
}


