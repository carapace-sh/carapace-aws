package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_listSinksCmd = &cobra.Command{
	Use:   "list-sinks",
	Short: "Use this operation in a monitoring account to return the list of sinks created in that account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_listSinksCmd).Standalone()

	oam_listSinksCmd.Flags().String("max-results", "", "Limits the number of returned links to the specified number.")
	oam_listSinksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	oamCmd.AddCommand(oam_listSinksCmd)
}
