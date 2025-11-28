package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed gamelift/aws.gamelift.get-game-session-log.yaml
var gamelift_getGameSessionLog []byte

//go:embed gamelift/aws.gamelift.upload-build.yaml
var gamelift_uploadBuild []byte

func init() {
	customizations["gamelift"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			gamelift_getGameSessionLog,
			gamelift_uploadBuild,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		return nil
	}
}
