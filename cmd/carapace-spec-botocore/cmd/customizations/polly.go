package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["polly.start-speech-synthesis-stream"] = func(cmd *command.Command) error {
		return &SkipError{}
	}
}
