package javac

import (
    "fmt"
    "os/exec"
)

type Compiler struct {
    Exe string
    ClassDestDir string
    ClassPath string
}

func (jc *Compiler) Compile(sourceFiles []string) error {
    exe := jc.Exe
    if exe == "" {
        exe = "javac"
    }

    args := []string{}

    if jc.ClassDestDir != "" {
        args = append(args, "-d", jc.ClassDestDir)
    }

    if jc.ClassPath != "" {
        args = append(args, "-cp", jc.ClassPath)
    }

    args = append(args, sourceFiles...)

    cmd := exec.Command(exe, args...)

    if err := cmd.Run(); err != nil {
        exitErr := err.(*exec.ExitError)
        
        return fmt.Errorf("javac failed: %s", string(exitErr.Stderr))
    }

    return nil
}