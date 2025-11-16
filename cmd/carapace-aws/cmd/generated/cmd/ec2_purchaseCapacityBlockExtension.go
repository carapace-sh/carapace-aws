package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_purchaseCapacityBlockExtensionCmd = &cobra.Command{
	Use:   "purchase-capacity-block-extension",
	Short: "Purchase the Capacity Block extension for use with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_purchaseCapacityBlockExtensionCmd).Standalone()

	ec2_purchaseCapacityBlockExtensionCmd.Flags().String("capacity-block-extension-offering-id", "", "The ID of the Capacity Block extension offering to purchase.")
	ec2_purchaseCapacityBlockExtensionCmd.Flags().String("capacity-reservation-id", "", "The ID of the Capacity reservation to be extended.")
	ec2_purchaseCapacityBlockExtensionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_purchaseCapacityBlockExtensionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_purchaseCapacityBlockExtensionCmd.MarkFlagRequired("capacity-block-extension-offering-id")
	ec2_purchaseCapacityBlockExtensionCmd.MarkFlagRequired("capacity-reservation-id")
	ec2_purchaseCapacityBlockExtensionCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_purchaseCapacityBlockExtensionCmd)
}
