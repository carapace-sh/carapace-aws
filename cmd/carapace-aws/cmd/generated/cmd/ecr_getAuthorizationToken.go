package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_getAuthorizationTokenCmd = &cobra.Command{
	Use:   "get-authorization-token",
	Short: "Retrieves an authorization token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_getAuthorizationTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_getAuthorizationTokenCmd).Standalone()

		ecr_getAuthorizationTokenCmd.Flags().String("registry-ids", "", "A list of Amazon Web Services account IDs that are associated with the registries for which to get AuthorizationData objects.")
	})
	ecrCmd.AddCommand(ecr_getAuthorizationTokenCmd)
}
