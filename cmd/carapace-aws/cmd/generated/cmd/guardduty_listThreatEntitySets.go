package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listThreatEntitySetsCmd = &cobra.Command{
	Use:   "list-threat-entity-sets",
	Short: "Lists the threat entity sets associated with the specified GuardDuty detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listThreatEntitySetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_listThreatEntitySetsCmd).Standalone()

		guardduty_listThreatEntitySetsCmd.Flags().String("detector-id", "", "The unique ID of the GuardDuty detector that is associated with this threat entity set.")
		guardduty_listThreatEntitySetsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
		guardduty_listThreatEntitySetsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
		guardduty_listThreatEntitySetsCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_listThreatEntitySetsCmd)
}
