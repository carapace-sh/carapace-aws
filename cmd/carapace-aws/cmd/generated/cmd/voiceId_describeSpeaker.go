package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_describeSpeakerCmd = &cobra.Command{
	Use:   "describe-speaker",
	Short: "Describes the specified speaker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_describeSpeakerCmd).Standalone()

	voiceId_describeSpeakerCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the speaker.")
	voiceId_describeSpeakerCmd.Flags().String("speaker-id", "", "The identifier of the speaker you are describing.")
	voiceId_describeSpeakerCmd.MarkFlagRequired("domain-id")
	voiceId_describeSpeakerCmd.MarkFlagRequired("speaker-id")
	voiceIdCmd.AddCommand(voiceId_describeSpeakerCmd)
}
