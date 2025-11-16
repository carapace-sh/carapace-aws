package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_createAliasCmd = &cobra.Command{
	Use:   "create-alias",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nCreates an alias for a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_createAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_createAliasCmd).Standalone()

		gamelift_createAliasCmd.Flags().String("description", "", "A human-readable description of the alias.")
		gamelift_createAliasCmd.Flags().String("name", "", "A descriptive label that is associated with an alias.")
		gamelift_createAliasCmd.Flags().String("routing-strategy", "", "The routing configuration, including routing type and fleet target, for the alias.")
		gamelift_createAliasCmd.Flags().String("tags", "", "A list of labels to assign to the new alias resource.")
		gamelift_createAliasCmd.MarkFlagRequired("name")
		gamelift_createAliasCmd.MarkFlagRequired("routing-strategy")
	})
	gameliftCmd.AddCommand(gamelift_createAliasCmd)
}
