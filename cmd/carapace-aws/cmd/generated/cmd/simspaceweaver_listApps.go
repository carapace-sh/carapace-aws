package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "Lists all custom apps or service apps for the given simulation and domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_listAppsCmd).Standalone()

	simspaceweaver_listAppsCmd.Flags().String("domain", "", "The name of the domain that you want to list apps for.")
	simspaceweaver_listAppsCmd.Flags().String("max-results", "", "The maximum number of apps to list.")
	simspaceweaver_listAppsCmd.Flags().String("next-token", "", "If SimSpace Weaver returns `nextToken`, then there are more results available.")
	simspaceweaver_listAppsCmd.Flags().String("simulation", "", "The name of the simulation that you want to list apps for.")
	simspaceweaver_listAppsCmd.MarkFlagRequired("simulation")
	simspaceweaverCmd.AddCommand(simspaceweaver_listAppsCmd)
}
