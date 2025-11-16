package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteLaunchTemplateVersionsCmd = &cobra.Command{
	Use:   "delete-launch-template-versions",
	Short: "Deletes one or more versions of a launch template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteLaunchTemplateVersionsCmd).Standalone()

	ec2_deleteLaunchTemplateVersionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteLaunchTemplateVersionsCmd.Flags().String("launch-template-id", "", "The ID of the launch template.")
	ec2_deleteLaunchTemplateVersionsCmd.Flags().String("launch-template-name", "", "The name of the launch template.")
	ec2_deleteLaunchTemplateVersionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteLaunchTemplateVersionsCmd.Flags().String("versions", "", "The version numbers of one or more launch template versions to delete.")
	ec2_deleteLaunchTemplateVersionsCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteLaunchTemplateVersionsCmd.MarkFlagRequired("versions")
	ec2Cmd.AddCommand(ec2_deleteLaunchTemplateVersionsCmd)
}
