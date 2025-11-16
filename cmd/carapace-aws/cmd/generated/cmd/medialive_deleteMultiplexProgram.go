package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteMultiplexProgramCmd = &cobra.Command{
	Use:   "delete-multiplex-program",
	Short: "Delete a program from a multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteMultiplexProgramCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_deleteMultiplexProgramCmd).Standalone()

		medialive_deleteMultiplexProgramCmd.Flags().String("multiplex-id", "", "The ID of the multiplex that the program belongs to.")
		medialive_deleteMultiplexProgramCmd.Flags().String("program-name", "", "The multiplex program name.")
		medialive_deleteMultiplexProgramCmd.MarkFlagRequired("multiplex-id")
		medialive_deleteMultiplexProgramCmd.MarkFlagRequired("program-name")
	})
	medialiveCmd.AddCommand(medialive_deleteMultiplexProgramCmd)
}
