package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftCmd = &cobra.Command{
	Use:   "gamelift",
	Short: "Amazon GameLift Servers provides solutions for hosting session-based multiplayer game servers in the cloud, including tools for deploying, operating, and scaling game servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftCmd).Standalone()

	})
	rootCmd.AddCommand(gameliftCmd)
}
