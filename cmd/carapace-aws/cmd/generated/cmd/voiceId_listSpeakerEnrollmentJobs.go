package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_listSpeakerEnrollmentJobsCmd = &cobra.Command{
	Use:   "list-speaker-enrollment-jobs",
	Short: "Lists all the speaker enrollment jobs in the domain with the specified `JobStatus`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_listSpeakerEnrollmentJobsCmd).Standalone()

	voiceId_listSpeakerEnrollmentJobsCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the speaker enrollment jobs.")
	voiceId_listSpeakerEnrollmentJobsCmd.Flags().String("job-status", "", "Provides the status of your speaker enrollment Job.")
	voiceId_listSpeakerEnrollmentJobsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	voiceId_listSpeakerEnrollmentJobsCmd.Flags().String("next-token", "", "If `NextToken` is returned, there are more results available.")
	voiceId_listSpeakerEnrollmentJobsCmd.MarkFlagRequired("domain-id")
	voiceIdCmd.AddCommand(voiceId_listSpeakerEnrollmentJobsCmd)
}
