package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeAliasCmd = &cobra.Command{
	Use:   "describe-alias",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves properties for an alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeAliasCmd).Standalone()

		gamelift_describeAliasCmd.Flags().String("alias-id", "", "The unique identifier for the fleet alias that you want to retrieve.")
		gamelift_describeAliasCmd.MarkFlagRequired("alias-id")
	})
	gameliftCmd.AddCommand(gamelift_describeAliasCmd)
}
