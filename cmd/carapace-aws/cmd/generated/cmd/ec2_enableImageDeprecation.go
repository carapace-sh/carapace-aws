package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableImageDeprecationCmd = &cobra.Command{
	Use:   "enable-image-deprecation",
	Short: "Enables deprecation of the specified AMI at the specified date and time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableImageDeprecationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableImageDeprecationCmd).Standalone()

		ec2_enableImageDeprecationCmd.Flags().String("deprecate-at", "", "The date and time to deprecate the AMI, in UTC, in the following format: *YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z.")
		ec2_enableImageDeprecationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableImageDeprecationCmd.Flags().String("image-id", "", "The ID of the AMI.")
		ec2_enableImageDeprecationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableImageDeprecationCmd.MarkFlagRequired("deprecate-at")
		ec2_enableImageDeprecationCmd.MarkFlagRequired("image-id")
		ec2_enableImageDeprecationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_enableImageDeprecationCmd)
}
