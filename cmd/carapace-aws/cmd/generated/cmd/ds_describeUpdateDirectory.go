package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeUpdateDirectoryCmd = &cobra.Command{
	Use:   "describe-update-directory",
	Short: "Describes the updates of a directory for a particular update type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeUpdateDirectoryCmd).Standalone()

	ds_describeUpdateDirectoryCmd.Flags().String("directory-id", "", "The unique identifier of the directory.")
	ds_describeUpdateDirectoryCmd.Flags().String("next-token", "", "The `DescribeUpdateDirectoryResult`.")
	ds_describeUpdateDirectoryCmd.Flags().String("region-name", "", "The name of the Region.")
	ds_describeUpdateDirectoryCmd.Flags().String("update-type", "", "The type of updates you want to describe for the directory.")
	ds_describeUpdateDirectoryCmd.MarkFlagRequired("directory-id")
	ds_describeUpdateDirectoryCmd.MarkFlagRequired("update-type")
	dsCmd.AddCommand(ds_describeUpdateDirectoryCmd)
}
