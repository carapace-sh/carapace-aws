package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteLaunchTemplateCmd = &cobra.Command{
	Use:   "delete-launch-template",
	Short: "Deletes a launch template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteLaunchTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteLaunchTemplateCmd).Standalone()

		ec2_deleteLaunchTemplateCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLaunchTemplateCmd.Flags().String("launch-template-id", "", "The ID of the launch template.")
		ec2_deleteLaunchTemplateCmd.Flags().String("launch-template-name", "", "The name of the launch template.")
		ec2_deleteLaunchTemplateCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteLaunchTemplateCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteLaunchTemplateCmd)
}
