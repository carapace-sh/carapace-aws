package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed cloudtrail/aws.cloudtrail.verify-query-results.yaml
var cloudTrail_verifyQueryResults []byte

//go:embed cloudtrail/aws.cloudtrail.validate-logs.yaml
var cloudTrail_validateLogs []byte

func init() {
	customizations["cloudtrail"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			cloudTrail_verifyQueryResults,
			cloudTrail_validateLogs,
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
