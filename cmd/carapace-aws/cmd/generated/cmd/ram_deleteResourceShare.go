package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_deleteResourceShareCmd = &cobra.Command{
	Use:   "delete-resource-share",
	Short: "Deletes the specified resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_deleteResourceShareCmd).Standalone()

	ram_deleteResourceShareCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ram_deleteResourceShareCmd.Flags().String("resource-share-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share to delete.")
	ram_deleteResourceShareCmd.MarkFlagRequired("resource-share-arn")
	ramCmd.AddCommand(ram_deleteResourceShareCmd)
}
