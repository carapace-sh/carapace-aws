package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putRetentionPolicyCmd = &cobra.Command{
	Use:   "put-retention-policy",
	Short: "Sets the retention of the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putRetentionPolicyCmd).Standalone()

	logs_putRetentionPolicyCmd.Flags().String("log-group-name", "", "The name of the log group.")
	logs_putRetentionPolicyCmd.Flags().String("retention-in-days", "", "")
	logs_putRetentionPolicyCmd.MarkFlagRequired("log-group-name")
	logs_putRetentionPolicyCmd.MarkFlagRequired("retention-in-days")
	logsCmd.AddCommand(logs_putRetentionPolicyCmd)
}
