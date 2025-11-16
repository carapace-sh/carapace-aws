package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listCoverageCmd = &cobra.Command{
	Use:   "list-coverage",
	Short: "Lists coverage details for your GuardDuty account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listCoverageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_listCoverageCmd).Standalone()

		guardduty_listCoverageCmd.Flags().String("detector-id", "", "The unique ID of the detector whose coverage details you want to retrieve.")
		guardduty_listCoverageCmd.Flags().String("filter-criteria", "", "Represents the criteria used to filter the coverage details.")
		guardduty_listCoverageCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		guardduty_listCoverageCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
		guardduty_listCoverageCmd.Flags().String("sort-criteria", "", "Represents the criteria used to sort the coverage details.")
		guardduty_listCoverageCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_listCoverageCmd)
}
