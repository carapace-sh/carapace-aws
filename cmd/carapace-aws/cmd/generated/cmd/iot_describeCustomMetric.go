package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeCustomMetricCmd = &cobra.Command{
	Use:   "describe-custom-metric",
	Short: "Gets information about a Device Defender detect custom metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeCustomMetricCmd).Standalone()

	iot_describeCustomMetricCmd.Flags().String("metric-name", "", "The name of the custom metric.")
	iot_describeCustomMetricCmd.MarkFlagRequired("metric-name")
	iotCmd.AddCommand(iot_describeCustomMetricCmd)
}
