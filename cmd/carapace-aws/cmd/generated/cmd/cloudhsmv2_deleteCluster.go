package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes the specified CloudHSM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_deleteClusterCmd).Standalone()

	cloudhsmv2_deleteClusterCmd.Flags().String("cluster-id", "", "The identifier (ID) of the cluster that you are deleting.")
	cloudhsmv2_deleteClusterCmd.MarkFlagRequired("cluster-id")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_deleteClusterCmd)
}
