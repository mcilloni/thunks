package embedder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

const tplStr = `// automatically generated from @{.fileName}; do not edit
package @{.pkg}

var @{.varName} = [...]byte{@{fmtBytes .bytes}}
`

func fmtByteLine(buf *bytes.Buffer, sl []byte) {
	buf.WriteRune('\t')

	for i, v := range sl {
		if i > 0 {
			buf.WriteRune(' ')
		}

		buf.WriteString(fmt.Sprintf("%d,", v))
	}

	buf.WriteRune('\n')
}

func fmtBytes(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	const ofs = 10

	buf := bytes.NewBufferString("\n")

	for ; len(b) > ofs; b = b[ofs:] {
		fmtByteLine(buf, b[:ofs])
	}

	fmtByteLine(buf, b)

	return buf.String()
}

func Generate(pkg, varName, inPath, outPath string) error {
	in, err := os.Open(inPath)
	if err != nil {
		return err
	}
	defer in.Close()

	inBytes, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	tpl := template.Must(template.New("embedder").Funcs(template.FuncMap{
		"fmtBytes": fmtBytes,
	}).Delims("@{","}").Parse(tplStr))

	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return tpl.Execute(out, map[string]interface{}{
		"pkg":     pkg,
		"fileName": filepath.Base(inPath),
		"varName": varName,
		"bytes":   inBytes,
	})
}
