package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed configservice/aws.configservice.get-status.yaml
var configservice_getStatus []byte

//go:embed configservice/aws.configservice.subscribe.yaml
var configservice_subscribe []byte

func init() {
	customizations["config"] = func(cmd *command.Command) error {
		cmd.Name = "configservice"

		for _, spec := range [][]byte{
			configservice_getStatus,
			configservice_subscribe,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		return nil
	}
	customizations["config.put-configuration-recorder"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "recording-group",
			Description: "Specifies which resource types are in scope for the configuration recorder to record",
			Value:       true,
		})
		return nil
	}
}
