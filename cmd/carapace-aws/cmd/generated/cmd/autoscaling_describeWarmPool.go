package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeWarmPoolCmd = &cobra.Command{
	Use:   "describe-warm-pool",
	Short: "Gets information about a warm pool and its instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeWarmPoolCmd).Standalone()

	autoscaling_describeWarmPoolCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_describeWarmPoolCmd.Flags().String("max-records", "", "The maximum number of instances to return with this call.")
	autoscaling_describeWarmPoolCmd.Flags().String("next-token", "", "The token for the next set of instances to return.")
	autoscaling_describeWarmPoolCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscalingCmd.AddCommand(autoscaling_describeWarmPoolCmd)
}
