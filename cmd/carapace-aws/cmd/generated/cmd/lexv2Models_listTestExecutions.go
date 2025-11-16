package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listTestExecutionsCmd = &cobra.Command{
	Use:   "list-test-executions",
	Short: "The list of test set executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listTestExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_listTestExecutionsCmd).Standalone()

		lexv2Models_listTestExecutionsCmd.Flags().String("max-results", "", "The maximum number of test executions to return in each page.")
		lexv2Models_listTestExecutionsCmd.Flags().String("next-token", "", "If the response from the ListTestExecutions operation contains more results than specified in the maxResults parameter, a token is returned in the response.")
		lexv2Models_listTestExecutionsCmd.Flags().String("sort-by", "", "The sort order of the test set executions.")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_listTestExecutionsCmd)
}
