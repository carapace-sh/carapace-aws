package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_listAttachedLinksCmd = &cobra.Command{
	Use:   "list-attached-links",
	Short: "Returns a list of source account links that are linked to this monitoring account sink.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_listAttachedLinksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(oam_listAttachedLinksCmd).Standalone()

		oam_listAttachedLinksCmd.Flags().String("max-results", "", "Limits the number of returned links to the specified number.")
		oam_listAttachedLinksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		oam_listAttachedLinksCmd.Flags().String("sink-identifier", "", "The ARN of the sink that you want to retrieve links for.")
		oam_listAttachedLinksCmd.MarkFlagRequired("sink-identifier")
	})
	oamCmd.AddCommand(oam_listAttachedLinksCmd)
}
