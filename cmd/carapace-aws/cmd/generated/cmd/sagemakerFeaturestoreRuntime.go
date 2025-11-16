package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerFeaturestoreRuntimeCmd = &cobra.Command{
	Use:   "sagemaker-featurestore-runtime",
	Short: "Contains all data plane API operations and data types for the Amazon SageMaker Feature Store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerFeaturestoreRuntimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerFeaturestoreRuntimeCmd).Standalone()

	})
	rootCmd.AddCommand(sagemakerFeaturestoreRuntimeCmd)
}
