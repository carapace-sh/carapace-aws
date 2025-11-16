package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listIdentityProvidersCmd = &cobra.Command{
	Use:   "list-identity-providers",
	Short: "Retrieves a list of identity providers for a specific web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listIdentityProvidersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_listIdentityProvidersCmd).Standalone()

		workspacesWeb_listIdentityProvidersCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
		workspacesWeb_listIdentityProvidersCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
		workspacesWeb_listIdentityProvidersCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_listIdentityProvidersCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_listIdentityProvidersCmd)
}
