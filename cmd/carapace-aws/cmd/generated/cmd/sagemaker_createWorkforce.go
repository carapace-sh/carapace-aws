package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createWorkforceCmd = &cobra.Command{
	Use:   "create-workforce",
	Short: "Use this operation to create a workforce.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createWorkforceCmd).Standalone()

	sagemaker_createWorkforceCmd.Flags().String("cognito-config", "", "Use this parameter to configure an Amazon Cognito private workforce.")
	sagemaker_createWorkforceCmd.Flags().String("ip-address-type", "", "Use this parameter to specify whether you want `IPv4` only or `dualstack` (`IPv4` and `IPv6`) to support your labeling workforce.")
	sagemaker_createWorkforceCmd.Flags().String("oidc-config", "", "Use this parameter to configure a private workforce using your own OIDC Identity Provider.")
	sagemaker_createWorkforceCmd.Flags().String("source-ip-config", "", "")
	sagemaker_createWorkforceCmd.Flags().String("tags", "", "An array of key-value pairs that contain metadata to help you categorize and organize our workforce.")
	sagemaker_createWorkforceCmd.Flags().String("workforce-name", "", "The name of the private workforce.")
	sagemaker_createWorkforceCmd.Flags().String("workforce-vpc-config", "", "Use this parameter to configure a workforce using VPC.")
	sagemaker_createWorkforceCmd.MarkFlagRequired("workforce-name")
	sagemakerCmd.AddCommand(sagemaker_createWorkforceCmd)
}
