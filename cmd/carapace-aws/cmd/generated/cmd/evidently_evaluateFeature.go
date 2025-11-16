package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_evaluateFeatureCmd = &cobra.Command{
	Use:   "evaluate-feature",
	Short: "This operation assigns a feature variation to one given user session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_evaluateFeatureCmd).Standalone()

	evidently_evaluateFeatureCmd.Flags().String("entity-id", "", "An internal ID that represents a unique user of the application.")
	evidently_evaluateFeatureCmd.Flags().String("evaluation-context", "", "A JSON object of attributes that you can optionally pass in as part of the evaluation event sent to Evidently from the user session.")
	evidently_evaluateFeatureCmd.Flags().String("feature", "", "The name of the feature being evaluated.")
	evidently_evaluateFeatureCmd.Flags().String("project", "", "The name or ARN of the project that contains this feature.")
	evidently_evaluateFeatureCmd.MarkFlagRequired("entity-id")
	evidently_evaluateFeatureCmd.MarkFlagRequired("feature")
	evidently_evaluateFeatureCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_evaluateFeatureCmd)
}
