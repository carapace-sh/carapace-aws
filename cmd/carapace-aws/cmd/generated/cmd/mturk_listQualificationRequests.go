package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listQualificationRequestsCmd = &cobra.Command{
	Use:   "list-qualification-requests",
	Short: "The `ListQualificationRequests` operation retrieves requests for Qualifications of a particular Qualification type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listQualificationRequestsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_listQualificationRequestsCmd).Standalone()

		mturk_listQualificationRequestsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		mturk_listQualificationRequestsCmd.Flags().String("next-token", "", "")
		mturk_listQualificationRequestsCmd.Flags().String("qualification-type-id", "", "The ID of the QualificationType.")
	})
	mturkCmd.AddCommand(mturk_listQualificationRequestsCmd)
}
