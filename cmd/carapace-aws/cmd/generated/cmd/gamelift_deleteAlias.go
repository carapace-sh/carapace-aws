package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteAliasCmd = &cobra.Command{
	Use:   "delete-alias",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nDeletes an alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_deleteAliasCmd).Standalone()

		gamelift_deleteAliasCmd.Flags().String("alias-id", "", "A unique identifier of the alias that you want to delete.")
		gamelift_deleteAliasCmd.MarkFlagRequired("alias-id")
	})
	gameliftCmd.AddCommand(gamelift_deleteAliasCmd)
}
