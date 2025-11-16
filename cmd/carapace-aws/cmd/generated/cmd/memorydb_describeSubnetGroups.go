package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeSubnetGroupsCmd = &cobra.Command{
	Use:   "describe-subnet-groups",
	Short: "Returns a list of subnet group descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeSubnetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_describeSubnetGroupsCmd).Standalone()

		memorydb_describeSubnetGroupsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		memorydb_describeSubnetGroupsCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
		memorydb_describeSubnetGroupsCmd.Flags().String("subnet-group-name", "", "The name of the subnet group to return details for.")
	})
	memorydbCmd.AddCommand(memorydb_describeSubnetGroupsCmd)
}
