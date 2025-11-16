package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteDedicatedIpPoolCmd = &cobra.Command{
	Use:   "delete-dedicated-ip-pool",
	Short: "Delete a dedicated IP pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteDedicatedIpPoolCmd).Standalone()

	sesv2_deleteDedicatedIpPoolCmd.Flags().String("pool-name", "", "The name of the dedicated IP pool that you want to delete.")
	sesv2_deleteDedicatedIpPoolCmd.MarkFlagRequired("pool-name")
	sesv2Cmd.AddCommand(sesv2_deleteDedicatedIpPoolCmd)
}
