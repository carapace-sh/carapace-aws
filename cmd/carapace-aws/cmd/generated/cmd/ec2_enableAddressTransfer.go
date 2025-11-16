package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableAddressTransferCmd = &cobra.Command{
	Use:   "enable-address-transfer",
	Short: "Enables Elastic IP address transfer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableAddressTransferCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableAddressTransferCmd).Standalone()

		ec2_enableAddressTransferCmd.Flags().String("allocation-id", "", "The allocation ID of an Elastic IP address.")
		ec2_enableAddressTransferCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableAddressTransferCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableAddressTransferCmd.Flags().String("transfer-account-id", "", "The ID of the account that you want to transfer the Elastic IP address to.")
		ec2_enableAddressTransferCmd.MarkFlagRequired("allocation-id")
		ec2_enableAddressTransferCmd.Flag("no-dry-run").Hidden = true
		ec2_enableAddressTransferCmd.MarkFlagRequired("transfer-account-id")
	})
	ec2Cmd.AddCommand(ec2_enableAddressTransferCmd)
}
