package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listMultiplexProgramsCmd = &cobra.Command{
	Use:   "list-multiplex-programs",
	Short: "List the programs that currently exist for a specific multiplex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listMultiplexProgramsCmd).Standalone()

	medialive_listMultiplexProgramsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	medialive_listMultiplexProgramsCmd.Flags().String("multiplex-id", "", "The ID of the multiplex that the programs belong to.")
	medialive_listMultiplexProgramsCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	medialive_listMultiplexProgramsCmd.MarkFlagRequired("multiplex-id")
	medialiveCmd.AddCommand(medialive_listMultiplexProgramsCmd)
}
