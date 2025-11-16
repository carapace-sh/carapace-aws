package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createMultiplexCmd = &cobra.Command{
	Use:   "create-multiplex",
	Short: "Create a new multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createMultiplexCmd).Standalone()

	medialive_createMultiplexCmd.Flags().String("availability-zones", "", "A list of availability zones for the multiplex.")
	medialive_createMultiplexCmd.Flags().String("multiplex-settings", "", "Configuration for a multiplex event.")
	medialive_createMultiplexCmd.Flags().String("name", "", "Name of multiplex.")
	medialive_createMultiplexCmd.Flags().String("request-id", "", "Unique request ID.")
	medialive_createMultiplexCmd.Flags().String("tags", "", "A collection of key-value pairs.")
	medialive_createMultiplexCmd.MarkFlagRequired("availability-zones")
	medialive_createMultiplexCmd.MarkFlagRequired("multiplex-settings")
	medialive_createMultiplexCmd.MarkFlagRequired("name")
	medialive_createMultiplexCmd.MarkFlagRequired("request-id")
	medialiveCmd.AddCommand(medialive_createMultiplexCmd)
}
