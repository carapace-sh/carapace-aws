package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_listEnvironmentVlansCmd = &cobra.Command{
	Use:   "list-environment-vlans",
	Short: "Lists environment VLANs that are associated with the specified environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_listEnvironmentVlansCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evs_listEnvironmentVlansCmd).Standalone()

		evs_listEnvironmentVlansCmd.Flags().String("environment-id", "", "A unique ID for the environment.")
		evs_listEnvironmentVlansCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		evs_listEnvironmentVlansCmd.Flags().String("next-token", "", "A unique pagination token for each page.")
		evs_listEnvironmentVlansCmd.MarkFlagRequired("environment-id")
	})
	evsCmd.AddCommand(evs_listEnvironmentVlansCmd)
}
