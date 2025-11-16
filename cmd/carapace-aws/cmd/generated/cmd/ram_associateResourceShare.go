package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_associateResourceShareCmd = &cobra.Command{
	Use:   "associate-resource-share",
	Short: "Adds the specified list of principals and list of resources to a resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_associateResourceShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_associateResourceShareCmd).Standalone()

		ram_associateResourceShareCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_associateResourceShareCmd.Flags().String("principals", "", "Specifies a list of principals to whom you want to the resource share.")
		ram_associateResourceShareCmd.Flags().String("resource-arns", "", "Specifies a list of [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resources that you want to share.")
		ram_associateResourceShareCmd.Flags().String("resource-share-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share that you want to add principals or resources to.")
		ram_associateResourceShareCmd.Flags().String("sources", "", "Specifies from which source accounts the service principal has access to the resources in this resource share.")
		ram_associateResourceShareCmd.MarkFlagRequired("resource-share-arn")
	})
	ramCmd.AddCommand(ram_associateResourceShareCmd)
}
