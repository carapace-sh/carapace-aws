package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_acceptQualificationRequestCmd = &cobra.Command{
	Use:   "accept-qualification-request",
	Short: "The `AcceptQualificationRequest` operation approves a Worker's request for a Qualification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_acceptQualificationRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_acceptQualificationRequestCmd).Standalone()

		mturk_acceptQualificationRequestCmd.Flags().String("integer-value", "", "The value of the Qualification.")
		mturk_acceptQualificationRequestCmd.Flags().String("qualification-request-id", "", "The ID of the Qualification request, as returned by the `GetQualificationRequests` operation.")
		mturk_acceptQualificationRequestCmd.MarkFlagRequired("qualification-request-id")
	})
	mturkCmd.AddCommand(mturk_acceptQualificationRequestCmd)
}
