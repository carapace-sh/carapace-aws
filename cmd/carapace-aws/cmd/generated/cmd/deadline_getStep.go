package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getStepCmd = &cobra.Command{
	Use:   "get-step",
	Short: "Gets a step.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getStepCmd).Standalone()

	deadline_getStepCmd.Flags().String("farm-id", "", "The farm ID for the step.")
	deadline_getStepCmd.Flags().String("job-id", "", "The job ID for the step.")
	deadline_getStepCmd.Flags().String("queue-id", "", "The queue ID for the step.")
	deadline_getStepCmd.Flags().String("step-id", "", "The step ID.")
	deadline_getStepCmd.MarkFlagRequired("farm-id")
	deadline_getStepCmd.MarkFlagRequired("job-id")
	deadline_getStepCmd.MarkFlagRequired("queue-id")
	deadline_getStepCmd.MarkFlagRequired("step-id")
	deadlineCmd.AddCommand(deadline_getStepCmd)
}
