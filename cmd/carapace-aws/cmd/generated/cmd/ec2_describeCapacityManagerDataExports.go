package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeCapacityManagerDataExportsCmd = &cobra.Command{
	Use:   "describe-capacity-manager-data-exports",
	Short: "Describes one or more Capacity Manager data export configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeCapacityManagerDataExportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeCapacityManagerDataExportsCmd).Standalone()

		ec2_describeCapacityManagerDataExportsCmd.Flags().String("capacity-manager-data-export-ids", "", "The IDs of the data export configurations to describe.")
		ec2_describeCapacityManagerDataExportsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityManagerDataExportsCmd.Flags().String("filters", "", "One or more filters to narrow the results.")
		ec2_describeCapacityManagerDataExportsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		ec2_describeCapacityManagerDataExportsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeCapacityManagerDataExportsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeCapacityManagerDataExportsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeCapacityManagerDataExportsCmd)
}
