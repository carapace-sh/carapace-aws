package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExportsCmd = &cobra.Command{
	Use:   "bcm-data-exports",
	Short: "You can use the Data Exports API to create customized exports from multiple Amazon Web Services cost management and billing datasets, such as cost and usage data and cost optimization recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDataExportsCmd).Standalone()

	})
	rootCmd.AddCommand(bcmDataExportsCmd)
}
