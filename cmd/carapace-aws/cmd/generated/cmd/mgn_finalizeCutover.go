package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_finalizeCutoverCmd = &cobra.Command{
	Use:   "finalize-cutover",
	Short: "Finalizes the cutover immediately for specific Source Servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_finalizeCutoverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_finalizeCutoverCmd).Standalone()

		mgn_finalizeCutoverCmd.Flags().String("account-id", "", "Request to finalize Cutover by Source Account ID.")
		mgn_finalizeCutoverCmd.Flags().String("source-server-id", "", "Request to finalize Cutover by Source Server ID.")
		mgn_finalizeCutoverCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_finalizeCutoverCmd)
}
