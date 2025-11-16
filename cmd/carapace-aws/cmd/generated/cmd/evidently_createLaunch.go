package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_createLaunchCmd = &cobra.Command{
	Use:   "create-launch",
	Short: "Creates a *launch* of a given feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_createLaunchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_createLaunchCmd).Standalone()

		evidently_createLaunchCmd.Flags().String("description", "", "An optional description for the launch.")
		evidently_createLaunchCmd.Flags().String("groups", "", "An array of structures that contains the feature and variations that are to be used for the launch.")
		evidently_createLaunchCmd.Flags().String("metric-monitors", "", "An array of structures that define the metrics that will be used to monitor the launch performance.")
		evidently_createLaunchCmd.Flags().String("name", "", "The name for the new launch.")
		evidently_createLaunchCmd.Flags().String("project", "", "The name or ARN of the project that you want to create the launch in.")
		evidently_createLaunchCmd.Flags().String("randomization-salt", "", "When Evidently assigns a particular user session to a launch, it must use a randomization ID to determine which variation the user session is served.")
		evidently_createLaunchCmd.Flags().String("scheduled-splits-config", "", "An array of structures that define the traffic allocation percentages among the feature variations during each step of the launch.")
		evidently_createLaunchCmd.Flags().String("tags", "", "Assigns one or more tags (key-value pairs) to the launch.")
		evidently_createLaunchCmd.MarkFlagRequired("groups")
		evidently_createLaunchCmd.MarkFlagRequired("name")
		evidently_createLaunchCmd.MarkFlagRequired("project")
	})
	evidentlyCmd.AddCommand(evidently_createLaunchCmd)
}
