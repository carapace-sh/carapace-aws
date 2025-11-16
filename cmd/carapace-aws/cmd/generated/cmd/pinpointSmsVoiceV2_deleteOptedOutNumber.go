package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteOptedOutNumberCmd = &cobra.Command{
	Use:   "delete-opted-out-number",
	Short: "Deletes an existing opted out destination phone number from the specified opt-out list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteOptedOutNumberCmd).Standalone()

	pinpointSmsVoiceV2_deleteOptedOutNumberCmd.Flags().String("opt-out-list-name", "", "The OptOutListName or OptOutListArn to remove the phone number from.")
	pinpointSmsVoiceV2_deleteOptedOutNumberCmd.Flags().String("opted-out-number", "", "The phone number, in E.164 format, to remove from the OptOutList.")
	pinpointSmsVoiceV2_deleteOptedOutNumberCmd.MarkFlagRequired("opt-out-list-name")
	pinpointSmsVoiceV2_deleteOptedOutNumberCmd.MarkFlagRequired("opted-out-number")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteOptedOutNumberCmd)
}
