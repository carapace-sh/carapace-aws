package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateWaveCmd = &cobra.Command{
	Use:   "update-wave",
	Short: "Update wave.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateWaveCmd).Standalone()

	mgn_updateWaveCmd.Flags().String("account-id", "", "Account ID.")
	mgn_updateWaveCmd.Flags().String("description", "", "Wave description.")
	mgn_updateWaveCmd.Flags().String("name", "", "Wave name.")
	mgn_updateWaveCmd.Flags().String("wave-id", "", "Wave ID.")
	mgn_updateWaveCmd.MarkFlagRequired("wave-id")
	mgnCmd.AddCommand(mgn_updateWaveCmd)
}
