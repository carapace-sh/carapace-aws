package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_disassociateDrtroleCmd = &cobra.Command{
	Use:   "disassociate-drtrole",
	Short: "Removes the Shield Response Team's (SRT) access to your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_disassociateDrtroleCmd).Standalone()

	shieldCmd.AddCommand(shield_disassociateDrtroleCmd)
}
