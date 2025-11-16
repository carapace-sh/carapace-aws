package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listResourceEvaluationsCmd = &cobra.Command{
	Use:   "list-resource-evaluations",
	Short: "Returns a list of proactive resource evaluations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listResourceEvaluationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_listResourceEvaluationsCmd).Standalone()

		config_listResourceEvaluationsCmd.Flags().String("filters", "", "Returns a `ResourceEvaluationFilters` object.")
		config_listResourceEvaluationsCmd.Flags().String("limit", "", "The maximum number of evaluations returned on each page.")
		config_listResourceEvaluationsCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	})
	configCmd.AddCommand(config_listResourceEvaluationsCmd)
}
