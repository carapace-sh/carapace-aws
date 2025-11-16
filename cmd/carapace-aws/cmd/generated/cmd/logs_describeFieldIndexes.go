package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeFieldIndexesCmd = &cobra.Command{
	Use:   "describe-field-indexes",
	Short: "Returns a list of custom and default field indexes which are discovered in log data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeFieldIndexesCmd).Standalone()

	logs_describeFieldIndexesCmd.Flags().String("log-group-identifiers", "", "An array containing the names or ARNs of the log groups that you want to retrieve field indexes for.")
	logs_describeFieldIndexesCmd.Flags().String("next-token", "", "")
	logs_describeFieldIndexesCmd.MarkFlagRequired("log-group-identifiers")
	logsCmd.AddCommand(logs_describeFieldIndexesCmd)
}
