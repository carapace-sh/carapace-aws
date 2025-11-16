package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteTransformerCmd = &cobra.Command{
	Use:   "delete-transformer",
	Short: "Deletes the log transformer for the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteTransformerCmd).Standalone()

	logs_deleteTransformerCmd.Flags().String("log-group-identifier", "", "Specify either the name or ARN of the log group to delete the transformer for.")
	logs_deleteTransformerCmd.MarkFlagRequired("log-group-identifier")
	logsCmd.AddCommand(logs_deleteTransformerCmd)
}
