package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getDataCatalogCmd = &cobra.Command{
	Use:   "get-data-catalog",
	Short: "Returns the specified data catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getDataCatalogCmd).Standalone()

	athena_getDataCatalogCmd.Flags().String("name", "", "The name of the data catalog to return.")
	athena_getDataCatalogCmd.Flags().String("work-group", "", "The name of the workgroup.")
	athena_getDataCatalogCmd.MarkFlagRequired("name")
	athenaCmd.AddCommand(athena_getDataCatalogCmd)
}
