package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_listLinksCmd = &cobra.Command{
	Use:   "list-links",
	Short: "Use this operation in a source account to return a list of links to monitoring account sinks that this source account has.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_listLinksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(oam_listLinksCmd).Standalone()

		oam_listLinksCmd.Flags().String("max-results", "", "Limits the number of returned links to the specified number.")
		oam_listLinksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	oamCmd.AddCommand(oam_listLinksCmd)
}
