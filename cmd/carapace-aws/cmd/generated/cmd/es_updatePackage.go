package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_updatePackageCmd = &cobra.Command{
	Use:   "update-package",
	Short: "Updates a package for use with Amazon ES domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_updatePackageCmd).Standalone()

	es_updatePackageCmd.Flags().String("commit-message", "", "An info message for the new version which will be shown as part of `GetPackageVersionHistoryResponse`.")
	es_updatePackageCmd.Flags().String("package-description", "", "New description of the package.")
	es_updatePackageCmd.Flags().String("package-id", "", "Unique identifier for the package.")
	es_updatePackageCmd.Flags().String("package-source", "", "")
	es_updatePackageCmd.MarkFlagRequired("package-id")
	es_updatePackageCmd.MarkFlagRequired("package-source")
	esCmd.AddCommand(es_updatePackageCmd)
}
