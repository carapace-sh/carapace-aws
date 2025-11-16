package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeReservedInstancesModificationsCmd = &cobra.Command{
	Use:   "describe-reserved-instances-modifications",
	Short: "Describes the modifications made to your Reserved Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeReservedInstancesModificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeReservedInstancesModificationsCmd).Standalone()

		ec2_describeReservedInstancesModificationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeReservedInstancesModificationsCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
		ec2_describeReservedInstancesModificationsCmd.Flags().String("reserved-instances-modification-ids", "", "IDs for the submitted modification request.")
	})
	ec2Cmd.AddCommand(ec2_describeReservedInstancesModificationsCmd)
}
