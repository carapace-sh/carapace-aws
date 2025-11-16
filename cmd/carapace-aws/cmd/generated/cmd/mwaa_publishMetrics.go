package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mwaa_publishMetricsCmd = &cobra.Command{
	Use:   "publish-metrics",
	Short: "**Internal only**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mwaa_publishMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mwaa_publishMetricsCmd).Standalone()

		mwaa_publishMetricsCmd.Flags().String("environment-name", "", "**Internal only**.")
		mwaa_publishMetricsCmd.Flags().String("metric-data", "", "**Internal only**.")
		mwaa_publishMetricsCmd.MarkFlagRequired("environment-name")
		mwaa_publishMetricsCmd.MarkFlagRequired("metric-data")
	})
	mwaaCmd.AddCommand(mwaa_publishMetricsCmd)
}
