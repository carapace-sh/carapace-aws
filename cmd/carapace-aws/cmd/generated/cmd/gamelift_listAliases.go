package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listAliasesCmd = &cobra.Command{
	Use:   "list-aliases",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves all aliases for this Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listAliasesCmd).Standalone()

		gamelift_listAliasesCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listAliasesCmd.Flags().String("name", "", "A descriptive label that is associated with an alias.")
		gamelift_listAliasesCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
		gamelift_listAliasesCmd.Flags().String("routing-strategy-type", "", "The routing type to filter results on.")
	})
	gameliftCmd.AddCommand(gamelift_listAliasesCmd)
}
