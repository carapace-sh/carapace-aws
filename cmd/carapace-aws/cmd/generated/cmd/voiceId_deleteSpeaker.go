package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_deleteSpeakerCmd = &cobra.Command{
	Use:   "delete-speaker",
	Short: "Deletes the specified speaker from Voice ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_deleteSpeakerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_deleteSpeakerCmd).Standalone()

		voiceId_deleteSpeakerCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the speaker.")
		voiceId_deleteSpeakerCmd.Flags().String("speaker-id", "", "The identifier of the speaker you want to delete.")
		voiceId_deleteSpeakerCmd.MarkFlagRequired("domain-id")
		voiceId_deleteSpeakerCmd.MarkFlagRequired("speaker-id")
	})
	voiceIdCmd.AddCommand(voiceId_deleteSpeakerCmd)
}
