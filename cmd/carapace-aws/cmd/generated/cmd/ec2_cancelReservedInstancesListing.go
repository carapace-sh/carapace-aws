package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_cancelReservedInstancesListingCmd = &cobra.Command{
	Use:   "cancel-reserved-instances-listing",
	Short: "Cancels the specified Reserved Instance listing in the Reserved Instance Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_cancelReservedInstancesListingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_cancelReservedInstancesListingCmd).Standalone()

		ec2_cancelReservedInstancesListingCmd.Flags().String("reserved-instances-listing-id", "", "The ID of the Reserved Instance listing.")
		ec2_cancelReservedInstancesListingCmd.MarkFlagRequired("reserved-instances-listing-id")
	})
	ec2Cmd.AddCommand(ec2_cancelReservedInstancesListingCmd)
}
