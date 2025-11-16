package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_associateKmsKeyCmd = &cobra.Command{
	Use:   "associate-kms-key",
	Short: "Associates the specified KMS key with either one log group in the account, or with all stored CloudWatch Logs query insights results in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_associateKmsKeyCmd).Standalone()

	logs_associateKmsKeyCmd.Flags().String("kms-key-id", "", "The Amazon Resource Name (ARN) of the KMS key to use when encrypting log data.")
	logs_associateKmsKeyCmd.Flags().String("log-group-name", "", "The name of the log group.")
	logs_associateKmsKeyCmd.Flags().String("resource-identifier", "", "Specifies the target for this operation.")
	logs_associateKmsKeyCmd.MarkFlagRequired("kms-key-id")
	logsCmd.AddCommand(logs_associateKmsKeyCmd)
}
