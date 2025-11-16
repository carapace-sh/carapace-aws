package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listHitsForQualificationTypeCmd = &cobra.Command{
	Use:   "list-hits-for-qualification-type",
	Short: "The `ListHITsForQualificationType` operation returns the HITs that use the given Qualification type for a Qualification requirement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listHitsForQualificationTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_listHitsForQualificationTypeCmd).Standalone()

		mturk_listHitsForQualificationTypeCmd.Flags().String("max-results", "", "Limit the number of results returned.")
		mturk_listHitsForQualificationTypeCmd.Flags().String("next-token", "", "Pagination Token")
		mturk_listHitsForQualificationTypeCmd.Flags().String("qualification-type-id", "", "The ID of the Qualification type to use when querying HITs.")
		mturk_listHitsForQualificationTypeCmd.MarkFlagRequired("qualification-type-id")
	})
	mturkCmd.AddCommand(mturk_listHitsForQualificationTypeCmd)
}
