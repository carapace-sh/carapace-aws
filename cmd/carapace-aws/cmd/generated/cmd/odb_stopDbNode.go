package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_stopDbNodeCmd = &cobra.Command{
	Use:   "stop-db-node",
	Short: "Stops the specified DB node in a VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_stopDbNodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_stopDbNodeCmd).Standalone()

		odb_stopDbNodeCmd.Flags().String("cloud-vm-cluster-id", "", "The unique identifier of the VM cluster that contains the DB node to stop.")
		odb_stopDbNodeCmd.Flags().String("db-node-id", "", "The unique identifier of the DB node to stop.")
		odb_stopDbNodeCmd.MarkFlagRequired("cloud-vm-cluster-id")
		odb_stopDbNodeCmd.MarkFlagRequired("db-node-id")
	})
	odbCmd.AddCommand(odb_stopDbNodeCmd)
}
