package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_rebootNodeCmd = &cobra.Command{
	Use:   "reboot-node",
	Short: "Reboots a single node of a DAX cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_rebootNodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_rebootNodeCmd).Standalone()

		dax_rebootNodeCmd.Flags().String("cluster-name", "", "The name of the DAX cluster containing the node to be rebooted.")
		dax_rebootNodeCmd.Flags().String("node-id", "", "The system-assigned ID of the node to be rebooted.")
		dax_rebootNodeCmd.MarkFlagRequired("cluster-name")
		dax_rebootNodeCmd.MarkFlagRequired("node-id")
	})
	daxCmd.AddCommand(dax_rebootNodeCmd)
}
