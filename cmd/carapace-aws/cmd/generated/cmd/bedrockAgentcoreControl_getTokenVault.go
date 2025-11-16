package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getTokenVaultCmd = &cobra.Command{
	Use:   "get-token-vault",
	Short: "Retrieves information about a token vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getTokenVaultCmd).Standalone()

	bedrockAgentcoreControl_getTokenVaultCmd.Flags().String("token-vault-id", "", "The unique identifier of the token vault to retrieve.")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getTokenVaultCmd)
}
