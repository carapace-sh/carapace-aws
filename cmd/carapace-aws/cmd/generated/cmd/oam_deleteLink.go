package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_deleteLinkCmd = &cobra.Command{
	Use:   "delete-link",
	Short: "Deletes a link between a monitoring account sink and a source account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_deleteLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(oam_deleteLinkCmd).Standalone()

		oam_deleteLinkCmd.Flags().String("identifier", "", "The ARN of the link to delete.")
		oam_deleteLinkCmd.MarkFlagRequired("identifier")
	})
	oamCmd.AddCommand(oam_deleteLinkCmd)
}
