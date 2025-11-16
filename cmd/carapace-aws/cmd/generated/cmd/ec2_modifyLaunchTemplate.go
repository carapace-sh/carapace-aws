package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyLaunchTemplateCmd = &cobra.Command{
	Use:   "modify-launch-template",
	Short: "Modifies a launch template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyLaunchTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyLaunchTemplateCmd).Standalone()

		ec2_modifyLaunchTemplateCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
		ec2_modifyLaunchTemplateCmd.Flags().String("default-version", "", "The version number of the launch template to set as the default version.")
		ec2_modifyLaunchTemplateCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyLaunchTemplateCmd.Flags().String("launch-template-id", "", "The ID of the launch template.")
		ec2_modifyLaunchTemplateCmd.Flags().String("launch-template-name", "", "The name of the launch template.")
		ec2_modifyLaunchTemplateCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyLaunchTemplateCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyLaunchTemplateCmd)
}
