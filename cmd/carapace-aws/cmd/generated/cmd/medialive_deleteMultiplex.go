package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteMultiplexCmd = &cobra.Command{
	Use:   "delete-multiplex",
	Short: "Delete a multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteMultiplexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_deleteMultiplexCmd).Standalone()

		medialive_deleteMultiplexCmd.Flags().String("multiplex-id", "", "The ID of the multiplex.")
		medialive_deleteMultiplexCmd.MarkFlagRequired("multiplex-id")
	})
	medialiveCmd.AddCommand(medialive_deleteMultiplexCmd)
}
