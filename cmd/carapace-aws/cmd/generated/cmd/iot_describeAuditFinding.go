package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeAuditFindingCmd = &cobra.Command{
	Use:   "describe-audit-finding",
	Short: "Gets information about a single audit finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeAuditFindingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeAuditFindingCmd).Standalone()

		iot_describeAuditFindingCmd.Flags().String("finding-id", "", "A unique identifier for a single audit finding.")
		iot_describeAuditFindingCmd.MarkFlagRequired("finding-id")
	})
	iotCmd.AddCommand(iot_describeAuditFindingCmd)
}
