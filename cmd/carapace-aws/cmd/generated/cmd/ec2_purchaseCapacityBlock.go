package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_purchaseCapacityBlockCmd = &cobra.Command{
	Use:   "purchase-capacity-block",
	Short: "Purchase the Capacity Block for use with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_purchaseCapacityBlockCmd).Standalone()

	ec2_purchaseCapacityBlockCmd.Flags().String("capacity-block-offering-id", "", "The ID of the Capacity Block offering.")
	ec2_purchaseCapacityBlockCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_purchaseCapacityBlockCmd.Flags().String("instance-platform", "", "The type of operating system for which to reserve capacity.")
	ec2_purchaseCapacityBlockCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_purchaseCapacityBlockCmd.Flags().String("tag-specifications", "", "The tags to apply to the Capacity Block during launch.")
	ec2_purchaseCapacityBlockCmd.MarkFlagRequired("capacity-block-offering-id")
	ec2_purchaseCapacityBlockCmd.MarkFlagRequired("instance-platform")
	ec2_purchaseCapacityBlockCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_purchaseCapacityBlockCmd)
}
