package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listSdiSourcesCmd = &cobra.Command{
	Use:   "list-sdi-sources",
	Short: "List all the SdiSources in the AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listSdiSourcesCmd).Standalone()

	medialive_listSdiSourcesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	medialive_listSdiSourcesCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	medialiveCmd.AddCommand(medialive_listSdiSourcesCmd)
}
