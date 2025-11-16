package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_updateFeatureCmd = &cobra.Command{
	Use:   "update-feature",
	Short: "Updates an existing feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_updateFeatureCmd).Standalone()

	evidently_updateFeatureCmd.Flags().String("add-or-update-variations", "", "To update variation configurations for this feature, or add new ones, specify this structure.")
	evidently_updateFeatureCmd.Flags().String("default-variation", "", "The name of the variation to use as the default variation.")
	evidently_updateFeatureCmd.Flags().String("description", "", "An optional description of the feature.")
	evidently_updateFeatureCmd.Flags().String("entity-overrides", "", "Specified users that should always be served a specific variation of a feature.")
	evidently_updateFeatureCmd.Flags().String("evaluation-strategy", "", "Specify `ALL_RULES` to activate the traffic allocation specified by any ongoing launches or experiments.")
	evidently_updateFeatureCmd.Flags().String("feature", "", "The name of the feature to be updated.")
	evidently_updateFeatureCmd.Flags().String("project", "", "The name or ARN of the project that contains the feature to be updated.")
	evidently_updateFeatureCmd.Flags().String("remove-variations", "", "Removes a variation from the feature.")
	evidently_updateFeatureCmd.MarkFlagRequired("feature")
	evidently_updateFeatureCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_updateFeatureCmd)
}
