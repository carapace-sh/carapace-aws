package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_describeHsmCmd = &cobra.Command{
	Use:   "describe-hsm",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_describeHsmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsm_describeHsmCmd).Standalone()

		cloudhsm_describeHsmCmd.Flags().String("hsm-arn", "", "The ARN of the HSM.")
		cloudhsm_describeHsmCmd.Flags().String("hsm-serial-number", "", "The serial number of the HSM.")
	})
	cloudhsmCmd.AddCommand(cloudhsm_describeHsmCmd)
}
