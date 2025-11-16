package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeParameterGroupsCmd = &cobra.Command{
	Use:   "describe-parameter-groups",
	Short: "Returns a list of parameter group descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeParameterGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_describeParameterGroupsCmd).Standalone()

		memorydb_describeParameterGroupsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		memorydb_describeParameterGroupsCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
		memorydb_describeParameterGroupsCmd.Flags().String("parameter-group-name", "", "The name of a specific parameter group to return details for.")
	})
	memorydbCmd.AddCommand(memorydb_describeParameterGroupsCmd)
}
