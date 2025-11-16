package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_associatePackageCmd = &cobra.Command{
	Use:   "associate-package",
	Short: "Associates a package with an Amazon ES domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_associatePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_associatePackageCmd).Standalone()

		es_associatePackageCmd.Flags().String("domain-name", "", "Name of the domain that you want to associate the package with.")
		es_associatePackageCmd.Flags().String("package-id", "", "Internal ID of the package that you want to associate with a domain.")
		es_associatePackageCmd.MarkFlagRequired("domain-name")
		es_associatePackageCmd.MarkFlagRequired("package-id")
	})
	esCmd.AddCommand(es_associatePackageCmd)
}
