package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listPortalsCmd = &cobra.Command{
	Use:   "list-portals",
	Short: "Retrieves a list or web portals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listPortalsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_listPortalsCmd).Standalone()

		workspacesWeb_listPortalsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
		workspacesWeb_listPortalsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_listPortalsCmd)
}
