package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeMultiplexProgramCmd = &cobra.Command{
	Use:   "describe-multiplex-program",
	Short: "Get the details for a program in a multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeMultiplexProgramCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeMultiplexProgramCmd).Standalone()

		medialive_describeMultiplexProgramCmd.Flags().String("multiplex-id", "", "The ID of the multiplex that the program belongs to.")
		medialive_describeMultiplexProgramCmd.Flags().String("program-name", "", "The name of the program.")
		medialive_describeMultiplexProgramCmd.MarkFlagRequired("multiplex-id")
		medialive_describeMultiplexProgramCmd.MarkFlagRequired("program-name")
	})
	medialiveCmd.AddCommand(medialive_describeMultiplexProgramCmd)
}
