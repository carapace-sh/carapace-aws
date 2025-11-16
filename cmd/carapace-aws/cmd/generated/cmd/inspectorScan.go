package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectorScanCmd = &cobra.Command{
	Use:   "inspector-scan",
	Short: "Amazon Inspector Scan is a vulnerability discovery service that scans a provided Software Bill of Materials (SBOM) for security vulnerabilities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectorScanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspectorScanCmd).Standalone()

	})
	rootCmd.AddCommand(inspectorScanCmd)
}
