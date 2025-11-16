package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_disassociateResourceShareCmd = &cobra.Command{
	Use:   "disassociate-resource-share",
	Short: "Removes the specified principals or resources from participating in the specified resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_disassociateResourceShareCmd).Standalone()

	ram_disassociateResourceShareCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ram_disassociateResourceShareCmd.Flags().String("principals", "", "Specifies a list of one or more principals that no longer are to have access to the resources in this resource share.")
	ram_disassociateResourceShareCmd.Flags().String("resource-arns", "", "Specifies a list of [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) for one or more resources that you want to remove from the resource share.")
	ram_disassociateResourceShareCmd.Flags().String("resource-share-arn", "", "Specifies [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share that you want to remove resources or principals from.")
	ram_disassociateResourceShareCmd.Flags().String("sources", "", "Specifies from which source accounts the service principal no longer has access to the resources in this resource share.")
	ram_disassociateResourceShareCmd.MarkFlagRequired("resource-share-arn")
	ramCmd.AddCommand(ram_disassociateResourceShareCmd)
}
