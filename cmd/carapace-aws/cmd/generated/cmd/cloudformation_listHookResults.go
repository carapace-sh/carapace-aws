package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listHookResultsCmd = &cobra.Command{
	Use:   "list-hook-results",
	Short: "Returns summaries of invoked Hooks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listHookResultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listHookResultsCmd).Standalone()

		cloudformation_listHookResultsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listHookResultsCmd.Flags().String("status", "", "Filters results by the status of Hook invocations.")
		cloudformation_listHookResultsCmd.Flags().String("target-id", "", "Filters results by the unique identifier of the target the Hook was invoked against.")
		cloudformation_listHookResultsCmd.Flags().String("target-type", "", "Filters results by target type.")
		cloudformation_listHookResultsCmd.Flags().String("type-arn", "", "Filters results by the ARN of the Hook.")
	})
	cloudformationCmd.AddCommand(cloudformation_listHookResultsCmd)
}
