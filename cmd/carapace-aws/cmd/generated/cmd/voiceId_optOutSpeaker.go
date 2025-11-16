package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_optOutSpeakerCmd = &cobra.Command{
	Use:   "opt-out-speaker",
	Short: "Opts out a speaker from Voice ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_optOutSpeakerCmd).Standalone()

	voiceId_optOutSpeakerCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the speaker.")
	voiceId_optOutSpeakerCmd.Flags().String("speaker-id", "", "The identifier of the speaker you want opted-out.")
	voiceId_optOutSpeakerCmd.MarkFlagRequired("domain-id")
	voiceId_optOutSpeakerCmd.MarkFlagRequired("speaker-id")
	voiceIdCmd.AddCommand(voiceId_optOutSpeakerCmd)
}
