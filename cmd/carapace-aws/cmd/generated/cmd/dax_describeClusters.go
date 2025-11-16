package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_describeClustersCmd = &cobra.Command{
	Use:   "describe-clusters",
	Short: "Returns information about all provisioned DAX clusters if no cluster identifier is specified, or about a specific DAX cluster if a cluster identifier is supplied.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_describeClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_describeClustersCmd).Standalone()

		dax_describeClustersCmd.Flags().String("cluster-names", "", "The names of the DAX clusters being described.")
		dax_describeClustersCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		dax_describeClustersCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	})
	daxCmd.AddCommand(dax_describeClustersCmd)
}
