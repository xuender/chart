package draw

import "github.com/tdewolff/canvas"

func Text(ctx *canvas.Context, text string, face *canvas.FontFace, rect canvas.Rect) {
	rich := canvas.NewRichText(face)
	rich.WriteString(text)

	ctx.DrawText(
		rect.X, rect.Y,
		rich.ToText(rect.W, rect.H, canvas.Center, canvas.Center, 0.0, 0.0),
	)
}
