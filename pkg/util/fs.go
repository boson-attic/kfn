package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func FileExist(p ...string) bool {
	if s, err := os.Stat(path.Join(p...)); os.IsNotExist(err) {
		return false
	} else {
		return !s.IsDir()
	}
}

func DirExist(p ...string) bool {
	if s, err := os.Stat(path.Join(p...)); os.IsNotExist(err) {
		return false
	} else {
		return s.IsDir()
	}
}

func FsExist(p ...string) bool {
	if _, err := os.Stat(path.Join(p...)); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func MkdirpIfNotExists(path string) error {
	if path != "" && path != "." && !FsExist(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

type WriteDest struct {
	Filename string
	Data     []byte
}

func WriteFiles(destDir string, dest ...WriteDest) error {
	for _, d := range dest {

		outFile, err := os.OpenFile(path.Join(destDir, d.Filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}

		_, err = outFile.Write(d.Data)
		if err != nil {
			return err
		}

		outFile.Close()
	}
	return nil
}

func PipeTemplateToFile(destination string, template *template.Template, data interface{}) error {
	outFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	err = template.Execute(outFile, data)
	if err != nil {
		return err
	}

	outFile.Close()

	return nil
}

func Copy(source string, dest string) error {
	if !FsExist(source) {
		return fmt.Errorf("Cannot find %s", source)
	}

	cmd := exec.Command("cp", "-r", source, dest)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Cannot copy %s to %s", source, dest)
	}
	return nil
}

func CopyContent(source string, dest string) error {
	if !FsExist(source) {
		return fmt.Errorf("Cannot find %s", source)
	}

	cmd := exec.Command("cp", "-r", path.Join(source, "."), dest)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Cannot copy %s to %s", source, dest)
	}
	return nil
}

func CopyOrLink(link bool) func(string, string) error {
	if link {
		return func(source string, dest string) error {
			if !FsExist(dest) {
				return os.Link(source, dest)
			} else {
				return nil
			}
		}
	} else {
		return Copy
	}
}

func RmR(path string) error {
	if !FsExist(path) {
		return nil
	}

	cmd := exec.Command("rm", "-r", path)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
