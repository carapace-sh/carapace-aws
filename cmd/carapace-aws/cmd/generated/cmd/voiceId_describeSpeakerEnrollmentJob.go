package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_describeSpeakerEnrollmentJobCmd = &cobra.Command{
	Use:   "describe-speaker-enrollment-job",
	Short: "Describes the specified speaker enrollment job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_describeSpeakerEnrollmentJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_describeSpeakerEnrollmentJobCmd).Standalone()

		voiceId_describeSpeakerEnrollmentJobCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the speaker enrollment job.")
		voiceId_describeSpeakerEnrollmentJobCmd.Flags().String("job-id", "", "The identifier of the speaker enrollment job you are describing.")
		voiceId_describeSpeakerEnrollmentJobCmd.MarkFlagRequired("domain-id")
		voiceId_describeSpeakerEnrollmentJobCmd.MarkFlagRequired("job-id")
	})
	voiceIdCmd.AddCommand(voiceId_describeSpeakerEnrollmentJobCmd)
}
