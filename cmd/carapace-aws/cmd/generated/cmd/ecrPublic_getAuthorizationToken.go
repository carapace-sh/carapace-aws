package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_getAuthorizationTokenCmd = &cobra.Command{
	Use:   "get-authorization-token",
	Short: "Retrieves an authorization token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_getAuthorizationTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrPublic_getAuthorizationTokenCmd).Standalone()

	})
	ecrPublicCmd.AddCommand(ecrPublic_getAuthorizationTokenCmd)
}
