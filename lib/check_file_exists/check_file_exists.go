package check_file_exists

import (
	"path/filepath"

	flag "github.com/spf13/pflag"
	"github.com/wataru0225/go_nagios_plugins/utils"
)

func Do() {
	fileName := flag.StringP("file", "f", "", "set filename")
	dirName := flag.StringP("dir", "d", "./", "set directory name")
	notExists := flag.BoolP("notexists", "n", false, "set notExists")
	flag.Parse()

	filePattern := filepath.Join(*dirName, *fileName)
	files, err := filepath.Glob(filePattern)

	if err != nil {
		utils.Critical(err.Error()).Exit()
	}

	if *notExists {
		if len(files) == 0 {
			msg := "OK STATUS"
			utils.Ok(msg).Exit()
		} else {
			msg := "CRITICAL ERROR"
			utils.Critical(msg).Exit()
		}
	} else {
		if len(files) != 0 {
			msg := "OK STATUS"
			utils.Ok(msg).Exit()
		} else {
			msg := "CRITICAL ERROR"
			utils.Critical(msg).Exit()
		}
	}
}
