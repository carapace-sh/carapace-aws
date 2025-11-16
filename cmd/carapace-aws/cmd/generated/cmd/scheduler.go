package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "Amazon EventBridge Scheduler is a serverless scheduler that allows you to create, run, and manage tasks from one central, managed service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schedulerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schedulerCmd).Standalone()

	})
	rootCmd.AddCommand(schedulerCmd)
}
