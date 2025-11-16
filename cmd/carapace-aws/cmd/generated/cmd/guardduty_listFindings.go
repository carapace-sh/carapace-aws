package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listFindingsCmd = &cobra.Command{
	Use:   "list-findings",
	Short: "Lists GuardDuty findings for the specified detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listFindingsCmd).Standalone()

	guardduty_listFindingsCmd.Flags().String("detector-id", "", "The ID of the detector that specifies the GuardDuty service whose findings you want to list.")
	guardduty_listFindingsCmd.Flags().String("finding-criteria", "", "Represents the criteria used for querying findings.")
	guardduty_listFindingsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
	guardduty_listFindingsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	guardduty_listFindingsCmd.Flags().String("sort-criteria", "", "Represents the criteria used for sorting findings.")
	guardduty_listFindingsCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_listFindingsCmd)
}
