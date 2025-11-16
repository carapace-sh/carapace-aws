package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeMultiRegionParametersCmd = &cobra.Command{
	Use:   "describe-multi-region-parameters",
	Short: "Returns the detailed parameter list for a particular multi-region parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeMultiRegionParametersCmd).Standalone()

	memorydb_describeMultiRegionParametersCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
	memorydb_describeMultiRegionParametersCmd.Flags().String("multi-region-parameter-group-name", "", "The name of the multi-region parameter group to return details for.")
	memorydb_describeMultiRegionParametersCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	memorydb_describeMultiRegionParametersCmd.Flags().String("source", "", "The parameter types to return.")
	memorydb_describeMultiRegionParametersCmd.MarkFlagRequired("multi-region-parameter-group-name")
	memorydbCmd.AddCommand(memorydb_describeMultiRegionParametersCmd)
}
