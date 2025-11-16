package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateMultiplexProgramCmd = &cobra.Command{
	Use:   "update-multiplex-program",
	Short: "Update a program in a multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateMultiplexProgramCmd).Standalone()

	medialive_updateMultiplexProgramCmd.Flags().String("multiplex-id", "", "The ID of the multiplex of the program to update.")
	medialive_updateMultiplexProgramCmd.Flags().String("multiplex-program-settings", "", "The new settings for a multiplex program.")
	medialive_updateMultiplexProgramCmd.Flags().String("program-name", "", "The name of the program to update.")
	medialive_updateMultiplexProgramCmd.MarkFlagRequired("multiplex-id")
	medialive_updateMultiplexProgramCmd.MarkFlagRequired("program-name")
	medialiveCmd.AddCommand(medialive_updateMultiplexProgramCmd)
}
