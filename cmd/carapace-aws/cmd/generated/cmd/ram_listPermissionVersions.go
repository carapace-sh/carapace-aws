package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listPermissionVersionsCmd = &cobra.Command{
	Use:   "list-permission-versions",
	Short: "Lists the available versions of the specified RAM permission.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listPermissionVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_listPermissionVersionsCmd).Standalone()

		ram_listPermissionVersionsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
		ram_listPermissionVersionsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ram_listPermissionVersionsCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the RAM permission whose versions you want to list.")
		ram_listPermissionVersionsCmd.MarkFlagRequired("permission-arn")
	})
	ramCmd.AddCommand(ram_listPermissionVersionsCmd)
}
