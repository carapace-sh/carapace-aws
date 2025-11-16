package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeParametersCmd = &cobra.Command{
	Use:   "describe-parameters",
	Short: "Returns the detailed parameter list for a particular parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_describeParametersCmd).Standalone()

		memorydb_describeParametersCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		memorydb_describeParametersCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
		memorydb_describeParametersCmd.Flags().String("parameter-group-name", "", "he name of a specific parameter group to return details for.")
		memorydb_describeParametersCmd.MarkFlagRequired("parameter-group-name")
	})
	memorydbCmd.AddCommand(memorydb_describeParametersCmd)
}
