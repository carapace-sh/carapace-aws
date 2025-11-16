package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmailmessageflow_putRawMessageContentCmd = &cobra.Command{
	Use:   "put-raw-message-content",
	Short: "Updates the raw content of an in-transit email message, in MIME format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmailmessageflow_putRawMessageContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmailmessageflow_putRawMessageContentCmd).Standalone()

		workmailmessageflow_putRawMessageContentCmd.Flags().String("content", "", "Describes the raw message content of the updated email message.")
		workmailmessageflow_putRawMessageContentCmd.Flags().String("message-id", "", "The identifier of the email message being updated.")
		workmailmessageflow_putRawMessageContentCmd.MarkFlagRequired("content")
		workmailmessageflow_putRawMessageContentCmd.MarkFlagRequired("message-id")
	})
	workmailmessageflowCmd.AddCommand(workmailmessageflow_putRawMessageContentCmd)
}
