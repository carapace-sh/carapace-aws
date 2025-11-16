package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_getScanCmd = &cobra.Command{
	Use:   "get-scan",
	Short: "Returns details about a scan, including whether or not a scan has completed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_getScanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruSecurity_getScanCmd).Standalone()

		codeguruSecurity_getScanCmd.Flags().String("run-id", "", "UUID that identifies the individual scan run you want to view details about.")
		codeguruSecurity_getScanCmd.Flags().String("scan-name", "", "The name of the scan you want to view details about.")
		codeguruSecurity_getScanCmd.MarkFlagRequired("scan-name")
	})
	codeguruSecurityCmd.AddCommand(codeguruSecurity_getScanCmd)
}
