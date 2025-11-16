package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_startLaunchCmd = &cobra.Command{
	Use:   "start-launch",
	Short: "Starts an existing launch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_startLaunchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_startLaunchCmd).Standalone()

		evidently_startLaunchCmd.Flags().String("launch", "", "The name of the launch to start.")
		evidently_startLaunchCmd.Flags().String("project", "", "The name or ARN of the project that contains the launch to start.")
		evidently_startLaunchCmd.MarkFlagRequired("launch")
		evidently_startLaunchCmd.MarkFlagRequired("project")
	})
	evidentlyCmd.AddCommand(evidently_startLaunchCmd)
}
