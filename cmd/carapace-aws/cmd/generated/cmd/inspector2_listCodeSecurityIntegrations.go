package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCodeSecurityIntegrationsCmd = &cobra.Command{
	Use:   "list-code-security-integrations",
	Short: "Lists all code security integrations in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCodeSecurityIntegrationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listCodeSecurityIntegrationsCmd).Standalone()

		inspector2_listCodeSecurityIntegrationsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		inspector2_listCodeSecurityIntegrationsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	})
	inspector2Cmd.AddCommand(inspector2_listCodeSecurityIntegrationsCmd)
}
