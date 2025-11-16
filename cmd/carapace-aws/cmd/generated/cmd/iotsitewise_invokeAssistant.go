package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_invokeAssistantCmd = &cobra.Command{
	Use:   "invoke-assistant",
	Short: "Invokes SiteWise Assistant to start or continue a conversation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_invokeAssistantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_invokeAssistantCmd).Standalone()

		iotsitewise_invokeAssistantCmd.Flags().String("conversation-id", "", "The ID assigned to a conversation.")
		iotsitewise_invokeAssistantCmd.Flags().String("enable-trace", "", "Specifies if to turn trace on or not.")
		iotsitewise_invokeAssistantCmd.Flags().String("message", "", "A text message sent to the SiteWise Assistant by the user.")
		iotsitewise_invokeAssistantCmd.MarkFlagRequired("message")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_invokeAssistantCmd)
}
