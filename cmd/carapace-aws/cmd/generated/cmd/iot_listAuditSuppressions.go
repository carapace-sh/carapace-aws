package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listAuditSuppressionsCmd = &cobra.Command{
	Use:   "list-audit-suppressions",
	Short: "Lists your Device Defender audit listings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listAuditSuppressionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listAuditSuppressionsCmd).Standalone()

		iot_listAuditSuppressionsCmd.Flags().String("ascending-order", "", "Determines whether suppressions are listed in ascending order by expiration date or not.")
		iot_listAuditSuppressionsCmd.Flags().String("check-name", "", "")
		iot_listAuditSuppressionsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listAuditSuppressionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		iot_listAuditSuppressionsCmd.Flags().String("resource-identifier", "", "")
	})
	iotCmd.AddCommand(iot_listAuditSuppressionsCmd)
}
