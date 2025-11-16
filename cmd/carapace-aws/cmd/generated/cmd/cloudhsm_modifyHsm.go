package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_modifyHsmCmd = &cobra.Command{
	Use:   "modify-hsm",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_modifyHsmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsm_modifyHsmCmd).Standalone()

		cloudhsm_modifyHsmCmd.Flags().String("eni-ip", "", "The new IP address for the elastic network interface (ENI) attached to the HSM.")
		cloudhsm_modifyHsmCmd.Flags().String("external-id", "", "The new external ID.")
		cloudhsm_modifyHsmCmd.Flags().String("hsm-arn", "", "The ARN of the HSM to modify.")
		cloudhsm_modifyHsmCmd.Flags().String("iam-role-arn", "", "The new IAM role ARN.")
		cloudhsm_modifyHsmCmd.Flags().String("subnet-id", "", "The new identifier of the subnet that the HSM is in.")
		cloudhsm_modifyHsmCmd.Flags().String("syslog-ip", "", "The new IP address for the syslog monitoring server.")
		cloudhsm_modifyHsmCmd.MarkFlagRequired("hsm-arn")
	})
	cloudhsmCmd.AddCommand(cloudhsm_modifyHsmCmd)
}
