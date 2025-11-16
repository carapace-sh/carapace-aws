package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createReservedInstancesListingCmd = &cobra.Command{
	Use:   "create-reserved-instances-listing",
	Short: "Creates a listing for Amazon EC2 Standard Reserved Instances to be sold in the Reserved Instance Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createReservedInstancesListingCmd).Standalone()

	ec2_createReservedInstancesListingCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of your listings.")
	ec2_createReservedInstancesListingCmd.Flags().String("instance-count", "", "The number of instances that are a part of a Reserved Instance account to be listed in the Reserved Instance Marketplace.")
	ec2_createReservedInstancesListingCmd.Flags().String("price-schedules", "", "A list specifying the price of the Standard Reserved Instance for each month remaining in the Reserved Instance term.")
	ec2_createReservedInstancesListingCmd.Flags().String("reserved-instances-id", "", "The ID of the active Standard Reserved Instance.")
	ec2_createReservedInstancesListingCmd.MarkFlagRequired("client-token")
	ec2_createReservedInstancesListingCmd.MarkFlagRequired("instance-count")
	ec2_createReservedInstancesListingCmd.MarkFlagRequired("price-schedules")
	ec2_createReservedInstancesListingCmd.MarkFlagRequired("reserved-instances-id")
	ec2Cmd.AddCommand(ec2_createReservedInstancesListingCmd)
}
