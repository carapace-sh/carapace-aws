package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_describeDataSourcesCmd = &cobra.Command{
	Use:   "describe-data-sources",
	Short: "Returns a list of `DataSource` that match the search criteria in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_describeDataSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_describeDataSourcesCmd).Standalone()

		machinelearning_describeDataSourcesCmd.Flags().String("eq", "", "The equal to operator.")
		machinelearning_describeDataSourcesCmd.Flags().String("filter-variable", "", "Use one of the following variables to filter a list of `DataSource`:")
		machinelearning_describeDataSourcesCmd.Flags().String("ge", "", "The greater than or equal to operator.")
		machinelearning_describeDataSourcesCmd.Flags().String("gt", "", "The greater than operator.")
		machinelearning_describeDataSourcesCmd.Flags().String("le", "", "The less than or equal to operator.")
		machinelearning_describeDataSourcesCmd.Flags().String("limit", "", "The maximum number of `DataSource` to include in the result.")
		machinelearning_describeDataSourcesCmd.Flags().String("lt", "", "The less than operator.")
		machinelearning_describeDataSourcesCmd.Flags().String("ne", "", "The not equal to operator.")
		machinelearning_describeDataSourcesCmd.Flags().String("next-token", "", "The ID of the page in the paginated results.")
		machinelearning_describeDataSourcesCmd.Flags().String("prefix", "", "A string that is found at the beginning of a variable, such as `Name` or `Id`.")
		machinelearning_describeDataSourcesCmd.Flags().String("sort-order", "", "A two-value parameter that determines the sequence of the resulting list of `DataSource`.")
	})
	machinelearningCmd.AddCommand(machinelearning_describeDataSourcesCmd)
}
