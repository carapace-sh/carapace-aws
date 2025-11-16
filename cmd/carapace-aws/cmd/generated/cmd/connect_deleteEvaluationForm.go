package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteEvaluationFormCmd = &cobra.Command{
	Use:   "delete-evaluation-form",
	Short: "Deletes an evaluation form in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteEvaluationFormCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteEvaluationFormCmd).Standalone()

		connect_deleteEvaluationFormCmd.Flags().String("evaluation-form-id", "", "The unique identifier for the evaluation form.")
		connect_deleteEvaluationFormCmd.Flags().String("evaluation-form-version", "", "The unique identifier for the evaluation form.")
		connect_deleteEvaluationFormCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteEvaluationFormCmd.MarkFlagRequired("evaluation-form-id")
		connect_deleteEvaluationFormCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_deleteEvaluationFormCmd)
}
