package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createActivationCmd = &cobra.Command{
	Use:   "create-activation",
	Short: "Generates an activation code and activation ID you can use to register your on-premises servers, edge devices, or virtual machine (VM) with Amazon Web Services Systems Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createActivationCmd).Standalone()

	ssm_createActivationCmd.Flags().String("default-instance-name", "", "The name of the registered, managed node as it will appear in the Amazon Web Services Systems Manager console or when you use the Amazon Web Services command line tools to list Systems Manager resources.")
	ssm_createActivationCmd.Flags().String("description", "", "A user-defined description of the resource that you want to register with Systems Manager.")
	ssm_createActivationCmd.Flags().String("expiration-date", "", "The date by which this activation request should expire, in timestamp format, such as \"2024-07-07T00:00:00\".")
	ssm_createActivationCmd.Flags().String("iam-role", "", "The name of the Identity and Access Management (IAM) role that you want to assign to the managed node.")
	ssm_createActivationCmd.Flags().String("registration-limit", "", "Specify the maximum number of managed nodes you want to register.")
	ssm_createActivationCmd.Flags().String("registration-metadata", "", "Reserved for internal use.")
	ssm_createActivationCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
	ssm_createActivationCmd.MarkFlagRequired("iam-role")
	ssmCmd.AddCommand(ssm_createActivationCmd)
}
