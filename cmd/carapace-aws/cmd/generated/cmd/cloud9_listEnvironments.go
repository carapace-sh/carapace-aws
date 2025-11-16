package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "Gets a list of Cloud9 development environment identifiers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_listEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9_listEnvironmentsCmd).Standalone()

		cloud9_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of environments to get identifiers for.")
		cloud9_listEnvironmentsCmd.Flags().String("next-token", "", "During a previous call, if there are more than 25 items in the list, only the first 25 items are returned, along with a unique string called a *next token*.")
	})
	cloud9Cmd.AddCommand(cloud9_listEnvironmentsCmd)
}
