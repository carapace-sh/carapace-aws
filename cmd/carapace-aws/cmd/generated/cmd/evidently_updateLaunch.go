package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_updateLaunchCmd = &cobra.Command{
	Use:   "update-launch",
	Short: "Updates a launch of a given feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_updateLaunchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_updateLaunchCmd).Standalone()

		evidently_updateLaunchCmd.Flags().String("description", "", "An optional description for the launch.")
		evidently_updateLaunchCmd.Flags().String("groups", "", "An array of structures that contains the feature and variations that are to be used for the launch.")
		evidently_updateLaunchCmd.Flags().String("launch", "", "The name of the launch that is to be updated.")
		evidently_updateLaunchCmd.Flags().String("metric-monitors", "", "An array of structures that define the metrics that will be used to monitor the launch performance.")
		evidently_updateLaunchCmd.Flags().String("project", "", "The name or ARN of the project that contains the launch that you want to update.")
		evidently_updateLaunchCmd.Flags().String("randomization-salt", "", "When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served.")
		evidently_updateLaunchCmd.Flags().String("scheduled-splits-config", "", "An array of structures that define the traffic allocation percentages among the feature variations during each step of the launch.")
		evidently_updateLaunchCmd.MarkFlagRequired("launch")
		evidently_updateLaunchCmd.MarkFlagRequired("project")
	})
	evidentlyCmd.AddCommand(evidently_updateLaunchCmd)
}
