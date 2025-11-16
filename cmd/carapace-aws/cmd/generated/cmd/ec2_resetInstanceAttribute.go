package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_resetInstanceAttributeCmd = &cobra.Command{
	Use:   "reset-instance-attribute",
	Short: "Resets an attribute of an instance to its default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_resetInstanceAttributeCmd).Standalone()

	ec2_resetInstanceAttributeCmd.Flags().String("attribute", "", "The attribute to reset.")
	ec2_resetInstanceAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_resetInstanceAttributeCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_resetInstanceAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_resetInstanceAttributeCmd.MarkFlagRequired("attribute")
	ec2_resetInstanceAttributeCmd.MarkFlagRequired("instance-id")
	ec2_resetInstanceAttributeCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_resetInstanceAttributeCmd)
}
