package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_evaluateExpressionCmd = &cobra.Command{
	Use:   "evaluate-expression",
	Short: "Task runners call `EvaluateExpression` to evaluate a string in the context of the specified object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_evaluateExpressionCmd).Standalone()

	datapipeline_evaluateExpressionCmd.Flags().String("expression", "", "The expression to evaluate.")
	datapipeline_evaluateExpressionCmd.Flags().String("object-id", "", "The ID of the object.")
	datapipeline_evaluateExpressionCmd.Flags().String("pipeline-id", "", "The ID of the pipeline.")
	datapipeline_evaluateExpressionCmd.MarkFlagRequired("expression")
	datapipeline_evaluateExpressionCmd.MarkFlagRequired("object-id")
	datapipeline_evaluateExpressionCmd.MarkFlagRequired("pipeline-id")
	datapipelineCmd.AddCommand(datapipeline_evaluateExpressionCmd)
}
