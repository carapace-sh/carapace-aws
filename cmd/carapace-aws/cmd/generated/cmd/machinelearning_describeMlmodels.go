package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_describeMlmodelsCmd = &cobra.Command{
	Use:   "describe-mlmodels",
	Short: "Returns a list of `MLModel` that match the search criteria in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_describeMlmodelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_describeMlmodelsCmd).Standalone()

		machinelearning_describeMlmodelsCmd.Flags().String("eq", "", "The equal to operator.")
		machinelearning_describeMlmodelsCmd.Flags().String("filter-variable", "", "Use one of the following variables to filter a list of `MLModel`:")
		machinelearning_describeMlmodelsCmd.Flags().String("ge", "", "The greater than or equal to operator.")
		machinelearning_describeMlmodelsCmd.Flags().String("gt", "", "The greater than operator.")
		machinelearning_describeMlmodelsCmd.Flags().String("le", "", "The less than or equal to operator.")
		machinelearning_describeMlmodelsCmd.Flags().String("limit", "", "The number of pages of information to include in the result.")
		machinelearning_describeMlmodelsCmd.Flags().String("lt", "", "The less than operator.")
		machinelearning_describeMlmodelsCmd.Flags().String("ne", "", "The not equal to operator.")
		machinelearning_describeMlmodelsCmd.Flags().String("next-token", "", "The ID of the page in the paginated results.")
		machinelearning_describeMlmodelsCmd.Flags().String("prefix", "", "A string that is found at the beginning of a variable, such as `Name` or `Id`.")
		machinelearning_describeMlmodelsCmd.Flags().String("sort-order", "", "A two-value parameter that determines the sequence of the resulting list of `MLModel`.")
	})
	machinelearningCmd.AddCommand(machinelearning_describeMlmodelsCmd)
}
