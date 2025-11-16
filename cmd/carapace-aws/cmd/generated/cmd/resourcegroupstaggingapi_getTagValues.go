package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapi_getTagValuesCmd = &cobra.Command{
	Use:   "get-tag-values",
	Short: "Returns all tag values for the specified key that are used in the specified Amazon Web Services Region for the calling account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapi_getTagValuesCmd).Standalone()

	resourcegroupstaggingapi_getTagValuesCmd.Flags().String("key", "", "Specifies the tag key for which you want to list all existing values that are currently used in the specified Amazon Web Services Region for the calling account.")
	resourcegroupstaggingapi_getTagValuesCmd.Flags().String("pagination-token", "", "Specifies a `PaginationToken` response value from a previous request to indicate that you want the next page of results.")
	resourcegroupstaggingapi_getTagValuesCmd.MarkFlagRequired("key")
	resourcegroupstaggingapiCmd.AddCommand(resourcegroupstaggingapi_getTagValuesCmd)
}
