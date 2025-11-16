package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_createLogGroupCmd = &cobra.Command{
	Use:   "create-log-group",
	Short: "Creates a log group with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_createLogGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_createLogGroupCmd).Standalone()

		logs_createLogGroupCmd.Flags().String("kms-key-id", "", "The Amazon Resource Name (ARN) of the KMS key to use when encrypting log data.")
		logs_createLogGroupCmd.Flags().String("log-group-class", "", "Use this parameter to specify the log group class for this log group.")
		logs_createLogGroupCmd.Flags().String("log-group-name", "", "A name for the log group.")
		logs_createLogGroupCmd.Flags().String("tags", "", "The key-value pairs to use for the tags.")
		logs_createLogGroupCmd.MarkFlagRequired("log-group-name")
	})
	logsCmd.AddCommand(logs_createLogGroupCmd)
}
