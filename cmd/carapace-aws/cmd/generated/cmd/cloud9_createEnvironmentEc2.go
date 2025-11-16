package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_createEnvironmentEc2Cmd = &cobra.Command{
	Use:   "create-environment-ec2",
	Short: "Creates an Cloud9 development environment, launches an Amazon Elastic Compute Cloud (Amazon EC2) instance, and then connects from the instance to the environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_createEnvironmentEc2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_createEnvironmentEc2Cmd).Standalone()

		cloud9_createEnvironmentEc2Cmd.Flags().String("automatic-stop-time-minutes", "", "The number of minutes until the running instance is shut down after the environment has last been used.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("client-request-token", "", "A unique, case-sensitive string that helps Cloud9 to ensure this operation completes no more than one time.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("connection-type", "", "The connection type used for connecting to an Amazon EC2 environment.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("description", "", "The description of the environment to create.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("dry-run", "", "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("image-id", "", "The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("instance-type", "", "The type of instance to connect to the environment (for example, `t2.micro`).")
		cloud9_createEnvironmentEc2Cmd.Flags().String("name", "", "The name of the environment to create.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("owner-arn", "", "The Amazon Resource Name (ARN) of the environment owner.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("subnet-id", "", "The ID of the subnet in Amazon VPC that Cloud9 will use to communicate with the Amazon EC2 instance.")
		cloud9_createEnvironmentEc2Cmd.Flags().String("tags", "", "An array of key-value pairs that will be associated with the new Cloud9 development environment.")
		cloud9_createEnvironmentEc2Cmd.MarkFlagRequired("image-id")
		cloud9_createEnvironmentEc2Cmd.MarkFlagRequired("instance-type")
		cloud9_createEnvironmentEc2Cmd.MarkFlagRequired("name")
	})
	cloud9Cmd.AddCommand(cloud9_createEnvironmentEc2Cmd)
}
