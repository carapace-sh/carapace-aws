package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listManagedWorkgroupsCmd = &cobra.Command{
	Use:   "list-managed-workgroups",
	Short: "Returns information about a list of specified managed workgroups in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listManagedWorkgroupsCmd).Standalone()

	redshiftServerless_listManagedWorkgroupsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	redshiftServerless_listManagedWorkgroupsCmd.Flags().String("next-token", "", "If your initial ListManagedWorkgroups operation returns a nextToken, you can include the returned nextToken in following ListManagedWorkgroups operations, which returns results in the next page.")
	redshiftServerless_listManagedWorkgroupsCmd.Flags().String("source-arn", "", "The Amazon Resource Name (ARN) for the managed workgroup in the AWS Glue Data Catalog.")
	redshiftServerlessCmd.AddCommand(redshiftServerless_listManagedWorkgroupsCmd)
}
