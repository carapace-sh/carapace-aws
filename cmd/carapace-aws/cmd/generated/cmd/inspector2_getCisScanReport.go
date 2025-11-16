package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getCisScanReportCmd = &cobra.Command{
	Use:   "get-cis-scan-report",
	Short: "Retrieves a CIS scan report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getCisScanReportCmd).Standalone()

	inspector2_getCisScanReportCmd.Flags().String("report-format", "", "The format of the report.")
	inspector2_getCisScanReportCmd.Flags().String("scan-arn", "", "The scan ARN.")
	inspector2_getCisScanReportCmd.Flags().String("target-accounts", "", "The target accounts.")
	inspector2_getCisScanReportCmd.MarkFlagRequired("scan-arn")
	inspector2Cmd.AddCommand(inspector2_getCisScanReportCmd)
}
