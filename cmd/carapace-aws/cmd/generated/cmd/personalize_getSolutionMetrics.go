package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_getSolutionMetricsCmd = &cobra.Command{
	Use:   "get-solution-metrics",
	Short: "Gets the metrics for the specified solution version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_getSolutionMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_getSolutionMetricsCmd).Standalone()

		personalize_getSolutionMetricsCmd.Flags().String("solution-version-arn", "", "The Amazon Resource Name (ARN) of the solution version for which to get metrics.")
		personalize_getSolutionMetricsCmd.MarkFlagRequired("solution-version-arn")
	})
	personalizeCmd.AddCommand(personalize_getSolutionMetricsCmd)
}
