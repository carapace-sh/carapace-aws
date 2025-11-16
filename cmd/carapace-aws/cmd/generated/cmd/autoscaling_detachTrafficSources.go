package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_detachTrafficSourcesCmd = &cobra.Command{
	Use:   "detach-traffic-sources",
	Short: "Detaches one or more traffic sources from the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_detachTrafficSourcesCmd).Standalone()

	autoscaling_detachTrafficSourcesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_detachTrafficSourcesCmd.Flags().String("traffic-sources", "", "The unique identifiers of one or more traffic sources.")
	autoscaling_detachTrafficSourcesCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_detachTrafficSourcesCmd.MarkFlagRequired("traffic-sources")
	autoscalingCmd.AddCommand(autoscaling_detachTrafficSourcesCmd)
}
