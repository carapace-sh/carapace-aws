package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed ecs/aws.ecs.deploy.yaml
var ecs_deploy []byte

//go:embed ecs/aws.ecs.monitor-express-gateway-service.yaml
var ecs_monitorExpressGatewayService []byte

func init() {
	customizations["ecs"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			ecs_deploy,
			ecs_monitorExpressGatewayService,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		return nil
	}

	customizations["ecs.update-express-gateway-service"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "monitor-resources",
			Description: " Enable live monitoring of service resource status",
			Value:       true,
		})
		return nil
	}
	customizations["ecs.delete-express-gateway-service"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "monitor-resources",
			Description: " Enable live monitoring of service resource status",
			Value:       true,
		})
		return nil
	}
	customizations["ecs.create-express-gateway-service"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "monitor-resources",
			Description: " Enable live monitoring of service resource status",
			Value:       true,
		})
		return nil
	}
}
