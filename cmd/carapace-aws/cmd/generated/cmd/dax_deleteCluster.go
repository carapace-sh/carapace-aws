package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes a previously provisioned DAX cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_deleteClusterCmd).Standalone()

	dax_deleteClusterCmd.Flags().String("cluster-name", "", "The name of the cluster to be deleted.")
	dax_deleteClusterCmd.MarkFlagRequired("cluster-name")
	daxCmd.AddCommand(dax_deleteClusterCmd)
}
