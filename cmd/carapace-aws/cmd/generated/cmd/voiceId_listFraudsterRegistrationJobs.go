package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_listFraudsterRegistrationJobsCmd = &cobra.Command{
	Use:   "list-fraudster-registration-jobs",
	Short: "Lists all the fraudster registration jobs in the domain with the given `JobStatus`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_listFraudsterRegistrationJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_listFraudsterRegistrationJobsCmd).Standalone()

		voiceId_listFraudsterRegistrationJobsCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the fraudster registration Jobs.")
		voiceId_listFraudsterRegistrationJobsCmd.Flags().String("job-status", "", "Provides the status of your fraudster registration job.")
		voiceId_listFraudsterRegistrationJobsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		voiceId_listFraudsterRegistrationJobsCmd.Flags().String("next-token", "", "If `NextToken` is returned, there are more results available.")
		voiceId_listFraudsterRegistrationJobsCmd.MarkFlagRequired("domain-id")
	})
	voiceIdCmd.AddCommand(voiceId_listFraudsterRegistrationJobsCmd)
}
