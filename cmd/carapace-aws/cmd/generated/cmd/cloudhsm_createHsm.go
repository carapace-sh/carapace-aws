package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_createHsmCmd = &cobra.Command{
	Use:   "create-hsm",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_createHsmCmd).Standalone()

	cloudhsm_createHsmCmd.Flags().String("client-token", "", "A user-defined token to ensure idempotence.")
	cloudhsm_createHsmCmd.Flags().String("eni-ip", "", "The IP address to assign to the HSM's ENI.")
	cloudhsm_createHsmCmd.Flags().String("external-id", "", "The external ID from `IamRoleArn`, if present.")
	cloudhsm_createHsmCmd.Flags().String("iam-role-arn", "", "The ARN of an IAM role to enable the AWS CloudHSM service to allocate an ENI on your behalf.")
	cloudhsm_createHsmCmd.Flags().String("ssh-key", "", "The SSH public key to install on the HSM.")
	cloudhsm_createHsmCmd.Flags().String("subnet-id", "", "The identifier of the subnet in your VPC in which to place the HSM.")
	cloudhsm_createHsmCmd.Flags().String("subscription-type", "", "")
	cloudhsm_createHsmCmd.Flags().String("syslog-ip", "", "The IP address for the syslog monitoring server.")
	cloudhsm_createHsmCmd.MarkFlagRequired("iam-role-arn")
	cloudhsm_createHsmCmd.MarkFlagRequired("ssh-key")
	cloudhsm_createHsmCmd.MarkFlagRequired("subnet-id")
	cloudhsm_createHsmCmd.MarkFlagRequired("subscription-type")
	cloudhsmCmd.AddCommand(cloudhsm_createHsmCmd)
}
