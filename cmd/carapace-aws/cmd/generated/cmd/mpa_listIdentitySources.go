package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_listIdentitySourcesCmd = &cobra.Command{
	Use:   "list-identity-sources",
	Short: "Returns a list of identity sources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_listIdentitySourcesCmd).Standalone()

	mpa_listIdentitySourcesCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
	mpa_listIdentitySourcesCmd.Flags().String("next-token", "", "If present, indicates that more output is available than is included in the current response.")
	mpaCmd.AddCommand(mpa_listIdentitySourcesCmd)
}
