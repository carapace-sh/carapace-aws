package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deleteIpsetCmd = &cobra.Command{
	Use:   "delete-ipset",
	Short: "Deletes the IPSet specified by the `ipSetId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deleteIpsetCmd).Standalone()

	guardduty_deleteIpsetCmd.Flags().String("detector-id", "", "The unique ID of the detector associated with the IPSet.")
	guardduty_deleteIpsetCmd.Flags().String("ip-set-id", "", "The unique ID of the IPSet to delete.")
	guardduty_deleteIpsetCmd.MarkFlagRequired("detector-id")
	guardduty_deleteIpsetCmd.MarkFlagRequired("ip-set-id")
	guarddutyCmd.AddCommand(guardduty_deleteIpsetCmd)
}
