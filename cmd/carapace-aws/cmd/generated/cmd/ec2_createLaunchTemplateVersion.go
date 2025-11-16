package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createLaunchTemplateVersionCmd = &cobra.Command{
	Use:   "create-launch-template-version",
	Short: "Creates a new version of a launch template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createLaunchTemplateVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createLaunchTemplateVersionCmd).Standalone()

		ec2_createLaunchTemplateVersionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
		ec2_createLaunchTemplateVersionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createLaunchTemplateVersionCmd.Flags().String("launch-template-data", "", "The information for the launch template.")
		ec2_createLaunchTemplateVersionCmd.Flags().String("launch-template-id", "", "The ID of the launch template.")
		ec2_createLaunchTemplateVersionCmd.Flags().String("launch-template-name", "", "The name of the launch template.")
		ec2_createLaunchTemplateVersionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createLaunchTemplateVersionCmd.Flags().Bool("no-resolve-alias", false, "If `true`, and if a Systems Manager parameter is specified for `ImageId`, the AMI ID is displayed in the response for `imageID`.")
		ec2_createLaunchTemplateVersionCmd.Flags().Bool("resolve-alias", false, "If `true`, and if a Systems Manager parameter is specified for `ImageId`, the AMI ID is displayed in the response for `imageID`.")
		ec2_createLaunchTemplateVersionCmd.Flags().String("source-version", "", "The version of the launch template on which to base the new version.")
		ec2_createLaunchTemplateVersionCmd.Flags().String("version-description", "", "A description for the version of the launch template.")
		ec2_createLaunchTemplateVersionCmd.MarkFlagRequired("launch-template-data")
		ec2_createLaunchTemplateVersionCmd.Flag("no-dry-run").Hidden = true
		ec2_createLaunchTemplateVersionCmd.Flag("no-resolve-alias").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createLaunchTemplateVersionCmd)
}
