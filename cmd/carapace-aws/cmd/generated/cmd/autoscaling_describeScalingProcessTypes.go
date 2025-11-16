package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeScalingProcessTypesCmd = &cobra.Command{
	Use:   "describe-scaling-process-types",
	Short: "Describes the scaling process types for use with the [ResumeProcesses](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ResumeProcesses.html) and [SuspendProcesses](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_SuspendProcesses.html) APIs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeScalingProcessTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeScalingProcessTypesCmd).Standalone()

	})
	autoscalingCmd.AddCommand(autoscaling_describeScalingProcessTypesCmd)
}
