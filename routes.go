package main

import (
	"fmt"

	"net/http"
	"os"

	"image"
	_ "image/jpeg"
	"image/png"
	"github.com/julienschmidt/httprouter"
)

//Generate and return one card given the background image and csv
func previewCard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Inside previewCard")
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println("Joshing around")
		fmt.Println(err)
	}

	for x, p := range r.MultipartForm.File {
		fmt.Println(x)

		fmt.Println("------------------------------------------------------------")
		fmt.Println(p)
		fmt.Println("************************************************************")
	}

	fmt.Println(r.MultipartForm.File["background"])
	//getting the multipart file, it comes in an array of objects, not sure why will need to test on future bigger images
	k := r.MultipartForm.File["background"][0]
	fmt.Println(k.Filename)
	fmt.Println(k.Header)
	fmt.Println(k.Size)
	l, err := k.Open()
	img, name, err := image.Decode(l)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(name)


	cardName := fmt.Sprintf("opp.png")
	f, _ := os.Create(cardName)
	png.Encode(f, img)
	/* fmt.Println(r.Form[pp])
	_,b,c := r.FormFile("background")
	if c != nil{
		fmt.Println("didnt get image correctly")
	}
	fmt.Println(b.Filename) */
	/* 	fmt.Println("------------------------------------------------------------")
	   	fmt.Println("------------------------------------------------------------")
	   	fmt.Println(r.Form["filename"])
	   	fmt.Println("------------------------------------------------------------")
	   	fmt.Println(r.PostFormValue("filename"))
	   	fmt.Println("------------------------------------------------------------")
	   	fmt.Println(r.PostFormValue("background")) */

	/* img, err := png.Decode(bytes.NewReader(r.Form["background"]))
	if err != nil {
		fmt.Println(err)
	}
	cardName := fmt.Sprintf("opp.png")
	f, _ := os.Create(cardName)
	png.Encode(f, img) */
}
