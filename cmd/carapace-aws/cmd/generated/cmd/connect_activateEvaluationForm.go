package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_activateEvaluationFormCmd = &cobra.Command{
	Use:   "activate-evaluation-form",
	Short: "Activates an evaluation form in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_activateEvaluationFormCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_activateEvaluationFormCmd).Standalone()

		connect_activateEvaluationFormCmd.Flags().String("evaluation-form-id", "", "The unique identifier for the evaluation form.")
		connect_activateEvaluationFormCmd.Flags().String("evaluation-form-version", "", "The version of the evaluation form to activate.")
		connect_activateEvaluationFormCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_activateEvaluationFormCmd.MarkFlagRequired("evaluation-form-id")
		connect_activateEvaluationFormCmd.MarkFlagRequired("evaluation-form-version")
		connect_activateEvaluationFormCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_activateEvaluationFormCmd)
}
