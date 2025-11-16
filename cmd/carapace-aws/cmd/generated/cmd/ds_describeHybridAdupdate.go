package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeHybridAdupdateCmd = &cobra.Command{
	Use:   "describe-hybrid-adupdate",
	Short: "Retrieves information about update activities for a hybrid directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeHybridAdupdateCmd).Standalone()

	ds_describeHybridAdupdateCmd.Flags().String("directory-id", "", "The identifier of the hybrid directory for which to retrieve update information.")
	ds_describeHybridAdupdateCmd.Flags().String("next-token", "", "The pagination token from a previous request to [DescribeHybridADUpdate]().")
	ds_describeHybridAdupdateCmd.Flags().String("update-type", "", "The type of update activities to retrieve.")
	ds_describeHybridAdupdateCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_describeHybridAdupdateCmd)
}
