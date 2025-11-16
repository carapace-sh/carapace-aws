package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listTapesCmd = &cobra.Command{
	Use:   "list-tapes",
	Short: "Lists virtual tapes in your virtual tape library (VTL) and your virtual tape shelf (VTS).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listTapesCmd).Standalone()

	storagegateway_listTapesCmd.Flags().String("limit", "", "An optional number limit for the tapes in the list returned by this call.")
	storagegateway_listTapesCmd.Flags().String("marker", "", "A string that indicates the position at which to begin the returned list of tapes.")
	storagegateway_listTapesCmd.Flags().String("tape-arns", "", "")
	storagegatewayCmd.AddCommand(storagegateway_listTapesCmd)
}
