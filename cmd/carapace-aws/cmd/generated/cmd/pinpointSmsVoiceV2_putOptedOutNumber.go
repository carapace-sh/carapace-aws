package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_putOptedOutNumberCmd = &cobra.Command{
	Use:   "put-opted-out-number",
	Short: "Creates an opted out destination phone number in the opt-out list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_putOptedOutNumberCmd).Standalone()

	pinpointSmsVoiceV2_putOptedOutNumberCmd.Flags().String("opt-out-list-name", "", "The OptOutListName or OptOutListArn to add the phone number to.")
	pinpointSmsVoiceV2_putOptedOutNumberCmd.Flags().String("opted-out-number", "", "The phone number to add to the OptOutList in E.164 format.")
	pinpointSmsVoiceV2_putOptedOutNumberCmd.MarkFlagRequired("opt-out-list-name")
	pinpointSmsVoiceV2_putOptedOutNumberCmd.MarkFlagRequired("opted-out-number")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_putOptedOutNumberCmd)
}
