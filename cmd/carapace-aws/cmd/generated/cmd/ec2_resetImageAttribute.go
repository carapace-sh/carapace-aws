package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_resetImageAttributeCmd = &cobra.Command{
	Use:   "reset-image-attribute",
	Short: "Resets an attribute of an AMI to its default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_resetImageAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_resetImageAttributeCmd).Standalone()

		ec2_resetImageAttributeCmd.Flags().String("attribute", "", "The attribute to reset (currently you can only reset the launch permission attribute).")
		ec2_resetImageAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetImageAttributeCmd.Flags().String("image-id", "", "The ID of the AMI.")
		ec2_resetImageAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetImageAttributeCmd.MarkFlagRequired("attribute")
		ec2_resetImageAttributeCmd.MarkFlagRequired("image-id")
		ec2_resetImageAttributeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_resetImageAttributeCmd)
}
