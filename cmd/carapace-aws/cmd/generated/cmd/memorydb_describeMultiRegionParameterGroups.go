package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeMultiRegionParameterGroupsCmd = &cobra.Command{
	Use:   "describe-multi-region-parameter-groups",
	Short: "Returns a list of multi-region parameter groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeMultiRegionParameterGroupsCmd).Standalone()

	memorydb_describeMultiRegionParameterGroupsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
	memorydb_describeMultiRegionParameterGroupsCmd.Flags().String("multi-region-parameter-group-name", "", "The request for information on a specific multi-region parameter group.")
	memorydb_describeMultiRegionParameterGroupsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	memorydbCmd.AddCommand(memorydb_describeMultiRegionParameterGroupsCmd)
}
