package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeDirectoriesCmd = &cobra.Command{
	Use:   "describe-directories",
	Short: "Obtains information about the directories that belong to this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeDirectoriesCmd).Standalone()

	ds_describeDirectoriesCmd.Flags().String("directory-ids", "", "A list of identifiers of the directories for which to obtain the information.")
	ds_describeDirectoriesCmd.Flags().String("limit", "", "The maximum number of items to return.")
	ds_describeDirectoriesCmd.Flags().String("next-token", "", "The `DescribeDirectoriesResult.NextToken` value from a previous call to [DescribeDirectories]().")
	dsCmd.AddCommand(ds_describeDirectoriesCmd)
}
