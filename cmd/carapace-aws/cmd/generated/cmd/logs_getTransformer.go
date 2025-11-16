package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getTransformerCmd = &cobra.Command{
	Use:   "get-transformer",
	Short: "Returns the information about the log transformer associated with this log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getTransformerCmd).Standalone()

	logs_getTransformerCmd.Flags().String("log-group-identifier", "", "Specify either the name or ARN of the log group to return transformer information for.")
	logs_getTransformerCmd.MarkFlagRequired("log-group-identifier")
	logsCmd.AddCommand(logs_getTransformerCmd)
}
