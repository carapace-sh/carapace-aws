package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listTestExecutionResultItemsCmd = &cobra.Command{
	Use:   "list-test-execution-result-items",
	Short: "Gets a list of test execution result items.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listTestExecutionResultItemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listTestExecutionResultItemsCmd).Standalone()

		lexv2Models_listTestExecutionResultItemsCmd.Flags().String("max-results", "", "The maximum number of test execution result items to return in each page.")
		lexv2Models_listTestExecutionResultItemsCmd.Flags().String("next-token", "", "If the response from the `ListTestExecutionResultItems` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
		lexv2Models_listTestExecutionResultItemsCmd.Flags().String("result-filter-by", "", "The filter for the list of results from the test set execution.")
		lexv2Models_listTestExecutionResultItemsCmd.Flags().String("test-execution-id", "", "The unique identifier of the test execution to list the result items.")
		lexv2Models_listTestExecutionResultItemsCmd.MarkFlagRequired("result-filter-by")
		lexv2Models_listTestExecutionResultItemsCmd.MarkFlagRequired("test-execution-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listTestExecutionResultItemsCmd)
}
