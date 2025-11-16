package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyAddressAttributeCmd = &cobra.Command{
	Use:   "modify-address-attribute",
	Short: "Modifies an attribute of the specified Elastic IP address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyAddressAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyAddressAttributeCmd).Standalone()

		ec2_modifyAddressAttributeCmd.Flags().String("allocation-id", "", "\\[EC2-VPC] The allocation ID.")
		ec2_modifyAddressAttributeCmd.Flags().String("domain-name", "", "The domain name to modify for the IP address.")
		ec2_modifyAddressAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyAddressAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyAddressAttributeCmd.MarkFlagRequired("allocation-id")
		ec2_modifyAddressAttributeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyAddressAttributeCmd)
}
