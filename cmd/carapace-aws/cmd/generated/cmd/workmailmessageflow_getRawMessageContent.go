package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmailmessageflow_getRawMessageContentCmd = &cobra.Command{
	Use:   "get-raw-message-content",
	Short: "Retrieves the raw content of an in-transit email message, in MIME format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmailmessageflow_getRawMessageContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmailmessageflow_getRawMessageContentCmd).Standalone()

		workmailmessageflow_getRawMessageContentCmd.Flags().String("message-id", "", "The identifier of the email message to retrieve.")
		workmailmessageflow_getRawMessageContentCmd.MarkFlagRequired("message-id")
	})
	workmailmessageflowCmd.AddCommand(workmailmessageflow_getRawMessageContentCmd)
}
