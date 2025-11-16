package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_startMetricsExportCmd = &cobra.Command{
	Use:   "start-metrics-export",
	Short: "Initiates the export task of metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_startMetricsExportCmd).Standalone()

	resiliencehub_startMetricsExportCmd.Flags().String("bucket-name", "", "(Optional) Specifies the name of the Amazon Simple Storage Service bucket where the exported metrics will be stored.")
	resiliencehub_startMetricsExportCmd.Flags().String("client-token", "", "Used for an idempotency token.")
	resiliencehubCmd.AddCommand(resiliencehub_startMetricsExportCmd)
}
