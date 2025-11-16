package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_createWaveCmd = &cobra.Command{
	Use:   "create-wave",
	Short: "Create wave.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_createWaveCmd).Standalone()

	mgn_createWaveCmd.Flags().String("account-id", "", "Account ID.")
	mgn_createWaveCmd.Flags().String("description", "", "Wave description.")
	mgn_createWaveCmd.Flags().String("name", "", "Wave name.")
	mgn_createWaveCmd.Flags().String("tags", "", "Wave tags.")
	mgn_createWaveCmd.MarkFlagRequired("name")
	mgnCmd.AddCommand(mgn_createWaveCmd)
}
