package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_getResourceSharesCmd = &cobra.Command{
	Use:   "get-resource-shares",
	Short: "Retrieves details about the resource shares that you own or that are shared with you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_getResourceSharesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_getResourceSharesCmd).Standalone()

		ram_getResourceSharesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_getResourceSharesCmd.Flags().String("name", "", "Specifies the name of an individual resource share that you want to retrieve details about.")
		ram_getResourceSharesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_getResourceSharesCmd.Flags().String("permission-arn", "", "Specifies that you want to retrieve details of only those resource shares that use the managed permission with this [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).")
		ram_getResourceSharesCmd.Flags().String("permission-version", "", "Specifies that you want to retrieve details for only those resource shares that use the specified version of the managed permission.")
		ram_getResourceSharesCmd.Flags().String("resource-owner", "", "Specifies that you want to retrieve details of only those resource shares that match the following:")
		ram_getResourceSharesCmd.Flags().String("resource-share-arns", "", "Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of individual resource shares that you want information about.")
		ram_getResourceSharesCmd.Flags().String("resource-share-status", "", "Specifies that you want to retrieve details of only those resource shares that have this status.")
		ram_getResourceSharesCmd.Flags().String("tag-filters", "", "Specifies that you want to retrieve details of only those resource shares that match the specified tag keys and values.")
		ram_getResourceSharesCmd.MarkFlagRequired("resource-owner")
	})
	ramCmd.AddCommand(ram_getResourceSharesCmd)
}
