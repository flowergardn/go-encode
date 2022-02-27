package main

import (
	b64 "encoding/base64"
	"github.com/kataras/iris/v12"
)

// checks to see if an error exists, if it does, then exit the application using panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// simple function to encrypt text, rather than writing all this out
func encode(toEncrypt string) string {
	uEnc := b64.URLEncoding.EncodeToString([]byte(toEncrypt))
	return uEnc
}

// function to decode text
func decode(toDecode string) string {
	uDec, _ := b64.URLEncoding.DecodeString(toDecode)
	return string(uDec)
}

func main() {
	// Creates an instance of Iris
	app := iris.New()

	// [GET] / - Sends a hello world :)
	app.Get("/", func(ctx iris.Context) {
		_, err := ctx.Writef("Hello, world!")
		check(err)
	})

	// [GET] /encode - Encodes a string for you
	// Ex: localhost:8080/encode?string=ur mom
	app.Get("/encode", func(ctx iris.Context) {
		// get the string to encode, if it's not specified, default it to lol
		stringToEncode := ctx.URLParamDefault("string", "lol")
		// We don't need the response back, so we set it to null using "_".
		_, err := ctx.JSON(iris.Map{
			"success": true,
			"encoded": encode(stringToEncode),
		})
		// we get the error from above ^, and check to see if there is any errors.
		check(err)
	})

	// [GET] /decode - Decodes a string for you
	// Ex: localhost:8080/decode?string=dXIgbW9t
	app.Get("/decode", func(ctx iris.Context) {
		// get the string to encode, if it's not specified, default it to bG9s (which equals lol in b64)
		stringToDecode := ctx.URLParamDefault("string", "bG9s")
		// We don't need the response back, so we set it to null using "_".
		_, err := ctx.JSON(iris.Map{
			"success": true,
			"decoded": decode(stringToDecode),
		})
		// we get the error from above ^, and check to see if there is any errors.
		check(err)
	})

	// Listen for requests on port 8080
	err := app.Listen(":8080")
	// check for errors
	check(err)
}
