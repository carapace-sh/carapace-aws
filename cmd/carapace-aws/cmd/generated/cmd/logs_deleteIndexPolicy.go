package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteIndexPolicyCmd = &cobra.Command{
	Use:   "delete-index-policy",
	Short: "Deletes a log-group level field index policy that was applied to a single log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteIndexPolicyCmd).Standalone()

	logs_deleteIndexPolicyCmd.Flags().String("log-group-identifier", "", "The log group to delete the index policy for.")
	logs_deleteIndexPolicyCmd.MarkFlagRequired("log-group-identifier")
	logsCmd.AddCommand(logs_deleteIndexPolicyCmd)
}
