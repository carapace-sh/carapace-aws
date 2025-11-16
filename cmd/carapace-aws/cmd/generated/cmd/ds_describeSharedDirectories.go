package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeSharedDirectoriesCmd = &cobra.Command{
	Use:   "describe-shared-directories",
	Short: "Returns the shared directories in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeSharedDirectoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeSharedDirectoriesCmd).Standalone()

		ds_describeSharedDirectoriesCmd.Flags().String("limit", "", "The number of shared directories to return in the response object.")
		ds_describeSharedDirectoriesCmd.Flags().String("next-token", "", "The `DescribeSharedDirectoriesResult.NextToken` value from a previous call to [DescribeSharedDirectories]().")
		ds_describeSharedDirectoriesCmd.Flags().String("owner-directory-id", "", "Returns the identifier of the directory in the directory owner account.")
		ds_describeSharedDirectoriesCmd.Flags().String("shared-directory-ids", "", "A list of identifiers of all shared directories in your account.")
		ds_describeSharedDirectoriesCmd.MarkFlagRequired("owner-directory-id")
	})
	dsCmd.AddCommand(ds_describeSharedDirectoriesCmd)
}
