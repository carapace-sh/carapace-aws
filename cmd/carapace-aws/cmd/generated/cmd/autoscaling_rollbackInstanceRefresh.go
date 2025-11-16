package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_rollbackInstanceRefreshCmd = &cobra.Command{
	Use:   "rollback-instance-refresh",
	Short: "Cancels an instance refresh that is in progress and rolls back any changes that it made.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_rollbackInstanceRefreshCmd).Standalone()

	autoscaling_rollbackInstanceRefreshCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_rollbackInstanceRefreshCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscalingCmd.AddCommand(autoscaling_rollbackInstanceRefreshCmd)
}
