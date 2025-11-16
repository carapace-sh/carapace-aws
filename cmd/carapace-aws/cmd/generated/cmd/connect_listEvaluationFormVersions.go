package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listEvaluationFormVersionsCmd = &cobra.Command{
	Use:   "list-evaluation-form-versions",
	Short: "Lists versions of an evaluation form in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listEvaluationFormVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listEvaluationFormVersionsCmd).Standalone()

		connect_listEvaluationFormVersionsCmd.Flags().String("evaluation-form-id", "", "The unique identifier for the evaluation form.")
		connect_listEvaluationFormVersionsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listEvaluationFormVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listEvaluationFormVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listEvaluationFormVersionsCmd.MarkFlagRequired("evaluation-form-id")
		connect_listEvaluationFormVersionsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listEvaluationFormVersionsCmd)
}
