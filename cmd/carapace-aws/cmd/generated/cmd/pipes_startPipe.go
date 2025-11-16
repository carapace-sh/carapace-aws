package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_startPipeCmd = &cobra.Command{
	Use:   "start-pipe",
	Short: "Start an existing pipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_startPipeCmd).Standalone()

	pipes_startPipeCmd.Flags().String("name", "", "The name of the pipe.")
	pipes_startPipeCmd.MarkFlagRequired("name")
	pipesCmd.AddCommand(pipes_startPipeCmd)
}
