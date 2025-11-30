package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
	"gopkg.in/yaml.v3"
)

//go:embed deploy/aws.deploy.deregister.yaml
var deploy_deregister []byte

//go:embed deploy/aws.deploy.install.yaml
var deploy_install []byte

//go:embed deploy/aws.deploy.push.yaml
var deploy_push []byte

//go:embed deploy/aws.deploy.register.yaml
var deploy_register []byte

//go:embed deploy/aws.deploy.uninstall.yaml
var deploy_uninstall []byte

func init() {
	customizations["codedeploy"] = func(cmd *command.Command) error {
		cmd.Name = "deploy"

		for _, spec := range [][]byte{
			deploy_deregister,
			deploy_install,
			deploy_push,
			deploy_register,
			deploy_uninstall,
		} {
			var specCommand command.Command
			if err := yaml.Unmarshal(spec, &specCommand); err != nil {
				return err
			}
			cmd.Commands = append(cmd.Commands, specCommand)
		}
		return nil
	}

	for _, operation := range []string{
		"codedeploy.create-deployment",
		"codedeploy.get-application-revision",
		"codedeploy.register-application-revision",
	} {
		customizations[operation] = func(cmd *command.Command) error {
			cmd.AddFlag(command.Flag{
				Longhand:    "s3-location",
				Description: "Information about the location of the application revision in Amazon S3",
				Value:       true,
			})
			cmd.AddFlag(command.Flag{
				Longhand:    "github-location",
				Description: "Information about the location of the application revision in GitHub",
				Value:       true,
			})
			return nil
		}
	}
}
