package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_createFargateProfileCmd = &cobra.Command{
	Use:   "create-fargate-profile",
	Short: "Creates an Fargate profile for your Amazon EKS cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_createFargateProfileCmd).Standalone()

	eks_createFargateProfileCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_createFargateProfileCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_createFargateProfileCmd.Flags().String("fargate-profile-name", "", "The name of the Fargate profile.")
	eks_createFargateProfileCmd.Flags().String("pod-execution-role-arn", "", "The Amazon Resource Name (ARN) of the `Pod` execution role to use for a `Pod` that matches the selectors in the Fargate profile.")
	eks_createFargateProfileCmd.Flags().String("selectors", "", "The selectors to match for a `Pod` to use this Fargate profile.")
	eks_createFargateProfileCmd.Flags().String("subnets", "", "The IDs of subnets to launch a `Pod` into.")
	eks_createFargateProfileCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
	eks_createFargateProfileCmd.MarkFlagRequired("cluster-name")
	eks_createFargateProfileCmd.MarkFlagRequired("fargate-profile-name")
	eks_createFargateProfileCmd.MarkFlagRequired("pod-execution-role-arn")
	eksCmd.AddCommand(eks_createFargateProfileCmd)
}
