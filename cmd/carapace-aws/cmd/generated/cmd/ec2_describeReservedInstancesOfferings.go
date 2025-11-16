package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeReservedInstancesOfferingsCmd = &cobra.Command{
	Use:   "describe-reserved-instances-offerings",
	Short: "Describes Reserved Instance offerings that are available for purchase.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeReservedInstancesOfferingsCmd).Standalone()

	ec2_describeReservedInstancesOfferingsCmd.Flags().String("availability-zone", "", "The Availability Zone in which the Reserved Instance can be used.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("availability-zone-id", "", "The ID of the Availability Zone.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().Bool("include-marketplace", false, "Include Reserved Instance Marketplace offerings in the response.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("instance-tenancy", "", "The tenancy of the instances covered by the reservation.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("instance-type", "", "The instance type that the reservation will cover (for example, `m1.small`).")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("max-duration", "", "The maximum duration (in seconds) to filter when searching for offerings.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("max-instance-count", "", "The maximum number of instances to filter when searching for offerings.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("min-duration", "", "The minimum duration (in seconds) to filter when searching for offerings.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().Bool("no-include-marketplace", false, "Include Reserved Instance Marketplace offerings in the response.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("offering-class", "", "The offering class of the Reserved Instance.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("offering-type", "", "The Reserved Instance offering type.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("product-description", "", "The Reserved Instance product platform description.")
	ec2_describeReservedInstancesOfferingsCmd.Flags().String("reserved-instances-offering-ids", "", "One or more Reserved Instances offering IDs.")
	ec2_describeReservedInstancesOfferingsCmd.Flag("no-dry-run").Hidden = true
	ec2_describeReservedInstancesOfferingsCmd.Flag("no-include-marketplace").Hidden = true
	ec2Cmd.AddCommand(ec2_describeReservedInstancesOfferingsCmd)
}
