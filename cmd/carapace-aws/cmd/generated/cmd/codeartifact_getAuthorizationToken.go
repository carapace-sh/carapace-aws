package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_getAuthorizationTokenCmd = &cobra.Command{
	Use:   "get-authorization-token",
	Short: "Generates a temporary authorization token for accessing repositories in the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_getAuthorizationTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_getAuthorizationTokenCmd).Standalone()

		codeartifact_getAuthorizationTokenCmd.Flags().String("domain", "", "The name of the domain that is in scope for the generated authorization token.")
		codeartifact_getAuthorizationTokenCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_getAuthorizationTokenCmd.Flags().String("duration-seconds", "", "The time, in seconds, that the generated authorization token is valid.")
		codeartifact_getAuthorizationTokenCmd.MarkFlagRequired("domain")
	})
	codeartifactCmd.AddCommand(codeartifact_getAuthorizationTokenCmd)
}
