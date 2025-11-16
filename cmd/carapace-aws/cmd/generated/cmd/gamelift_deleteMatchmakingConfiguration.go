package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteMatchmakingConfigurationCmd = &cobra.Command{
	Use:   "delete-matchmaking-configuration",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nPermanently removes a FlexMatch matchmaking configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteMatchmakingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_deleteMatchmakingConfigurationCmd).Standalone()

		gamelift_deleteMatchmakingConfigurationCmd.Flags().String("name", "", "A unique identifier for the matchmaking configuration.")
		gamelift_deleteMatchmakingConfigurationCmd.MarkFlagRequired("name")
	})
	gameliftCmd.AddCommand(gamelift_deleteMatchmakingConfigurationCmd)
}
