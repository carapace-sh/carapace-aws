package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_deleteDataProviderCmd = &cobra.Command{
	Use:   "delete-data-provider",
	Short: "Deletes the specified data provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_deleteDataProviderCmd).Standalone()

	dms_deleteDataProviderCmd.Flags().String("data-provider-identifier", "", "The identifier of the data provider to delete.")
	dms_deleteDataProviderCmd.MarkFlagRequired("data-provider-identifier")
	dmsCmd.AddCommand(dms_deleteDataProviderCmd)
}
