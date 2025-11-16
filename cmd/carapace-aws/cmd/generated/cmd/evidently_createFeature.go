package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_createFeatureCmd = &cobra.Command{
	Use:   "create-feature",
	Short: "Creates an Evidently *feature* that you want to launch or test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_createFeatureCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_createFeatureCmd).Standalone()

		evidently_createFeatureCmd.Flags().String("default-variation", "", "The name of the variation to use as the default variation.")
		evidently_createFeatureCmd.Flags().String("description", "", "An optional description of the feature.")
		evidently_createFeatureCmd.Flags().String("entity-overrides", "", "Specify users that should always be served a specific variation of a feature.")
		evidently_createFeatureCmd.Flags().String("evaluation-strategy", "", "Specify `ALL_RULES` to activate the traffic allocation specified by any ongoing launches or experiments.")
		evidently_createFeatureCmd.Flags().String("name", "", "The name for the new feature.")
		evidently_createFeatureCmd.Flags().String("project", "", "The name or ARN of the project that is to contain the new feature.")
		evidently_createFeatureCmd.Flags().String("tags", "", "Assigns one or more tags (key-value pairs) to the feature.")
		evidently_createFeatureCmd.Flags().String("variations", "", "An array of structures that contain the configuration of the feature's different variations.")
		evidently_createFeatureCmd.MarkFlagRequired("name")
		evidently_createFeatureCmd.MarkFlagRequired("project")
		evidently_createFeatureCmd.MarkFlagRequired("variations")
	})
	evidentlyCmd.AddCommand(evidently_createFeatureCmd)
}
