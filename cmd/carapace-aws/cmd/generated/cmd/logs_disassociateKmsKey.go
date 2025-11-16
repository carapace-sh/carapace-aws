package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_disassociateKmsKeyCmd = &cobra.Command{
	Use:   "disassociate-kms-key",
	Short: "Disassociates the specified KMS key from the specified log group or from all CloudWatch Logs Insights query results in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_disassociateKmsKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_disassociateKmsKeyCmd).Standalone()

		logs_disassociateKmsKeyCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_disassociateKmsKeyCmd.Flags().String("resource-identifier", "", "Specifies the target for this operation.")
	})
	logsCmd.AddCommand(logs_disassociateKmsKeyCmd)
}
