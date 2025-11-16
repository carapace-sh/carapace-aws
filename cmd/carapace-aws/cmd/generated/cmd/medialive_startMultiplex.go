package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_startMultiplexCmd = &cobra.Command{
	Use:   "start-multiplex",
	Short: "Start (run) the multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_startMultiplexCmd).Standalone()

	medialive_startMultiplexCmd.Flags().String("multiplex-id", "", "The ID of the multiplex.")
	medialive_startMultiplexCmd.MarkFlagRequired("multiplex-id")
	medialiveCmd.AddCommand(medialive_startMultiplexCmd)
}
