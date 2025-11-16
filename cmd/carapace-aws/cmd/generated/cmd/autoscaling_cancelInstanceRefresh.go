package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_cancelInstanceRefreshCmd = &cobra.Command{
	Use:   "cancel-instance-refresh",
	Short: "Cancels an instance refresh or rollback that is in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_cancelInstanceRefreshCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_cancelInstanceRefreshCmd).Standalone()

		autoscaling_cancelInstanceRefreshCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_cancelInstanceRefreshCmd.Flags().String("wait-for-transitioning-instances", "", "When cancelling an instance refresh, this indicates whether to wait for in-flight launches and terminations to complete.")
		autoscaling_cancelInstanceRefreshCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_cancelInstanceRefreshCmd)
}
