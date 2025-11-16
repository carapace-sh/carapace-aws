package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getInAppMessagesCmd = &cobra.Command{
	Use:   "get-in-app-messages",
	Short: "Retrieves the in-app messages targeted for the provided endpoint ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getInAppMessagesCmd).Standalone()

	pinpoint_getInAppMessagesCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_getInAppMessagesCmd.Flags().String("endpoint-id", "", "The unique identifier for the endpoint.")
	pinpoint_getInAppMessagesCmd.MarkFlagRequired("application-id")
	pinpoint_getInAppMessagesCmd.MarkFlagRequired("endpoint-id")
	pinpointCmd.AddCommand(pinpoint_getInAppMessagesCmd)
}
