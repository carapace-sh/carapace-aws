package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listThreatIntelSetsCmd = &cobra.Command{
	Use:   "list-threat-intel-sets",
	Short: "Lists the ThreatIntelSets of the GuardDuty service specified by the detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listThreatIntelSetsCmd).Standalone()

	guardduty_listThreatIntelSetsCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with the threatIntelSet.")
	guardduty_listThreatIntelSetsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items that you want in the response.")
	guardduty_listThreatIntelSetsCmd.Flags().String("next-token", "", "You can use this parameter to paginate results in the response.")
	guardduty_listThreatIntelSetsCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_listThreatIntelSetsCmd)
}
