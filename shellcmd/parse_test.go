package shellcmd

import (
	"testing"

	"github.com/matryer/is"
)

func TestParse(t *testing.T) {
	tests := map[string][]string{
		`go build .`:                 {`go`, `build`, `.`},
		``:                           {``},
		`ls`:                         {`ls`},
		`cmd --str ""`:               {`cmd`, `--str`, ``},
		`cmd --escaped \"`:           {`cmd`, `--escaped`, `"`},
		`cmd --single ''`:            {`cmd`, `--single`, ``},
		`cmd --singleEscaped \'`:     {`cmd`, `--singleEscaped`, `'`},
		`cmd "with space"`:           {`cmd`, `with space`},
		`cmd 'single with	tab'`:      {`cmd`, `single with	tab`},
		`cmd "single ' in double"`:   {`cmd`, `single ' in double`},
		`cmd 'double " in single'`:   {`cmd`, `double " in single`},
		`cmd "escaped \' in double"`: {`cmd`, `escaped \' in double`},
		`cmd 'escaped \" in single'`: {`cmd`, `escaped \" in single`},
		`cmd "double \" in double"`:  {`cmd`, `double " in double`},
		`cmd 'single \' in single'`:  {`cmd`, `single ' in single`},
		`cmd partially" quo"ted`:     {`cmd`, `partially quoted`},
		`cmd partially' sing'le`:     {`cmd`, `partially single`},
		`cmd \c\h\a\r \e\s\c\a\p\e`:  {`cmd`, `\c\h\a\r`, `\e\s\c\a\p\e`},
		`cmd space\ escape`:          {`cmd`, `space escape`},
		`cmd tab\	escape`:            {`cmd`, `tab	escape`},
		`cmd "double space\ escape"`: {`cmd`, `double space\ escape`},
		`cmd 'single space\ escape'`: {`cmd`, `single space\ escape`},
		`cmd ending\`:                {`cmd`, `ending\`},
		`cmd backslash \\`:           {`cmd`, `backslash`, `\`},
		` cmd extra   space `:        {`cmd`, `extra`, `space`},
	}

	for cmd, res := range tests {
		t.Run(cmd, func(t *testing.T) {
			is := is.New(t)

			args, err := new(cmdParser).parse(cmd)
			is.NoErr(err)
			is.Equal(args, res) // Wrong args returned
		})
	}
}

func TestParse_unterminatedQuote(t *testing.T) {
	is := is.New(t)

	_, err := new(cmdParser).parse(`cmd "unterminated`)
	is.Equal(err.Error(), `shellcmd: unterminated quote in command: cmd "unterminated`)
}

func TestParse_unterminatedSingleQuote(t *testing.T) {
	is := is.New(t)

	_, err := new(cmdParser).parse(`cmd 'unterminated`)
	is.Equal(err.Error(), `shellcmd: unterminated quote in command: cmd 'unterminated`)
}
