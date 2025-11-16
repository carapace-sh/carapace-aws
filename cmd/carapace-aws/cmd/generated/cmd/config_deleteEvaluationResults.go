package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteEvaluationResultsCmd = &cobra.Command{
	Use:   "delete-evaluation-results",
	Short: "Deletes the evaluation results for the specified Config rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteEvaluationResultsCmd).Standalone()

	config_deleteEvaluationResultsCmd.Flags().String("config-rule-name", "", "The name of the Config rule for which you want to delete the evaluation results.")
	config_deleteEvaluationResultsCmd.MarkFlagRequired("config-rule-name")
	configCmd.AddCommand(config_deleteEvaluationResultsCmd)
}
