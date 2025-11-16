package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeAccountLimitsCmd = &cobra.Command{
	Use:   "describe-account-limits",
	Short: "Describes the current Amazon EC2 Auto Scaling resource quotas for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeAccountLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeAccountLimitsCmd).Standalone()

	})
	autoscalingCmd.AddCommand(autoscaling_describeAccountLimitsCmd)
}
