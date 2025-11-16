package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_predictQaresultsCmd = &cobra.Command{
	Use:   "predict-qaresults",
	Short: "Predicts existing visuals or generates new visuals to answer a given query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_predictQaresultsCmd).Standalone()

	quicksight_predictQaresultsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that the user wants to execute Predict QA results in.")
	quicksight_predictQaresultsCmd.Flags().String("include-generated-answer", "", "Indicates whether generated answers are included or excluded.")
	quicksight_predictQaresultsCmd.Flags().String("include-quick-sight-qindex", "", "Indicates whether Q indicies are included or excluded.")
	quicksight_predictQaresultsCmd.Flags().String("max-topics-to-consider", "", "The number of maximum topics to be considered to predict QA results.")
	quicksight_predictQaresultsCmd.Flags().String("query-text", "", "The query text to be used to predict QA results.")
	quicksight_predictQaresultsCmd.MarkFlagRequired("aws-account-id")
	quicksight_predictQaresultsCmd.MarkFlagRequired("query-text")
	quicksightCmd.AddCommand(quicksight_predictQaresultsCmd)
}
