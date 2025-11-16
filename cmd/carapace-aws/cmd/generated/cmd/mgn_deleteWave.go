package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_deleteWaveCmd = &cobra.Command{
	Use:   "delete-wave",
	Short: "Delete wave.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_deleteWaveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_deleteWaveCmd).Standalone()

		mgn_deleteWaveCmd.Flags().String("account-id", "", "Account ID.")
		mgn_deleteWaveCmd.Flags().String("wave-id", "", "Wave ID.")
		mgn_deleteWaveCmd.MarkFlagRequired("wave-id")
	})
	mgnCmd.AddCommand(mgn_deleteWaveCmd)
}
