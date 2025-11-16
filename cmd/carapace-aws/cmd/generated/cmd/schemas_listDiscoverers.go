package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_listDiscoverersCmd = &cobra.Command{
	Use:   "list-discoverers",
	Short: "List the discoverers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_listDiscoverersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_listDiscoverersCmd).Standalone()

		schemas_listDiscoverersCmd.Flags().String("discoverer-id-prefix", "", "Specifying this limits the results to only those discoverer IDs that start with the specified prefix.")
		schemas_listDiscoverersCmd.Flags().String("limit", "", "")
		schemas_listDiscoverersCmd.Flags().String("next-token", "", "The token that specifies the next page of results to return.")
		schemas_listDiscoverersCmd.Flags().String("source-arn-prefix", "", "Specifying this limits the results to only those ARNs that start with the specified prefix.")
	})
	schemasCmd.AddCommand(schemas_listDiscoverersCmd)
}
