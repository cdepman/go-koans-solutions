package go_koans

import (
	"bytes"
	"fmt"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		out.ReadFrom(in)
		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		portion, err := in.ReadBytes('o')
		if err != nil {
		}
		out.Write(portion)
		fmt.Printf("%s", out.String())
		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
