package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed eks/aws.eks.get-token.yaml
var eks_getToken []byte

//go:embed eks/aws.eks.update-kubeconfig.yaml
var eks_updateKubeconfig []byte

func init() {
	customizations["eks"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			eks_getToken,
			eks_updateKubeconfig,
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
