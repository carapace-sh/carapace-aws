package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describePackageVersionCmd = &cobra.Command{
	Use:   "describe-package-version",
	Short: "Returns information about a package version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describePackageVersionCmd).Standalone()

	panorama_describePackageVersionCmd.Flags().String("owner-account", "", "The version's owner account.")
	panorama_describePackageVersionCmd.Flags().String("package-id", "", "The version's ID.")
	panorama_describePackageVersionCmd.Flags().String("package-version", "", "The version's version.")
	panorama_describePackageVersionCmd.Flags().String("patch-version", "", "The version's patch version.")
	panorama_describePackageVersionCmd.MarkFlagRequired("package-id")
	panorama_describePackageVersionCmd.MarkFlagRequired("package-version")
	panoramaCmd.AddCommand(panorama_describePackageVersionCmd)
}
