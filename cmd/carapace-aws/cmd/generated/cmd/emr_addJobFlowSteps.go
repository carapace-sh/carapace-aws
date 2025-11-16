package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_addJobFlowStepsCmd = &cobra.Command{
	Use:   "add-job-flow-steps",
	Short: "AddJobFlowSteps adds new steps to a running cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_addJobFlowStepsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_addJobFlowStepsCmd).Standalone()

		emr_addJobFlowStepsCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the runtime role for a step on the cluster.")
		emr_addJobFlowStepsCmd.Flags().String("job-flow-id", "", "A string that uniquely identifies the job flow.")
		emr_addJobFlowStepsCmd.Flags().String("steps", "", "A list of [StepConfig]() to be executed by the job flow.")
		emr_addJobFlowStepsCmd.MarkFlagRequired("job-flow-id")
		emr_addJobFlowStepsCmd.MarkFlagRequired("steps")
	})
	emrCmd.AddCommand(emr_addJobFlowStepsCmd)
}
