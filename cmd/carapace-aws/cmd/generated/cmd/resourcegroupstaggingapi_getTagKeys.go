package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapi_getTagKeysCmd = &cobra.Command{
	Use:   "get-tag-keys",
	Short: "Returns all tag keys currently in use in the specified Amazon Web Services Region for the calling account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapi_getTagKeysCmd).Standalone()

	resourcegroupstaggingapi_getTagKeysCmd.Flags().String("pagination-token", "", "Specifies a `PaginationToken` response value from a previous request to indicate that you want the next page of results.")
	resourcegroupstaggingapiCmd.AddCommand(resourcegroupstaggingapi_getTagKeysCmd)
}
