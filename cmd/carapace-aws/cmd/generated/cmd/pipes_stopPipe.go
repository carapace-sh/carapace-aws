package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_stopPipeCmd = &cobra.Command{
	Use:   "stop-pipe",
	Short: "Stop an existing pipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_stopPipeCmd).Standalone()

	pipes_stopPipeCmd.Flags().String("name", "", "The name of the pipe.")
	pipes_stopPipeCmd.MarkFlagRequired("name")
	pipesCmd.AddCommand(pipes_stopPipeCmd)
}
