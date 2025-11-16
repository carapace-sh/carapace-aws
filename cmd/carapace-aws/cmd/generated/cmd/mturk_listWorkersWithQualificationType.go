package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listWorkersWithQualificationTypeCmd = &cobra.Command{
	Use:   "list-workers-with-qualification-type",
	Short: "The `ListWorkersWithQualificationType` operation returns all of the Workers that have been associated with a given Qualification type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listWorkersWithQualificationTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_listWorkersWithQualificationTypeCmd).Standalone()

		mturk_listWorkersWithQualificationTypeCmd.Flags().String("max-results", "", "Limit the number of results returned.")
		mturk_listWorkersWithQualificationTypeCmd.Flags().String("next-token", "", "Pagination Token")
		mturk_listWorkersWithQualificationTypeCmd.Flags().String("qualification-type-id", "", "The ID of the Qualification type of the Qualifications to return.")
		mturk_listWorkersWithQualificationTypeCmd.Flags().String("status", "", "The status of the Qualifications to return.")
		mturk_listWorkersWithQualificationTypeCmd.MarkFlagRequired("qualification-type-id")
	})
	mturkCmd.AddCommand(mturk_listWorkersWithQualificationTypeCmd)
}
