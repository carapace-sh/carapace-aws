package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_listKeyspacesCmd = &cobra.Command{
	Use:   "list-keyspaces",
	Short: "The `ListKeyspaces` operation returns a list of keyspaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_listKeyspacesCmd).Standalone()

	keyspaces_listKeyspacesCmd.Flags().String("max-results", "", "The total number of keyspaces to return in the output.")
	keyspaces_listKeyspacesCmd.Flags().String("next-token", "", "The pagination token.")
	keyspacesCmd.AddCommand(keyspaces_listKeyspacesCmd)
}
