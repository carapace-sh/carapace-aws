package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeReservedInstancesListingsCmd = &cobra.Command{
	Use:   "describe-reserved-instances-listings",
	Short: "Describes your account's Reserved Instance listings in the Reserved Instance Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeReservedInstancesListingsCmd).Standalone()

	ec2_describeReservedInstancesListingsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeReservedInstancesListingsCmd.Flags().String("reserved-instances-id", "", "One or more Reserved Instance IDs.")
	ec2_describeReservedInstancesListingsCmd.Flags().String("reserved-instances-listing-id", "", "One or more Reserved Instance listing IDs.")
	ec2Cmd.AddCommand(ec2_describeReservedInstancesListingsCmd)
}
