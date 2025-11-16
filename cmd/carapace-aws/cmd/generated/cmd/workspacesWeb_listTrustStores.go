package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listTrustStoresCmd = &cobra.Command{
	Use:   "list-trust-stores",
	Short: "Retrieves a list of trust stores.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listTrustStoresCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_listTrustStoresCmd).Standalone()

		workspacesWeb_listTrustStoresCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
		workspacesWeb_listTrustStoresCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_listTrustStoresCmd)
}
