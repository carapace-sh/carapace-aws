package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_listCrlsCmd = &cobra.Command{
	Use:   "list-crls",
	Short: "Lists all certificate revocation lists (CRL) in the authenticated account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_listCrlsCmd).Standalone()

	rolesanywhere_listCrlsCmd.Flags().String("next-token", "", "A token that indicates where the output should continue from, if a previous request did not show all results.")
	rolesanywhere_listCrlsCmd.Flags().String("page-size", "", "The number of resources in the paginated list.")
	rolesanywhereCmd.AddCommand(rolesanywhere_listCrlsCmd)
}
