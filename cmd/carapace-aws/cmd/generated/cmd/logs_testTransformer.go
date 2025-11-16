package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_testTransformerCmd = &cobra.Command{
	Use:   "test-transformer",
	Short: "Use this operation to test a log transformer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_testTransformerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_testTransformerCmd).Standalone()

		logs_testTransformerCmd.Flags().String("log-event-messages", "", "An array of the raw log events that you want to use to test this transformer.")
		logs_testTransformerCmd.Flags().String("transformer-config", "", "This structure contains the configuration of this log transformer that you want to test.")
		logs_testTransformerCmd.MarkFlagRequired("log-event-messages")
		logs_testTransformerCmd.MarkFlagRequired("transformer-config")
	})
	logsCmd.AddCommand(logs_testTransformerCmd)
}
