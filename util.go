package main

import (
	"archive/tar"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/packer/packer"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Copies a file by copying the contents of the file to another place.
func CopyContents(dst, src string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()

	dstF, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstF.Close()

	if _, err := io.Copy(dstF, srcF); err != nil {
		return err
	}

	return nil
}

// DirToBox takes the directory and compresses it into a Vagrant-compatible
// box. This function does not perform checks to verify that dir is
// actually a proper box. This is an expected precondition.
func DirToBox(dst, dir string, ui packer.Ui, level int) error {
	log.Printf("Turning dir into box: %s => %s", dir, dst)
	dstF, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstF.Close()

	var dstWriter io.Writer = dstF
	if level != flate.NoCompression {
		log.Printf("Compressing with gzip compression level: %d", level)
		gzipWriter, err := gzip.NewWriterLevel(dstWriter, level)
		if err != nil {
			return err
		}
		defer gzipWriter.Close()

		dstWriter = gzipWriter
	}

	tarWriter := tar.NewWriter(dstWriter)
	defer tarWriter.Close()

	// This is the walk func that tars each of the files in the dir
	tarWalk := func(path string, info os.FileInfo, prevErr error) error {
		// If there was a prior error, return it
		if prevErr != nil {
			return prevErr
		}

		// Skip directories
		if info.IsDir() {
			log.Printf("Skipping directory '%s' for box '%s'", path, dst)
			return nil
		}

		log.Printf("Box add: '%s' to '%s'", path, dst)
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		// We have to set the Name explicitly because it is supposed to
		// be a relative path to the root. Otherwise, the tar ends up
		// being a bunch of files in the root, even if they're actually
		// nested in a dir in the original "dir" param.
		header.Name, err = filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		if ui != nil {
			ui.Message(fmt.Sprintf("Compressing: %s", header.Name))
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if _, err := io.Copy(tarWriter, f); err != nil {
			return err
		}

		return nil
	}

	// Tar.gz everything up
	return filepath.Walk(dir, tarWalk)
}

// WriteMetadata writes the "metadata.json" file for a Vagrant box.
func WriteMetadata(dir string, contents interface{}) error {
	f, err := os.Create(filepath.Join(dir, "metadata.json"))
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(contents)
}

// func RunAndLog(cmd *exec.Cmd) (string, string, error) {
// 	var stdout, stderr bytes.Buffer

// 	log.Printf("Executing: %s %v", cmd.Path, cmd.Args[1:])
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr
// 	err := cmd.Run()

// 	stdoutString := strings.TrimSpace(stdout.String())
// 	stderrString := strings.TrimSpace(stderr.String())

// 	if _, ok := err.(*exec.ExitError); ok {
// 		err = fmt.Errorf("VMware error: %s", stderrString)
// 	}

// 	log.Printf("stdout: %s", stdoutString)
// 	log.Printf("stderr: %s", stderrString)

// 	// Replace these for Windows, we only want to deal with Unix
// 	// style line endings.
// 	returnStdout := strings.Replace(stdout.String(), "\r\n", "\n", -1)
// 	returnStderr := strings.Replace(stderr.String(), "\r\n", "\n", -1)

// 	return returnStdout, returnStderr, err
// }

func FindOvfTool() (ovftool string, err error) {
	// Accumulate any errors
	errs := new(packer.MultiError)

	// use ovftool in PATH, so use can decide which one to use
	ovftool = "ovftool"
	if _, err := exec.LookPath(ovftool); err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("ovftool not found in path: %s", err))

		files := make([]string, 0, 6)

		// search ovftool at some specific places
		files = append(files, "/Applications/VMware Fusion.app/Contents/Library/VMware OVF Tool/ovftool")

		if os.Getenv("ProgramFiles(x86)") != "" {
			files = append(files,
				filepath.Join(os.Getenv("ProgramFiles(x86)"), "/VMware/Client Integration Plug-in 5.5/ovftool/ovftool.exe"))
		}

		if os.Getenv("ProgramFiles") != "" {
			files = append(files,
				filepath.Join(os.Getenv("ProgramFiles"), "/VMware/Client Integration Plug-in 5.5/ovftool/ovftool.exe"))
		}

		if os.Getenv("ProgramFiles(x86)") != "" {
			files = append(files,
				filepath.Join(os.Getenv("ProgramFiles(x86)"), "/VMware/VMware Workstation/ovftool/ovftool.exe"))
		}

		if os.Getenv("ProgramFiles") != "" {
			files = append(files,
				filepath.Join(os.Getenv("ProgramFiles"), "/VMware/VMware Workstation/ovftool/ovftool.exe"))
		}

		file := findFile(files)
		if file == "" {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("ovftool not found: %s", err))
		} else {
			ovftool = file
		}
	}
	return
}

func findFile(files []string) string {
	for _, file := range files {
		file = normalizePath(file)
		log.Printf("Searching for file '%s'", file)

		if _, err := os.Stat(file); err == nil {
			log.Printf("Found file '%s'", file)
			return file
		}
	}

	log.Printf("File not found")
	return ""
}

func normalizePath(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	path = strings.Replace(path, "//", "/", -1)
	path = strings.TrimRight(path, "/")
	return path
}
