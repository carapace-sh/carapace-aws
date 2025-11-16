package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_describeMetricsExportCmd = &cobra.Command{
	Use:   "describe-metrics-export",
	Short: "Describes the metrics of the application configuration being exported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_describeMetricsExportCmd).Standalone()

	resiliencehub_describeMetricsExportCmd.Flags().String("metrics-export-id", "", "Identifier of the metrics export task.")
	resiliencehub_describeMetricsExportCmd.MarkFlagRequired("metrics-export-id")
	resiliencehubCmd.AddCommand(resiliencehub_describeMetricsExportCmd)
}
