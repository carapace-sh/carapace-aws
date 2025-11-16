package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectorScan_scanSbomCmd = &cobra.Command{
	Use:   "scan-sbom",
	Short: "Scans a provided CycloneDX 1.5 SBOM and reports on any vulnerabilities discovered in that SBOM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectorScan_scanSbomCmd).Standalone()

	inspectorScan_scanSbomCmd.Flags().String("output-format", "", "The output format for the vulnerability report.")
	inspectorScan_scanSbomCmd.Flags().String("sbom", "", "The JSON file for the SBOM you want to scan.")
	inspectorScan_scanSbomCmd.MarkFlagRequired("sbom")
	inspectorScanCmd.AddCommand(inspectorScan_scanSbomCmd)
}
