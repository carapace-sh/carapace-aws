package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putIndexPolicyCmd = &cobra.Command{
	Use:   "put-index-policy",
	Short: "Creates or updates a *field index policy* for the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putIndexPolicyCmd).Standalone()

	logs_putIndexPolicyCmd.Flags().String("log-group-identifier", "", "Specify either the log group name or log group ARN to apply this field index policy to.")
	logs_putIndexPolicyCmd.Flags().String("policy-document", "", "The index policy document, in JSON format.")
	logs_putIndexPolicyCmd.MarkFlagRequired("log-group-identifier")
	logs_putIndexPolicyCmd.MarkFlagRequired("policy-document")
	logsCmd.AddCommand(logs_putIndexPolicyCmd)
}
