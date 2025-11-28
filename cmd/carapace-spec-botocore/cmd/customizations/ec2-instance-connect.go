package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed ec2-instance-connect/aws.ec2-instance-connect.open-tunnel.yaml
var ec2InstanceConnect_openTunnel []byte

//go:embed ec2-instance-connect/aws.ec2-instance-connect.ssh.yaml
var ec2InstanceConnect_ssh []byte

func init() {
	customizations["ec2-instance-connect"] = func(cmd *command.Command) error {
		for _, spec := range [][]byte{
			ec2InstanceConnect_openTunnel,
			ec2InstanceConnect_ssh,
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
