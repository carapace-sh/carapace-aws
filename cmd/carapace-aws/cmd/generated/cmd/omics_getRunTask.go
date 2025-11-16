package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getRunTaskCmd = &cobra.Command{
	Use:   "get-run-task",
	Short: "Gets detailed information about a run task using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getRunTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_getRunTaskCmd).Standalone()

		omics_getRunTaskCmd.Flags().String("id", "", "The workflow run ID.")
		omics_getRunTaskCmd.Flags().String("task-id", "", "The task's ID.")
		omics_getRunTaskCmd.MarkFlagRequired("id")
		omics_getRunTaskCmd.MarkFlagRequired("task-id")
	})
	omicsCmd.AddCommand(omics_getRunTaskCmd)
}
