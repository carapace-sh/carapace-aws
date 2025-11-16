package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_deregisterPackageVersionCmd = &cobra.Command{
	Use:   "deregister-package-version",
	Short: "Deregisters a package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_deregisterPackageVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_deregisterPackageVersionCmd).Standalone()

		panorama_deregisterPackageVersionCmd.Flags().String("owner-account", "", "An owner account.")
		panorama_deregisterPackageVersionCmd.Flags().String("package-id", "", "A package ID.")
		panorama_deregisterPackageVersionCmd.Flags().String("package-version", "", "A package version.")
		panorama_deregisterPackageVersionCmd.Flags().String("patch-version", "", "A patch version.")
		panorama_deregisterPackageVersionCmd.Flags().String("updated-latest-patch-version", "", "If the version was marked latest, the new version to maker as latest.")
		panorama_deregisterPackageVersionCmd.MarkFlagRequired("package-id")
		panorama_deregisterPackageVersionCmd.MarkFlagRequired("package-version")
		panorama_deregisterPackageVersionCmd.MarkFlagRequired("patch-version")
	})
	panoramaCmd.AddCommand(panorama_deregisterPackageVersionCmd)
}
