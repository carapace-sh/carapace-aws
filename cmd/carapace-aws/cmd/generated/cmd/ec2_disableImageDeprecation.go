package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableImageDeprecationCmd = &cobra.Command{
	Use:   "disable-image-deprecation",
	Short: "Cancels the deprecation of the specified AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableImageDeprecationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disableImageDeprecationCmd).Standalone()

		ec2_disableImageDeprecationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableImageDeprecationCmd.Flags().String("image-id", "", "The ID of the AMI.")
		ec2_disableImageDeprecationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableImageDeprecationCmd.MarkFlagRequired("image-id")
		ec2_disableImageDeprecationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disableImageDeprecationCmd)
}
