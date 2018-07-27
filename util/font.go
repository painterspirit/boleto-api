package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"

	"github.com/golang/freetype"
)

type font struct {
	FtFont *truetype.Font
}

var fontStr font

func GetFont() font {
	if (font{}) == fontStr {
		fontBytes, err := ioutil.ReadFile("./boleto/Arial.ttf")		
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}

		f, err := freetype.ParseFont(fontBytes)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}

		fontStr = font{
			FtFont: f,
		}
	}

	return fontStr
}
