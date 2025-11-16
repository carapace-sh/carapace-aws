package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createVoiceProfileCmd = &cobra.Command{
	Use:   "create-voice-profile",
	Short: "Creates a voice profile, which consists of an enrolled user and their latest voice print.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createVoiceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_createVoiceProfileCmd).Standalone()

		chimeSdkVoice_createVoiceProfileCmd.Flags().String("speaker-search-task-id", "", "The ID of the speaker search task.")
		chimeSdkVoice_createVoiceProfileCmd.MarkFlagRequired("speaker-search-task-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createVoiceProfileCmd)
}
