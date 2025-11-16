package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_purchaseReservedInstancesOfferingCmd = &cobra.Command{
	Use:   "purchase-reserved-instances-offering",
	Short: "Purchases a Reserved Instance for use with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_purchaseReservedInstancesOfferingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_purchaseReservedInstancesOfferingCmd).Standalone()

		ec2_purchaseReservedInstancesOfferingCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_purchaseReservedInstancesOfferingCmd.Flags().String("instance-count", "", "The number of Reserved Instances to purchase.")
		ec2_purchaseReservedInstancesOfferingCmd.Flags().String("limit-price", "", "Specified for Reserved Instance Marketplace offerings to limit the total order and ensure that the Reserved Instances are not purchased at unexpected prices.")
		ec2_purchaseReservedInstancesOfferingCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_purchaseReservedInstancesOfferingCmd.Flags().String("purchase-time", "", "The time at which to purchase the Reserved Instance, in UTC format (for example, *YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z).")
		ec2_purchaseReservedInstancesOfferingCmd.Flags().String("reserved-instances-offering-id", "", "The ID of the Reserved Instance offering to purchase.")
		ec2_purchaseReservedInstancesOfferingCmd.MarkFlagRequired("instance-count")
		ec2_purchaseReservedInstancesOfferingCmd.Flag("no-dry-run").Hidden = true
		ec2_purchaseReservedInstancesOfferingCmd.MarkFlagRequired("reserved-instances-offering-id")
	})
	ec2Cmd.AddCommand(ec2_purchaseReservedInstancesOfferingCmd)
}
