package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_stopLaunchCmd = &cobra.Command{
	Use:   "stop-launch",
	Short: "Stops a launch that is currently running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_stopLaunchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_stopLaunchCmd).Standalone()

		evidently_stopLaunchCmd.Flags().String("desired-state", "", "Specify whether to consider the launch as `COMPLETED` or `CANCELLED` after it stops.")
		evidently_stopLaunchCmd.Flags().String("launch", "", "The name of the launch to stop.")
		evidently_stopLaunchCmd.Flags().String("project", "", "The name or ARN of the project that contains the launch that you want to stop.")
		evidently_stopLaunchCmd.Flags().String("reason", "", "A string that describes why you are stopping the launch.")
		evidently_stopLaunchCmd.MarkFlagRequired("launch")
		evidently_stopLaunchCmd.MarkFlagRequired("project")
	})
	evidentlyCmd.AddCommand(evidently_stopLaunchCmd)
}
