package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_cancelLoaderJobCmd = &cobra.Command{
	Use:   "cancel-loader-job",
	Short: "Cancels a specified load job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_cancelLoaderJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_cancelLoaderJobCmd).Standalone()

		neptunedata_cancelLoaderJobCmd.Flags().String("load-id", "", "The ID of the load job to be deleted.")
		neptunedata_cancelLoaderJobCmd.MarkFlagRequired("load-id")
	})
	neptunedataCmd.AddCommand(neptunedata_cancelLoaderJobCmd)
}
