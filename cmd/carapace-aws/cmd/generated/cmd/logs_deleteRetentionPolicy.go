package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteRetentionPolicyCmd = &cobra.Command{
	Use:   "delete-retention-policy",
	Short: "Deletes the specified retention policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteRetentionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteRetentionPolicyCmd).Standalone()

		logs_deleteRetentionPolicyCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_deleteRetentionPolicyCmd.MarkFlagRequired("log-group-name")
	})
	logsCmd.AddCommand(logs_deleteRetentionPolicyCmd)
}
