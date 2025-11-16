package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_cancelStepsCmd = &cobra.Command{
	Use:   "cancel-steps",
	Short: "Cancels a pending step or steps in a running cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_cancelStepsCmd).Standalone()

	emr_cancelStepsCmd.Flags().String("cluster-id", "", "The `ClusterID` for the specified steps that will be canceled.")
	emr_cancelStepsCmd.Flags().String("step-cancellation-option", "", "The option to choose to cancel `RUNNING` steps.")
	emr_cancelStepsCmd.Flags().String("step-ids", "", "The list of `StepIDs` to cancel.")
	emr_cancelStepsCmd.MarkFlagRequired("cluster-id")
	emr_cancelStepsCmd.MarkFlagRequired("step-ids")
	emrCmd.AddCommand(emr_cancelStepsCmd)
}
