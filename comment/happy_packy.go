package comment

import (
	"bytes"
	"text/template"
)

type CommentFormatter struct {
	Comment string
}

func CreatePackerFriend(comment string) (string, error) {
	// TODO: Make packer-friend resilient to long text strings that need to
	// be broken into multiple lines.

	PackerSay := `
     |\
     |  \
\    |    \
| \  |      \
|   \ \       \
|     \ \       \
|       \ \       \
|        | | ◕  ◕  *
|        | |        |
|        | |    o   |   *{{.Comment}}*
|        | |        |  /
|        | |        |
|        | |   ヮ   /
|        | |______-
\        |
  \      |
    \    |
      \  |
        \|
`

	tpl := template.Must(template.New("createVM").Parse(PackerSay))
	var b bytes.Buffer
	c := CommentFormatter{comment}

	err := tpl.Execute(&b, c)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
