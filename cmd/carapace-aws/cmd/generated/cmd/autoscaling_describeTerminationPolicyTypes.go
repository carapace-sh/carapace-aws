package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeTerminationPolicyTypesCmd = &cobra.Command{
	Use:   "describe-termination-policy-types",
	Short: "Describes the termination policies supported by Amazon EC2 Auto Scaling.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeTerminationPolicyTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeTerminationPolicyTypesCmd).Standalone()

	})
	autoscalingCmd.AddCommand(autoscaling_describeTerminationPolicyTypesCmd)
}
