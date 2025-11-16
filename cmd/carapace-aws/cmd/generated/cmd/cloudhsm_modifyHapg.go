package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_modifyHapgCmd = &cobra.Command{
	Use:   "modify-hapg",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_modifyHapgCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsm_modifyHapgCmd).Standalone()

		cloudhsm_modifyHapgCmd.Flags().String("hapg-arn", "", "The ARN of the high-availability partition group to modify.")
		cloudhsm_modifyHapgCmd.Flags().String("label", "", "The new label for the high-availability partition group.")
		cloudhsm_modifyHapgCmd.Flags().String("partition-serial-list", "", "The list of partition serial numbers to make members of the high-availability partition group.")
		cloudhsm_modifyHapgCmd.MarkFlagRequired("hapg-arn")
	})
	cloudhsmCmd.AddCommand(cloudhsm_modifyHapgCmd)
}
