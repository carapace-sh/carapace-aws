package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableFastLaunchCmd = &cobra.Command{
	Use:   "enable-fast-launch",
	Short: "When you enable Windows fast launch for a Windows AMI, images are pre-provisioned, using snapshots to launch instances up to 65% faster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableFastLaunchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableFastLaunchCmd).Standalone()

		ec2_enableFastLaunchCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableFastLaunchCmd.Flags().String("image-id", "", "Specify the ID of the image for which to enable Windows fast launch.")
		ec2_enableFastLaunchCmd.Flags().String("launch-template", "", "The launch template to use when launching Windows instances from pre-provisioned snapshots.")
		ec2_enableFastLaunchCmd.Flags().String("max-parallel-launches", "", "The maximum number of instances that Amazon EC2 can launch at the same time to create pre-provisioned snapshots for Windows fast launch.")
		ec2_enableFastLaunchCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableFastLaunchCmd.Flags().String("resource-type", "", "The type of resource to use for pre-provisioning the AMI for Windows fast launch.")
		ec2_enableFastLaunchCmd.Flags().String("snapshot-configuration", "", "Configuration settings for creating and managing the snapshots that are used for pre-provisioning the AMI for Windows fast launch.")
		ec2_enableFastLaunchCmd.MarkFlagRequired("image-id")
		ec2_enableFastLaunchCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_enableFastLaunchCmd)
}
