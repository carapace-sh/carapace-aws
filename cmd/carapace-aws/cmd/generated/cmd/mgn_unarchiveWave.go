package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_unarchiveWaveCmd = &cobra.Command{
	Use:   "unarchive-wave",
	Short: "Unarchive wave.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_unarchiveWaveCmd).Standalone()

	mgn_unarchiveWaveCmd.Flags().String("account-id", "", "Account ID.")
	mgn_unarchiveWaveCmd.Flags().String("wave-id", "", "Wave ID.")
	mgn_unarchiveWaveCmd.MarkFlagRequired("wave-id")
	mgnCmd.AddCommand(mgn_unarchiveWaveCmd)
}
