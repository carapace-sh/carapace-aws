package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listQualificationTypesCmd = &cobra.Command{
	Use:   "list-qualification-types",
	Short: "The `ListQualificationTypes` operation returns a list of Qualification types, filtered by an optional search term.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listQualificationTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_listQualificationTypesCmd).Standalone()

		mturk_listQualificationTypesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		mturk_listQualificationTypesCmd.Flags().Bool("must-be-owned-by-caller", false, "Specifies that only Qualification types that the Requester created are returned.")
		mturk_listQualificationTypesCmd.Flags().Bool("must-be-requestable", false, "Specifies that only Qualification types that a user can request through the Amazon Mechanical Turk web site, such as by taking a Qualification test, are returned as results of the search.")
		mturk_listQualificationTypesCmd.Flags().String("next-token", "", "")
		mturk_listQualificationTypesCmd.Flags().Bool("no-must-be-owned-by-caller", false, "Specifies that only Qualification types that the Requester created are returned.")
		mturk_listQualificationTypesCmd.Flags().Bool("no-must-be-requestable", false, "Specifies that only Qualification types that a user can request through the Amazon Mechanical Turk web site, such as by taking a Qualification test, are returned as results of the search.")
		mturk_listQualificationTypesCmd.Flags().String("query", "", "A text query against all of the searchable attributes of Qualification types.")
		mturk_listQualificationTypesCmd.MarkFlagRequired("must-be-requestable")
		mturk_listQualificationTypesCmd.Flag("no-must-be-owned-by-caller").Hidden = true
		mturk_listQualificationTypesCmd.Flag("no-must-be-requestable").Hidden = true
		mturk_listQualificationTypesCmd.MarkFlagRequired("no-must-be-requestable")
	})
	mturkCmd.AddCommand(mturk_listQualificationTypesCmd)
}
