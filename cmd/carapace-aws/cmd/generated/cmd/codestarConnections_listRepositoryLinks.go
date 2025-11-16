package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_listRepositoryLinksCmd = &cobra.Command{
	Use:   "list-repository-links",
	Short: "Lists the repository links created for connections in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_listRepositoryLinksCmd).Standalone()

	codestarConnections_listRepositoryLinksCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codestarConnections_listRepositoryLinksCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codestarConnectionsCmd.AddCommand(codestarConnections_listRepositoryLinksCmd)
}
