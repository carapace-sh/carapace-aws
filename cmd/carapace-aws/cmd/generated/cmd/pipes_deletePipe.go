package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_deletePipeCmd = &cobra.Command{
	Use:   "delete-pipe",
	Short: "Delete an existing pipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_deletePipeCmd).Standalone()

	pipes_deletePipeCmd.Flags().String("name", "", "The name of the pipe.")
	pipes_deletePipeCmd.MarkFlagRequired("name")
	pipesCmd.AddCommand(pipes_deletePipeCmd)
}
