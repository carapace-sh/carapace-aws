package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_batchEvaluateFeatureCmd = &cobra.Command{
	Use:   "batch-evaluate-feature",
	Short: "This operation assigns feature variation to user sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_batchEvaluateFeatureCmd).Standalone()

	evidently_batchEvaluateFeatureCmd.Flags().String("project", "", "The name or ARN of the project that contains the feature being evaluated.")
	evidently_batchEvaluateFeatureCmd.Flags().String("requests", "", "An array of structures, where each structure assigns a feature variation to one user session.")
	evidently_batchEvaluateFeatureCmd.MarkFlagRequired("project")
	evidently_batchEvaluateFeatureCmd.MarkFlagRequired("requests")
	evidentlyCmd.AddCommand(evidently_batchEvaluateFeatureCmd)
}
