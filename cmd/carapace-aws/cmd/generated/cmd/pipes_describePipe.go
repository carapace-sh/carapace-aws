package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_describePipeCmd = &cobra.Command{
	Use:   "describe-pipe",
	Short: "Get the information about an existing pipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_describePipeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pipes_describePipeCmd).Standalone()

		pipes_describePipeCmd.Flags().String("name", "", "The name of the pipe.")
		pipes_describePipeCmd.MarkFlagRequired("name")
	})
	pipesCmd.AddCommand(pipes_describePipeCmd)
}
