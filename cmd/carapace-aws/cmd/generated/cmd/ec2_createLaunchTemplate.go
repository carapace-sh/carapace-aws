package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createLaunchTemplateCmd = &cobra.Command{
	Use:   "create-launch-template",
	Short: "Creates a launch template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createLaunchTemplateCmd).Standalone()

	ec2_createLaunchTemplateCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
	ec2_createLaunchTemplateCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createLaunchTemplateCmd.Flags().String("launch-template-data", "", "The information for the launch template.")
	ec2_createLaunchTemplateCmd.Flags().String("launch-template-name", "", "A name for the launch template.")
	ec2_createLaunchTemplateCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createLaunchTemplateCmd.Flags().String("operator", "", "Reserved for internal use.")
	ec2_createLaunchTemplateCmd.Flags().String("tag-specifications", "", "The tags to apply to the launch template on creation.")
	ec2_createLaunchTemplateCmd.Flags().String("version-description", "", "A description for the first version of the launch template.")
	ec2_createLaunchTemplateCmd.MarkFlagRequired("launch-template-data")
	ec2_createLaunchTemplateCmd.MarkFlagRequired("launch-template-name")
	ec2_createLaunchTemplateCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createLaunchTemplateCmd)
}
