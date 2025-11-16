package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateMultiplexCmd = &cobra.Command{
	Use:   "update-multiplex",
	Short: "Updates a multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateMultiplexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_updateMultiplexCmd).Standalone()

		medialive_updateMultiplexCmd.Flags().String("multiplex-id", "", "ID of the multiplex to update.")
		medialive_updateMultiplexCmd.Flags().String("multiplex-settings", "", "The new settings for a multiplex.")
		medialive_updateMultiplexCmd.Flags().String("name", "", "Name of the multiplex.")
		medialive_updateMultiplexCmd.Flags().String("packet-identifiers-mapping", "", "")
		medialive_updateMultiplexCmd.MarkFlagRequired("multiplex-id")
	})
	medialiveCmd.AddCommand(medialive_updateMultiplexCmd)
}
