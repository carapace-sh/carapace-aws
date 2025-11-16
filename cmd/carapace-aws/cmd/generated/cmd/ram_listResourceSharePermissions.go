package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listResourceSharePermissionsCmd = &cobra.Command{
	Use:   "list-resource-share-permissions",
	Short: "Lists the RAM permissions that are associated with a resource share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listResourceSharePermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_listResourceSharePermissionsCmd).Standalone()

		ram_listResourceSharePermissionsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_listResourceSharePermissionsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_listResourceSharePermissionsCmd.Flags().String("resource-share-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share for which you want to retrieve the associated permissions.")
		ram_listResourceSharePermissionsCmd.MarkFlagRequired("resource-share-arn")
	})
	ramCmd.AddCommand(ram_listResourceSharePermissionsCmd)
}
