package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_registerPackageVersionCmd = &cobra.Command{
	Use:   "register-package-version",
	Short: "Registers a package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_registerPackageVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_registerPackageVersionCmd).Standalone()

		panorama_registerPackageVersionCmd.Flags().String("mark-latest", "", "Whether to mark the new version as the latest version.")
		panorama_registerPackageVersionCmd.Flags().String("owner-account", "", "An owner account.")
		panorama_registerPackageVersionCmd.Flags().String("package-id", "", "A package ID.")
		panorama_registerPackageVersionCmd.Flags().String("package-version", "", "A package version.")
		panorama_registerPackageVersionCmd.Flags().String("patch-version", "", "A patch version.")
		panorama_registerPackageVersionCmd.MarkFlagRequired("package-id")
		panorama_registerPackageVersionCmd.MarkFlagRequired("package-version")
		panorama_registerPackageVersionCmd.MarkFlagRequired("patch-version")
	})
	panoramaCmd.AddCommand(panorama_registerPackageVersionCmd)
}
