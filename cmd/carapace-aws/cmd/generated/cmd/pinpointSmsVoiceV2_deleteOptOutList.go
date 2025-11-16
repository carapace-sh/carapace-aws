package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointSmsVoiceV2_deleteOptOutListCmd = &cobra.Command{
	Use:   "delete-opt-out-list",
	Short: "Deletes an existing opt-out list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointSmsVoiceV2_deleteOptOutListCmd).Standalone()

	pinpointSmsVoiceV2_deleteOptOutListCmd.Flags().String("opt-out-list-name", "", "The OptOutListName or OptOutListArn of the OptOutList to delete.")
	pinpointSmsVoiceV2_deleteOptOutListCmd.MarkFlagRequired("opt-out-list-name")
	pinpointSmsVoiceV2Cmd.AddCommand(pinpointSmsVoiceV2_deleteOptOutListCmd)
}
