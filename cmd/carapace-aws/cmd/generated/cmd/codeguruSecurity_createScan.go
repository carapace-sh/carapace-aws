package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_createScanCmd = &cobra.Command{
	Use:   "create-scan",
	Short: "Use to create a scan using code uploaded to an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_createScanCmd).Standalone()

	codeguruSecurity_createScanCmd.Flags().String("analysis-type", "", "The type of analysis you want CodeGuru Security to perform in the scan, either `Security` or `All`.")
	codeguruSecurity_createScanCmd.Flags().String("client-token", "", "The idempotency token for the request.")
	codeguruSecurity_createScanCmd.Flags().String("resource-id", "", "The identifier for the resource object to be scanned.")
	codeguruSecurity_createScanCmd.Flags().String("scan-name", "", "The unique name that CodeGuru Security uses to track revisions across multiple scans of the same resource.")
	codeguruSecurity_createScanCmd.Flags().String("scan-type", "", "The type of scan, either `Standard` or `Express`.")
	codeguruSecurity_createScanCmd.Flags().String("tags", "", "An array of key-value pairs used to tag a scan.")
	codeguruSecurity_createScanCmd.MarkFlagRequired("resource-id")
	codeguruSecurity_createScanCmd.MarkFlagRequired("scan-name")
	codeguruSecurityCmd.AddCommand(codeguruSecurity_createScanCmd)
}
