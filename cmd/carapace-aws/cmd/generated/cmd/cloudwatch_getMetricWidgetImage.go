package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_getMetricWidgetImageCmd = &cobra.Command{
	Use:   "get-metric-widget-image",
	Short: "You can use the `GetMetricWidgetImage` API to retrieve a snapshot graph of one or more Amazon CloudWatch metrics as a bitmap image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_getMetricWidgetImageCmd).Standalone()

	cloudwatch_getMetricWidgetImageCmd.Flags().String("metric-widget", "", "A JSON string that defines the bitmap graph to be retrieved.")
	cloudwatch_getMetricWidgetImageCmd.Flags().String("output-format", "", "The format of the resulting image.")
	cloudwatch_getMetricWidgetImageCmd.MarkFlagRequired("metric-widget")
	cloudwatchCmd.AddCommand(cloudwatch_getMetricWidgetImageCmd)
}
