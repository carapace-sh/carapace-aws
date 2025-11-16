package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getIpsetCmd = &cobra.Command{
	Use:   "get-ipset",
	Short: "Retrieves the IPSet specified by the `ipSetId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getIpsetCmd).Standalone()

	guardduty_getIpsetCmd.Flags().String("detector-id", "", "The unique ID of the detector that is associated with the IPSet.")
	guardduty_getIpsetCmd.Flags().String("ip-set-id", "", "The unique ID of the IPSet to retrieve.")
	guardduty_getIpsetCmd.MarkFlagRequired("detector-id")
	guardduty_getIpsetCmd.MarkFlagRequired("ip-set-id")
	guarddutyCmd.AddCommand(guardduty_getIpsetCmd)
}
