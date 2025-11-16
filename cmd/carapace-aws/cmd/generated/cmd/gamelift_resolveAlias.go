package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_resolveAliasCmd = &cobra.Command{
	Use:   "resolve-alias",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nAttempts to retrieve a fleet ID that is associated with an alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_resolveAliasCmd).Standalone()

	gamelift_resolveAliasCmd.Flags().String("alias-id", "", "The unique identifier of the alias that you want to retrieve a fleet ID for.")
	gamelift_resolveAliasCmd.MarkFlagRequired("alias-id")
	gameliftCmd.AddCommand(gamelift_resolveAliasCmd)
}
