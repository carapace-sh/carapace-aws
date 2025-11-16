package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_startResourceScanCmd = &cobra.Command{
	Use:   "start-resource-scan",
	Short: "Starts a scan of the resources in this account in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_startResourceScanCmd).Standalone()

	cloudformation_startResourceScanCmd.Flags().String("client-request-token", "", "A unique identifier for this `StartResourceScan` request.")
	cloudformation_startResourceScanCmd.Flags().String("scan-filters", "", "The scan filters to use.")
	cloudformationCmd.AddCommand(cloudformation_startResourceScanCmd)
}
