package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_getQualificationScoreCmd = &cobra.Command{
	Use:   "get-qualification-score",
	Short: "The `GetQualificationScore` operation returns the value of a Worker's Qualification for a given Qualification type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_getQualificationScoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_getQualificationScoreCmd).Standalone()

		mturk_getQualificationScoreCmd.Flags().String("qualification-type-id", "", "The ID of the QualificationType.")
		mturk_getQualificationScoreCmd.Flags().String("worker-id", "", "The ID of the Worker whose Qualification is being updated.")
		mturk_getQualificationScoreCmd.MarkFlagRequired("qualification-type-id")
		mturk_getQualificationScoreCmd.MarkFlagRequired("worker-id")
	})
	mturkCmd.AddCommand(mturk_getQualificationScoreCmd)
}
