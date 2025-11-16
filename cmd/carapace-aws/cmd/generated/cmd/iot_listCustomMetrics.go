package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listCustomMetricsCmd = &cobra.Command{
	Use:   "list-custom-metrics",
	Short: "Lists your Device Defender detect custom metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listCustomMetricsCmd).Standalone()

	iot_listCustomMetricsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listCustomMetricsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iotCmd.AddCommand(iot_listCustomMetricsCmd)
}
