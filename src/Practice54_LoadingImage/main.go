package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
   "image/png"
	"os"
)

func main(){
  f, err := os.Open("./moon.jpg")
  fmt.Print("file open----",f)
if err != nil {
    fmt.Print(err)
}
defer f.Close()

img, fmtName, err := image.Decode(f)
if err != nil {
  fmt.Print("Error",err)
}
fmt.Print("img",img)
// fmt.Print("img",img)
fmt.Print("fmtName",fmtName)


//Encoding(Saving the image)
fn, err := os.Create("./outimage.jpg")
if err != nil {
    // Handle error
}
defer fn.Close()

// Specify the quality, between 0-100
// Higher is better
opt := jpeg.Options{
    Quality: 90,
}
err = jpeg.Encode(fn, img, &opt)
if err != nil {
    // Handle error
}

//Encoding(Saving the image in png)
fpng, err := os.Create("pngimage.png")
if err != nil {
    // Handle error
}
defer f.Close()

// Encode to `PNG` with `DefaultCompression` level
// then save to file
err = png.Encode(fpng, img)
if err != nil {
    // Handle error
}


}