package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listResourcesCmd = &cobra.Command{
	Use:   "list-resources",
	Short: "Lists the resources that you added to a resource share or the resources that are shared with you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_listResourcesCmd).Standalone()

		ram_listResourcesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_listResourcesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_listResourcesCmd.Flags().String("principal", "", "Specifies that you want to list only the resource shares that are associated with the specified principal.")
		ram_listResourcesCmd.Flags().String("resource-arns", "", "Specifies that you want to list only the resource shares that include resources with the specified [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).")
		ram_listResourcesCmd.Flags().String("resource-owner", "", "Specifies that you want to list only the resource shares that match the following:")
		ram_listResourcesCmd.Flags().String("resource-region-scope", "", "Specifies that you want the results to include only resources that have the specified scope.")
		ram_listResourcesCmd.Flags().String("resource-share-arns", "", "Specifies that you want to list only resources in the resource shares identified by the specified [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).")
		ram_listResourcesCmd.Flags().String("resource-type", "", "Specifies that you want to list only the resource shares that include resources of the specified resource type.")
		ram_listResourcesCmd.MarkFlagRequired("resource-owner")
	})
	ramCmd.AddCommand(ram_listResourcesCmd)
}
