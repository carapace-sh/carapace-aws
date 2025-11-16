package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getFilterCmd = &cobra.Command{
	Use:   "get-filter",
	Short: "Returns the details of the filter specified by the filter name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getFilterCmd).Standalone()

		guardduty_getFilterCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with this filter.")
		guardduty_getFilterCmd.Flags().String("filter-name", "", "The name of the filter you want to get.")
		guardduty_getFilterCmd.MarkFlagRequired("detector-id")
		guardduty_getFilterCmd.MarkFlagRequired("filter-name")
	})
	guarddutyCmd.AddCommand(guardduty_getFilterCmd)
}
