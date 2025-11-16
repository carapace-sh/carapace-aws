package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_getSigningPlatformCmd = &cobra.Command{
	Use:   "get-signing-platform",
	Short: "Returns information on a specific signing platform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_getSigningPlatformCmd).Standalone()

	signer_getSigningPlatformCmd.Flags().String("platform-id", "", "The ID of the target signing platform.")
	signer_getSigningPlatformCmd.MarkFlagRequired("platform-id")
	signerCmd.AddCommand(signer_getSigningPlatformCmd)
}
