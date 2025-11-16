package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeExplainabilityCmd = &cobra.Command{
	Use:   "describe-explainability",
	Short: "Describes an Explainability resource created using the [CreateExplainability]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeExplainabilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describeExplainabilityCmd).Standalone()

		forecast_describeExplainabilityCmd.Flags().String("explainability-arn", "", "The Amazon Resource Name (ARN) of the Explaianability to describe.")
		forecast_describeExplainabilityCmd.MarkFlagRequired("explainability-arn")
	})
	forecastCmd.AddCommand(forecast_describeExplainabilityCmd)
}
