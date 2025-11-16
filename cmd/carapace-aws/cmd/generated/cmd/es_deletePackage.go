package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_deletePackageCmd = &cobra.Command{
	Use:   "delete-package",
	Short: "Delete the package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_deletePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_deletePackageCmd).Standalone()

		es_deletePackageCmd.Flags().String("package-id", "", "Internal ID of the package that you want to delete.")
		es_deletePackageCmd.MarkFlagRequired("package-id")
	})
	esCmd.AddCommand(es_deletePackageCmd)
}
