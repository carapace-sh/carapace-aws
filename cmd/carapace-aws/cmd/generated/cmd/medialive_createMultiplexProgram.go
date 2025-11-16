package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createMultiplexProgramCmd = &cobra.Command{
	Use:   "create-multiplex-program",
	Short: "Create a new program in the multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createMultiplexProgramCmd).Standalone()

	medialive_createMultiplexProgramCmd.Flags().String("multiplex-id", "", "ID of the multiplex where the program is to be created.")
	medialive_createMultiplexProgramCmd.Flags().String("multiplex-program-settings", "", "The settings for this multiplex program.")
	medialive_createMultiplexProgramCmd.Flags().String("program-name", "", "Name of multiplex program.")
	medialive_createMultiplexProgramCmd.Flags().String("request-id", "", "Unique request ID.")
	medialive_createMultiplexProgramCmd.MarkFlagRequired("multiplex-id")
	medialive_createMultiplexProgramCmd.MarkFlagRequired("multiplex-program-settings")
	medialive_createMultiplexProgramCmd.MarkFlagRequired("program-name")
	medialive_createMultiplexProgramCmd.MarkFlagRequired("request-id")
	medialiveCmd.AddCommand(medialive_createMultiplexProgramCmd)
}
