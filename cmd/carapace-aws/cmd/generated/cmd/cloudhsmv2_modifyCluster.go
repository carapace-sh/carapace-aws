package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_modifyClusterCmd = &cobra.Command{
	Use:   "modify-cluster",
	Short: "Modifies CloudHSM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_modifyClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsmv2_modifyClusterCmd).Standalone()

		cloudhsmv2_modifyClusterCmd.Flags().String("backup-retention-policy", "", "A policy that defines how the service retains backups.")
		cloudhsmv2_modifyClusterCmd.Flags().String("cluster-id", "", "The identifier (ID) of the cluster that you want to modify.")
		cloudhsmv2_modifyClusterCmd.Flags().String("hsm-type", "", "The desired HSM type of the cluster.")
		cloudhsmv2_modifyClusterCmd.MarkFlagRequired("cluster-id")
	})
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_modifyClusterCmd)
}
