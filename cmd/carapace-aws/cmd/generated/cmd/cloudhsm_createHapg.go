package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_createHapgCmd = &cobra.Command{
	Use:   "create-hapg",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_createHapgCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsm_createHapgCmd).Standalone()

		cloudhsm_createHapgCmd.Flags().String("label", "", "The label of the new high-availability partition group.")
		cloudhsm_createHapgCmd.MarkFlagRequired("label")
	})
	cloudhsmCmd.AddCommand(cloudhsm_createHapgCmd)
}
