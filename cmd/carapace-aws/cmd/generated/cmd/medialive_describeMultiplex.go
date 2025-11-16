package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeMultiplexCmd = &cobra.Command{
	Use:   "describe-multiplex",
	Short: "Gets details about a multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeMultiplexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeMultiplexCmd).Standalone()

		medialive_describeMultiplexCmd.Flags().String("multiplex-id", "", "The ID of the multiplex.")
		medialive_describeMultiplexCmd.MarkFlagRequired("multiplex-id")
	})
	medialiveCmd.AddCommand(medialive_describeMultiplexCmd)
}
