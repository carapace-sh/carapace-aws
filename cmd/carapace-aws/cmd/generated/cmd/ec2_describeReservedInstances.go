package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeReservedInstancesCmd = &cobra.Command{
	Use:   "describe-reserved-instances",
	Short: "Describes one or more of the Reserved Instances that you purchased.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeReservedInstancesCmd).Standalone()

	ec2_describeReservedInstancesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeReservedInstancesCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeReservedInstancesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeReservedInstancesCmd.Flags().String("offering-class", "", "Describes whether the Reserved Instance is Standard or Convertible.")
	ec2_describeReservedInstancesCmd.Flags().String("offering-type", "", "The Reserved Instance offering type.")
	ec2_describeReservedInstancesCmd.Flags().String("reserved-instances-ids", "", "One or more Reserved Instance IDs.")
	ec2_describeReservedInstancesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeReservedInstancesCmd)
}
