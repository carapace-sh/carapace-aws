package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_listPipesCmd = &cobra.Command{
	Use:   "list-pipes",
	Short: "Get the pipes associated with this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_listPipesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pipes_listPipesCmd).Standalone()

		pipes_listPipesCmd.Flags().String("current-state", "", "The state the pipe is in.")
		pipes_listPipesCmd.Flags().String("desired-state", "", "The state the pipe should be in.")
		pipes_listPipesCmd.Flags().String("limit", "", "The maximum number of pipes to include in the response.")
		pipes_listPipesCmd.Flags().String("name-prefix", "", "A value that will return a subset of the pipes associated with this account.")
		pipes_listPipesCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
		pipes_listPipesCmd.Flags().String("source-prefix", "", "The prefix matching the pipe source.")
		pipes_listPipesCmd.Flags().String("target-prefix", "", "The prefix matching the pipe target.")
	})
	pipesCmd.AddCommand(pipes_listPipesCmd)
}
