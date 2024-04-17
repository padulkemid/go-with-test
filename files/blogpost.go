package blogpost

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleTag = "Title: "
	descTag  = "Description: "
	tagsTag  = "Tags: "
)

func readLine(s *bufio.Scanner, tag string) string {
	s.Scan()
	return strings.TrimPrefix(s.Text(), tag)
}

func readBody(s *bufio.Scanner) string {
	s.Scan() // Scans "---" and get to the body

	bf := bytes.Buffer{}

	for s.Scan() {
		fmt.Fprintln(&bf, s.Text())
	}

	return strings.TrimSuffix(bf.String(), "\n")
}

func readPost(pf io.Reader) (Post, error) {
	s := bufio.NewScanner(pf)

	title := readLine(s, titleTag)
	desc := readLine(s, descTag)
	tag := readLine(s, tagsTag)
	body := readBody(s)

	post := Post{
		Title:       title,
		Description: desc,
		Tags:        strings.Split(tag, ", "),
		Body:        body,
	}

	return post, nil
}

func getPost(fileSystem fs.FS, f string) (Post, error) {
	pf, err := fileSystem.Open(f)
	if err != nil {
		return Post{}, err
	}

	defer pf.Close()

	return readPost(pf)
}

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	var posts []Post

	dirs, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	for _, f := range dirs {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
