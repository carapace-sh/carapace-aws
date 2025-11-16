package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreamsCmd = &cobra.Command{
	Use:   "gameliftstreams",
	Short: "Amazon GameLift Streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreamsCmd).Standalone()

	})
	rootCmd.AddCommand(gameliftstreamsCmd)
}
