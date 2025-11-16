package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_listEnvironmentHostsCmd = &cobra.Command{
	Use:   "list-environment-hosts",
	Short: "List the hosts within an environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_listEnvironmentHostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evs_listEnvironmentHostsCmd).Standalone()

		evs_listEnvironmentHostsCmd.Flags().String("environment-id", "", "A unique ID for the environment.")
		evs_listEnvironmentHostsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		evs_listEnvironmentHostsCmd.Flags().String("next-token", "", "A unique pagination token for each page.")
		evs_listEnvironmentHostsCmd.MarkFlagRequired("environment-id")
	})
	evsCmd.AddCommand(evs_listEnvironmentHostsCmd)
}
