package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_createOptOutListCmd = &cobra.Command{
	Use:   "create-opt-out-list",
	Short: "Creates a new opt-out list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_createOptOutListCmd).Standalone()

	pinpointSmsVoiceV2_createOptOutListCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	pinpointSmsVoiceV2_createOptOutListCmd.Flags().String("opt-out-list-name", "", "The name of the new OptOutList.")
	pinpointSmsVoiceV2_createOptOutListCmd.Flags().String("tags", "", "An array of tags (key and value pairs) to associate with the new OptOutList.")
	pinpointSmsVoiceV2_createOptOutListCmd.MarkFlagRequired("opt-out-list-name")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_createOptOutListCmd)
}
