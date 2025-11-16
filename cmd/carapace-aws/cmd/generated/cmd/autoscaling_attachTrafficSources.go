package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_attachTrafficSourcesCmd = &cobra.Command{
	Use:   "attach-traffic-sources",
	Short: "Attaches one or more traffic sources to the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_attachTrafficSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_attachTrafficSourcesCmd).Standalone()

		autoscaling_attachTrafficSourcesCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_attachTrafficSourcesCmd.Flags().String("skip-zonal-shift-validation", "", "If you enable zonal shift with cross-zone disabled load balancers, capacity could become imbalanced across Availability Zones.")
		autoscaling_attachTrafficSourcesCmd.Flags().String("traffic-sources", "", "The unique identifiers of one or more traffic sources.")
		autoscaling_attachTrafficSourcesCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_attachTrafficSourcesCmd.MarkFlagRequired("traffic-sources")
	})
	autoscalingCmd.AddCommand(autoscaling_attachTrafficSourcesCmd)
}
