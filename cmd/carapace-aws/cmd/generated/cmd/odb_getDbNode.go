package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getDbNodeCmd = &cobra.Command{
	Use:   "get-db-node",
	Short: "Returns information about the specified DB node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getDbNodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_getDbNodeCmd).Standalone()

		odb_getDbNodeCmd.Flags().String("cloud-vm-cluster-id", "", "The unique identifier of the VM cluster that contains the DB node.")
		odb_getDbNodeCmd.Flags().String("db-node-id", "", "The unique identifier of the DB node to retrieve information about.")
		odb_getDbNodeCmd.MarkFlagRequired("cloud-vm-cluster-id")
		odb_getDbNodeCmd.MarkFlagRequired("db-node-id")
	})
	odbCmd.AddCommand(odb_getDbNodeCmd)
}
