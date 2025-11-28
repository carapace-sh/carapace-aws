package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed servicecatalog/aws.servicecatalog.generate.yaml
var servicecatalog_generate []byte

func init() {
	customizations["servicecatalog"] = func(cmd *command.Command) error {
		var specCommand command.Command
		if err := yaml.Unmarshal(servicecatalog_generate, &specCommand); err != nil {
			return err
		}
		cmd.Commands = append(cmd.Commands, specCommand)
		return nil
	}
}
