package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_regenerateSecurityTokenCmd = &cobra.Command{
	Use:   "regenerate-security-token",
	Short: "Regenerates the security token for a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_regenerateSecurityTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_regenerateSecurityTokenCmd).Standalone()

		chime_regenerateSecurityTokenCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_regenerateSecurityTokenCmd.Flags().String("bot-id", "", "The bot ID.")
		chime_regenerateSecurityTokenCmd.MarkFlagRequired("account-id")
		chime_regenerateSecurityTokenCmd.MarkFlagRequired("bot-id")
	})
	chimeCmd.AddCommand(chime_regenerateSecurityTokenCmd)
}
