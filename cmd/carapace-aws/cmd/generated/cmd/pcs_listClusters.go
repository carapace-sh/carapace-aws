package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Returns a list of running clusters in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_listClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_listClustersCmd).Standalone()

		pcs_listClustersCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		pcs_listClustersCmd.Flags().String("next-token", "", "The value of `nextToken` is a unique pagination token for each page of results returned.")
	})
	pcsCmd.AddCommand(pcs_listClustersCmd)
}
