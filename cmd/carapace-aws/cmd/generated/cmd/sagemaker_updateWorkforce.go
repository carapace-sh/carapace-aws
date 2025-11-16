package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateWorkforceCmd = &cobra.Command{
	Use:   "update-workforce",
	Short: "Use this operation to update your workforce.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateWorkforceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateWorkforceCmd).Standalone()

		sagemaker_updateWorkforceCmd.Flags().String("ip-address-type", "", "Use this parameter to specify whether you want `IPv4` only or `dualstack` (`IPv4` and `IPv6`) to support your labeling workforce.")
		sagemaker_updateWorkforceCmd.Flags().String("oidc-config", "", "Use this parameter to update your OIDC Identity Provider (IdP) configuration for a workforce made using your own IdP.")
		sagemaker_updateWorkforceCmd.Flags().String("source-ip-config", "", "A list of one to ten worker IP address ranges ([CIDRs](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)) that can be used to access tasks assigned to this workforce.")
		sagemaker_updateWorkforceCmd.Flags().String("workforce-name", "", "The name of the private workforce that you want to update.")
		sagemaker_updateWorkforceCmd.Flags().String("workforce-vpc-config", "", "Use this parameter to update your VPC configuration for a workforce.")
		sagemaker_updateWorkforceCmd.MarkFlagRequired("workforce-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateWorkforceCmd)
}
