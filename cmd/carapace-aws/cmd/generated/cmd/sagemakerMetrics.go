package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerMetricsCmd = &cobra.Command{
	Use:   "sagemaker-metrics",
	Short: "Contains all data plane API operations and data types for Amazon SageMaker Metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerMetricsCmd).Standalone()

	rootCmd.AddCommand(sagemakerMetricsCmd)
}
