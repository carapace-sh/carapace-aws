package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listIpsetsCmd = &cobra.Command{
	Use:   "list-ipsets",
	Short: "Lists the IPSets of the GuardDuty service specified by the detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listIpsetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_listIpsetsCmd).Standalone()

		guardduty_listIpsetsCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with IPSet.")
		guardduty_listIpsetsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
		guardduty_listIpsetsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
		guardduty_listIpsetsCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_listIpsetsCmd)
}
