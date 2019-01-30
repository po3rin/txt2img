# txt2img

<img src="https://img.shields.io/badge/go-v1.11-blue.svg"/>

Fill in the image with the letter using Go.

## Usage

```go
func main() {
  // ...

  // using txt2img package
  d, _ := txt2img.NewDrawer(
    txt2img.Params{
      Img:               img,
      FontSize:          34,
      Font:              font,
      TextColor:         color,
      TextPosHorizontal: 30,
      TextPosVertical:   30,
    },
  )
  result, _ := d.Draw(text)
  // ...
}
```
