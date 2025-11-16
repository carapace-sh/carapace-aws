package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listAutonomousVirtualMachinesCmd = &cobra.Command{
	Use:   "list-autonomous-virtual-machines",
	Short: "Lists all Autonomous VMs in an Autonomous VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listAutonomousVirtualMachinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_listAutonomousVirtualMachinesCmd).Standalone()

		odb_listAutonomousVirtualMachinesCmd.Flags().String("cloud-autonomous-vm-cluster-id", "", "The unique identifier of the Autonomous VM cluster whose virtual machines you're listing.")
		odb_listAutonomousVirtualMachinesCmd.Flags().String("max-results", "", "The maximum number of items to return per page.")
		odb_listAutonomousVirtualMachinesCmd.Flags().String("next-token", "", "The pagination token to continue listing from.")
		odb_listAutonomousVirtualMachinesCmd.MarkFlagRequired("cloud-autonomous-vm-cluster-id")
	})
	odbCmd.AddCommand(odb_listAutonomousVirtualMachinesCmd)
}
