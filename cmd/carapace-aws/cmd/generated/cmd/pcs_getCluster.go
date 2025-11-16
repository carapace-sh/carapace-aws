package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_getClusterCmd = &cobra.Command{
	Use:   "get-cluster",
	Short: "Returns detailed information about a running cluster in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_getClusterCmd).Standalone()

	pcs_getClusterCmd.Flags().String("cluster-identifier", "", "The name or ID of the cluster.")
	pcs_getClusterCmd.MarkFlagRequired("cluster-identifier")
	pcsCmd.AddCommand(pcs_getClusterCmd)
}
