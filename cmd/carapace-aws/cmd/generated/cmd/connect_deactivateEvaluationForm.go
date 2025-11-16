package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deactivateEvaluationFormCmd = &cobra.Command{
	Use:   "deactivate-evaluation-form",
	Short: "Deactivates an evaluation form in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deactivateEvaluationFormCmd).Standalone()

	connect_deactivateEvaluationFormCmd.Flags().String("evaluation-form-id", "", "The unique identifier for the evaluation form.")
	connect_deactivateEvaluationFormCmd.Flags().String("evaluation-form-version", "", "A version of the evaluation form.")
	connect_deactivateEvaluationFormCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deactivateEvaluationFormCmd.MarkFlagRequired("evaluation-form-id")
	connect_deactivateEvaluationFormCmd.MarkFlagRequired("evaluation-form-version")
	connect_deactivateEvaluationFormCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_deactivateEvaluationFormCmd)
}
