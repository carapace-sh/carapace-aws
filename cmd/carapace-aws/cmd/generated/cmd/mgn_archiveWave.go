package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_archiveWaveCmd = &cobra.Command{
	Use:   "archive-wave",
	Short: "Archive wave.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_archiveWaveCmd).Standalone()

	mgn_archiveWaveCmd.Flags().String("account-id", "", "Account ID.")
	mgn_archiveWaveCmd.Flags().String("wave-id", "", "Wave ID.")
	mgn_archiveWaveCmd.MarkFlagRequired("wave-id")
	mgnCmd.AddCommand(mgn_archiveWaveCmd)
}
