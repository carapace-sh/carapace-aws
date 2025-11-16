package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getBlacklistReportsCmd = &cobra.Command{
	Use:   "get-blacklist-reports",
	Short: "Retrieve a list of the blacklists that your dedicated IP addresses appear on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getBlacklistReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getBlacklistReportsCmd).Standalone()

		sesv2_getBlacklistReportsCmd.Flags().String("blacklist-item-names", "", "A list of IP addresses that you want to retrieve blacklist information about.")
		sesv2_getBlacklistReportsCmd.MarkFlagRequired("blacklist-item-names")
	})
	sesv2Cmd.AddCommand(sesv2_getBlacklistReportsCmd)
}
