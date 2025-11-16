package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_startDbNodeCmd = &cobra.Command{
	Use:   "start-db-node",
	Short: "Starts the specified DB node in a VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_startDbNodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_startDbNodeCmd).Standalone()

		odb_startDbNodeCmd.Flags().String("cloud-vm-cluster-id", "", "The unique identifier of the VM cluster that contains the DB node to start.")
		odb_startDbNodeCmd.Flags().String("db-node-id", "", "The unique identifier of the DB node to start.")
		odb_startDbNodeCmd.MarkFlagRequired("cloud-vm-cluster-id")
		odb_startDbNodeCmd.MarkFlagRequired("db-node-id")
	})
	odbCmd.AddCommand(odb_startDbNodeCmd)
}
