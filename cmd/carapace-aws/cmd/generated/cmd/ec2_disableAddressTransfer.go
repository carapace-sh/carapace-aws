package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableAddressTransferCmd = &cobra.Command{
	Use:   "disable-address-transfer",
	Short: "Disables Elastic IP address transfer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableAddressTransferCmd).Standalone()

	ec2_disableAddressTransferCmd.Flags().String("allocation-id", "", "The allocation ID of an Elastic IP address.")
	ec2_disableAddressTransferCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableAddressTransferCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableAddressTransferCmd.MarkFlagRequired("allocation-id")
	ec2_disableAddressTransferCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disableAddressTransferCmd)
}
