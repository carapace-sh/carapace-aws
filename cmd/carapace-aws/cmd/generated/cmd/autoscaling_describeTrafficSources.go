package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeTrafficSourcesCmd = &cobra.Command{
	Use:   "describe-traffic-sources",
	Short: "Gets information about the traffic sources for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeTrafficSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeTrafficSourcesCmd).Standalone()

		autoscaling_describeTrafficSourcesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_describeTrafficSourcesCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
		autoscaling_describeTrafficSourcesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		autoscaling_describeTrafficSourcesCmd.Flags().String("traffic-source-type", "", "The traffic source type that you want to describe.")
		autoscaling_describeTrafficSourcesCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_describeTrafficSourcesCmd)
}
