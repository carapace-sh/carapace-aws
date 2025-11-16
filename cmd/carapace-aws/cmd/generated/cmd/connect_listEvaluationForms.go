package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listEvaluationFormsCmd = &cobra.Command{
	Use:   "list-evaluation-forms",
	Short: "Lists evaluation forms in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listEvaluationFormsCmd).Standalone()

	connect_listEvaluationFormsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listEvaluationFormsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listEvaluationFormsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listEvaluationFormsCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listEvaluationFormsCmd)
}
