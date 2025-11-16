package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_deleteHsmCmd = &cobra.Command{
	Use:   "delete-hsm",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_deleteHsmCmd).Standalone()

	cloudhsm_deleteHsmCmd.Flags().String("hsm-arn", "", "The ARN of the HSM to delete.")
	cloudhsm_deleteHsmCmd.MarkFlagRequired("hsm-arn")
	cloudhsmCmd.AddCommand(cloudhsm_deleteHsmCmd)
}
