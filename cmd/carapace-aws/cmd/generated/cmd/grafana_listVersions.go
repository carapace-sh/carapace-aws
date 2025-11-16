package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_listVersionsCmd = &cobra.Command{
	Use:   "list-versions",
	Short: "Lists available versions of Grafana.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_listVersionsCmd).Standalone()

	grafana_listVersionsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	grafana_listVersionsCmd.Flags().String("next-token", "", "The token to use when requesting the next set of results.")
	grafana_listVersionsCmd.Flags().String("workspace-id", "", "The ID of the workspace to list the available upgrade versions.")
	grafanaCmd.AddCommand(grafana_listVersionsCmd)
}
