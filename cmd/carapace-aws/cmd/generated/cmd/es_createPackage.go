package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_createPackageCmd = &cobra.Command{
	Use:   "create-package",
	Short: "Create a package for use with Amazon ES domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_createPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_createPackageCmd).Standalone()

		es_createPackageCmd.Flags().String("package-description", "", "Description of the package.")
		es_createPackageCmd.Flags().String("package-name", "", "Unique identifier for the package.")
		es_createPackageCmd.Flags().String("package-source", "", "The customer S3 location `PackageSource` for importing the package.")
		es_createPackageCmd.Flags().String("package-type", "", "Type of package.")
		es_createPackageCmd.MarkFlagRequired("package-name")
		es_createPackageCmd.MarkFlagRequired("package-source")
		es_createPackageCmd.MarkFlagRequired("package-type")
	})
	esCmd.AddCommand(es_createPackageCmd)
}
