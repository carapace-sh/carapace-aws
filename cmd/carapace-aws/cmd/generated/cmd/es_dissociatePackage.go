package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_dissociatePackageCmd = &cobra.Command{
	Use:   "dissociate-package",
	Short: "Dissociates a package from the Amazon ES domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_dissociatePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_dissociatePackageCmd).Standalone()

		es_dissociatePackageCmd.Flags().String("domain-name", "", "Name of the domain that you want to associate the package with.")
		es_dissociatePackageCmd.Flags().String("package-id", "", "Internal ID of the package that you want to associate with a domain.")
		es_dissociatePackageCmd.MarkFlagRequired("domain-name")
		es_dissociatePackageCmd.MarkFlagRequired("package-id")
	})
	esCmd.AddCommand(es_dissociatePackageCmd)
}
