package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listAccessTokensCmd = &cobra.Command{
	Use:   "list-access-tokens",
	Short: "Lists all personal access tokens (PATs) associated with the user who calls the API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listAccessTokensCmd).Standalone()

	codecatalyst_listAccessTokensCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
	codecatalyst_listAccessTokensCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
	codecatalystCmd.AddCommand(codecatalyst_listAccessTokensCmd)
}
