package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeMetricCollectionTypesCmd = &cobra.Command{
	Use:   "describe-metric-collection-types",
	Short: "Describes the available CloudWatch metrics for Amazon EC2 Auto Scaling.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeMetricCollectionTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeMetricCollectionTypesCmd).Standalone()

	})
	autoscalingCmd.AddCommand(autoscaling_describeMetricCollectionTypesCmd)
}
