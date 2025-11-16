package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateAliasCmd = &cobra.Command{
	Use:   "update-alias",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nUpdates properties for an alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateAliasCmd).Standalone()

	gamelift_updateAliasCmd.Flags().String("alias-id", "", "A unique identifier for the alias that you want to update.")
	gamelift_updateAliasCmd.Flags().String("description", "", "A human-readable description of the alias.")
	gamelift_updateAliasCmd.Flags().String("name", "", "A descriptive label that is associated with an alias.")
	gamelift_updateAliasCmd.Flags().String("routing-strategy", "", "The routing configuration, including routing type and fleet target, for the alias.")
	gamelift_updateAliasCmd.MarkFlagRequired("alias-id")
	gameliftCmd.AddCommand(gamelift_updateAliasCmd)
}
