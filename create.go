package main

//go:generate go run class-build/build.go -sym main.thunk -dir java thunks.Thunk

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)

var manifest = "Manifest-Version: 1.0\r\nMain-Class: thunks.Thunk\r\n"

func errf(f string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, f, args...)
	os.Exit(1)
}

func errln(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func makeThunk(binName, jarName string) error {
    binFile, err := os.Open(binName)
    if err != nil {
        return err
    }
    defer binFile.Close()

    out, err := os.Create(jarName)
    if err != nil {
        return fmt.Errorf("while creating %s: %v", jarName, err)
    }
    defer out.Close()

    wr := zip.NewWriter(out)
    defer wr.Close()

    mf, err := wr.Create("META-INF/MANIFEST.MF")
    if err != nil {
        return err
    }

    if _, err := mf.Write([]byte(manifest)); err != nil {
        return fmt.Errorf("while adding manifest: %v", err)
    }

    thunkFile, err := wr.Create("thunks/Thunk.class")
    if err != nil {
        return err
    }

    if _, err := thunkFile.Write(thunk); err != nil {
        return fmt.Errorf("while adding thunk class: %v", err)
    }

    realExe, err := wr.Create("bins/exefile")
    if err != nil {
        return err
    }

    _, err = io.Copy(realExe, binFile)    

    return err
}

func main() {
    if len(os.Args) < 2 {
        errln("error: not enough arguments, one expected")
    }

    binName := os.Args[1]
    jarName := strings.Split(filepath.Base(binName), ".")[0] + ".jar"

    if err := makeThunk(binName, jarName); err != nil {
        errf("error: %v\n", err)
    }
}