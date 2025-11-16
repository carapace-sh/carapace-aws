package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_getClusterCmd = &cobra.Command{
	Use:   "get-cluster",
	Short: "Retrieves information about a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_getClusterCmd).Standalone()

	dsql_getClusterCmd.Flags().String("identifier", "", "The ID of the cluster to retrieve.")
	dsql_getClusterCmd.MarkFlagRequired("identifier")
	dsqlCmd.AddCommand(dsql_getClusterCmd)
}
