package lazyinit

import (
	"fmt"
	"runtime"
	"strings"
)

func Powershell(completers []string) string {
	snippet := `%v%v

$_carapace_lazy = {
    param($wordToComplete, $commandAst, $cursorPosition)
    $completer = $commandAst.CommandElements[0].Value
    carapace $completer powershell | Out-String | Invoke-Expression
    & (Get-Item "Function:_${completer}_completer") $wordToComplete $commandAst $cursorPosition
}
%v
`

	prefix := "# "
	if runtime.GOOS == "windows" {
		prefix = ""
	}

	complete := make([]string, 0, len(completers)*2)
	for _, completer := range completers {
		complete = append(complete, fmt.Sprintf(`Register-ArgumentCompleter -Native -CommandName '%v'     -ScriptBlock $_carapace_lazy`, completer))
		complete = append(complete, fmt.Sprintf(`%vRegister-ArgumentCompleter -Native -CommandName '%v.exe' -ScriptBlock $_carapace_lazy`, prefix, completer))
	}
	return fmt.Sprintf(snippet, pathSnippet("powershell"), envSnippet("powershell"), strings.Join(complete, "\n"))
}
