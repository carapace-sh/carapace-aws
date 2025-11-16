package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_rejectQualificationRequestCmd = &cobra.Command{
	Use:   "reject-qualification-request",
	Short: "The `RejectQualificationRequest` operation rejects a user's request for a Qualification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_rejectQualificationRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_rejectQualificationRequestCmd).Standalone()

		mturk_rejectQualificationRequestCmd.Flags().String("qualification-request-id", "", "The ID of the Qualification request, as returned by the `ListQualificationRequests` operation.")
		mturk_rejectQualificationRequestCmd.Flags().String("reason", "", "A text message explaining why the request was rejected, to be shown to the Worker who made the request.")
		mturk_rejectQualificationRequestCmd.MarkFlagRequired("qualification-request-id")
	})
	mturkCmd.AddCommand(mturk_rejectQualificationRequestCmd)
}
