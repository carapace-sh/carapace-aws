package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateEvaluationFormCmd = &cobra.Command{
	Use:   "update-evaluation-form",
	Short: "Updates details about a specific evaluation form version in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateEvaluationFormCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateEvaluationFormCmd).Standalone()

		connect_updateEvaluationFormCmd.Flags().String("auto-evaluation-configuration", "", "Whether automated evaluations are enabled.")
		connect_updateEvaluationFormCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_updateEvaluationFormCmd.Flags().String("create-new-version", "", "A flag indicating whether the operation must create a new version.")
		connect_updateEvaluationFormCmd.Flags().String("description", "", "The description of the evaluation form.")
		connect_updateEvaluationFormCmd.Flags().String("evaluation-form-id", "", "The unique identifier for the evaluation form.")
		connect_updateEvaluationFormCmd.Flags().String("evaluation-form-version", "", "A version of the evaluation form to update.")
		connect_updateEvaluationFormCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateEvaluationFormCmd.Flags().String("items", "", "Items that are part of the evaluation form.")
		connect_updateEvaluationFormCmd.Flags().String("scoring-strategy", "", "A scoring strategy of the evaluation form.")
		connect_updateEvaluationFormCmd.Flags().String("title", "", "A title of the evaluation form.")
		connect_updateEvaluationFormCmd.MarkFlagRequired("evaluation-form-id")
		connect_updateEvaluationFormCmd.MarkFlagRequired("evaluation-form-version")
		connect_updateEvaluationFormCmd.MarkFlagRequired("instance-id")
		connect_updateEvaluationFormCmd.MarkFlagRequired("items")
		connect_updateEvaluationFormCmd.MarkFlagRequired("title")
	})
	connectCmd.AddCommand(connect_updateEvaluationFormCmd)
}
