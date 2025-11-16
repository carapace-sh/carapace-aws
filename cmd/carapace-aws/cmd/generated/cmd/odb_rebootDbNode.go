package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_rebootDbNodeCmd = &cobra.Command{
	Use:   "reboot-db-node",
	Short: "Reboots the specified DB node in a VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_rebootDbNodeCmd).Standalone()

	odb_rebootDbNodeCmd.Flags().String("cloud-vm-cluster-id", "", "The unique identifier of the VM cluster that contains the DB node to reboot.")
	odb_rebootDbNodeCmd.Flags().String("db-node-id", "", "The unique identifier of the DB node to reboot.")
	odb_rebootDbNodeCmd.MarkFlagRequired("cloud-vm-cluster-id")
	odb_rebootDbNodeCmd.MarkFlagRequired("db-node-id")
	odbCmd.AddCommand(odb_rebootDbNodeCmd)
}
