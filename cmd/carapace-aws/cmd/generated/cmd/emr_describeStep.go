package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_describeStepCmd = &cobra.Command{
	Use:   "describe-step",
	Short: "Provides more detail about the cluster step.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_describeStepCmd).Standalone()

	emr_describeStepCmd.Flags().String("cluster-id", "", "The identifier of the cluster with steps to describe.")
	emr_describeStepCmd.Flags().String("step-id", "", "The identifier of the step to describe.")
	emr_describeStepCmd.MarkFlagRequired("cluster-id")
	emr_describeStepCmd.MarkFlagRequired("step-id")
	emrCmd.AddCommand(emr_describeStepCmd)
}
