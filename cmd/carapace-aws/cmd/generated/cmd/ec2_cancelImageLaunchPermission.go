package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelImageLaunchPermissionCmd = &cobra.Command{
	Use:   "cancel-image-launch-permission",
	Short: "Removes your Amazon Web Services account from the launch permissions for the specified AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelImageLaunchPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_cancelImageLaunchPermissionCmd).Standalone()

		ec2_cancelImageLaunchPermissionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelImageLaunchPermissionCmd.Flags().String("image-id", "", "The ID of the AMI that was shared with your Amazon Web Services account.")
		ec2_cancelImageLaunchPermissionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_cancelImageLaunchPermissionCmd.MarkFlagRequired("image-id")
		ec2_cancelImageLaunchPermissionCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_cancelImageLaunchPermissionCmd)
}
