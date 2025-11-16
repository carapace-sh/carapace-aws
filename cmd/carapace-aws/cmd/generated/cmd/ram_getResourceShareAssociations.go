package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_getResourceShareAssociationsCmd = &cobra.Command{
	Use:   "get-resource-share-associations",
	Short: "Retrieves the lists of resources and principals that associated for resource shares that you own.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_getResourceShareAssociationsCmd).Standalone()

	ram_getResourceShareAssociationsCmd.Flags().String("association-status", "", "Specifies that you want to retrieve only associations that have this status.")
	ram_getResourceShareAssociationsCmd.Flags().String("association-type", "", "Specifies whether you want to retrieve the associations that involve a specified resource or principal.")
	ram_getResourceShareAssociationsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
	ram_getResourceShareAssociationsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	ram_getResourceShareAssociationsCmd.Flags().String("principal", "", "Specifies the ID of the principal whose resource shares you want to retrieve.")
	ram_getResourceShareAssociationsCmd.Flags().String("resource-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of a resource whose resource shares you want to retrieve.")
	ram_getResourceShareAssociationsCmd.Flags().String("resource-share-arns", "", "Specifies a list of [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share whose associations you want to retrieve.")
	ram_getResourceShareAssociationsCmd.MarkFlagRequired("association-type")
	ramCmd.AddCommand(ram_getResourceShareAssociationsCmd)
}
