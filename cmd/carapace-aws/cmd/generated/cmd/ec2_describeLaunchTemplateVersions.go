package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeLaunchTemplateVersionsCmd = &cobra.Command{
	Use:   "describe-launch-template-versions",
	Short: "Describes one or more versions of a specified launch template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeLaunchTemplateVersionsCmd).Standalone()

	ec2_describeLaunchTemplateVersionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().String("launch-template-id", "", "The ID of the launch template.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().String("launch-template-name", "", "The name of the launch template.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().String("max-version", "", "The version number up to which to describe launch template versions.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().String("min-version", "", "The version number after which to describe launch template versions.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().Bool("no-resolve-alias", false, "If `true`, and if a Systems Manager parameter is specified for `ImageId`, the AMI ID is displayed in the response for `imageId`.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().Bool("resolve-alias", false, "If `true`, and if a Systems Manager parameter is specified for `ImageId`, the AMI ID is displayed in the response for `imageId`.")
	ec2_describeLaunchTemplateVersionsCmd.Flags().String("versions", "", "One or more versions of the launch template.")
	ec2_describeLaunchTemplateVersionsCmd.Flag("no-dry-run").Hidden = true
	ec2_describeLaunchTemplateVersionsCmd.Flag("no-resolve-alias").Hidden = true
	ec2Cmd.AddCommand(ec2_describeLaunchTemplateVersionsCmd)
}
