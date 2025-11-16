package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deleteFilterCmd = &cobra.Command{
	Use:   "delete-filter",
	Short: "Deletes the filter specified by the filter name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deleteFilterCmd).Standalone()

	guardduty_deleteFilterCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with the filter.")
	guardduty_deleteFilterCmd.Flags().String("filter-name", "", "The name of the filter that you want to delete.")
	guardduty_deleteFilterCmd.MarkFlagRequired("detector-id")
	guardduty_deleteFilterCmd.MarkFlagRequired("filter-name")
	guarddutyCmd.AddCommand(guardduty_deleteFilterCmd)
}
