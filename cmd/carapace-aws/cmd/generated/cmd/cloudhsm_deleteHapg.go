package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_deleteHapgCmd = &cobra.Command{
	Use:   "delete-hapg",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_deleteHapgCmd).Standalone()

	cloudhsm_deleteHapgCmd.Flags().String("hapg-arn", "", "The ARN of the high-availability partition group to delete.")
	cloudhsm_deleteHapgCmd.MarkFlagRequired("hapg-arn")
	cloudhsmCmd.AddCommand(cloudhsm_deleteHapgCmd)
}
