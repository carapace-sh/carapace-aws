package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_stopMultiplexCmd = &cobra.Command{
	Use:   "stop-multiplex",
	Short: "Stops a running multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_stopMultiplexCmd).Standalone()

	medialive_stopMultiplexCmd.Flags().String("multiplex-id", "", "The ID of the multiplex.")
	medialive_stopMultiplexCmd.MarkFlagRequired("multiplex-id")
	medialiveCmd.AddCommand(medialive_stopMultiplexCmd)
}
