package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed emr-containers/aws.emr-containers.delete-role-associations.yaml
var emrContainers_deleteRoleAssociations []byte

//go:embed emr-containers/aws.emr-containers.create-role-associations.yaml
var emrContainers_createRoleAssociations []byte

//go:embed emr-containers/aws.emr-containers.update-role-trust-policy.yaml
var emrContainers_updateRoleAssociations []byte

func init() {
	customizations["emr-containers"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			emrContainers_createRoleAssociations,
			emrContainers_deleteRoleAssociations,
			emrContainers_updateRoleAssociations,
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
