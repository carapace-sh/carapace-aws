package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_resetAddressAttributeCmd = &cobra.Command{
	Use:   "reset-address-attribute",
	Short: "Resets the attribute of the specified IP address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_resetAddressAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_resetAddressAttributeCmd).Standalone()

		ec2_resetAddressAttributeCmd.Flags().String("allocation-id", "", "\\[EC2-VPC] The allocation ID.")
		ec2_resetAddressAttributeCmd.Flags().String("attribute", "", "The attribute of the IP address.")
		ec2_resetAddressAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetAddressAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetAddressAttributeCmd.MarkFlagRequired("allocation-id")
		ec2_resetAddressAttributeCmd.MarkFlagRequired("attribute")
		ec2_resetAddressAttributeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_resetAddressAttributeCmd)
}
