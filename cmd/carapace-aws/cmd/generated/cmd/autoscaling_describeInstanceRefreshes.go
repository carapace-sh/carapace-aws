package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeInstanceRefreshesCmd = &cobra.Command{
	Use:   "describe-instance-refreshes",
	Short: "Gets information about the instance refreshes for the specified Auto Scaling group from the previous six weeks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeInstanceRefreshesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeInstanceRefreshesCmd).Standalone()

		autoscaling_describeInstanceRefreshesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_describeInstanceRefreshesCmd.Flags().String("instance-refresh-ids", "", "One or more instance refresh IDs.")
		autoscaling_describeInstanceRefreshesCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
		autoscaling_describeInstanceRefreshesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		autoscaling_describeInstanceRefreshesCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_describeInstanceRefreshesCmd)
}
