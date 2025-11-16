package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_createPackageCmd = &cobra.Command{
	Use:   "create-package",
	Short: "Creates a package and storage location in an Amazon S3 access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_createPackageCmd).Standalone()

	panorama_createPackageCmd.Flags().String("package-name", "", "A name for the package.")
	panorama_createPackageCmd.Flags().String("tags", "", "Tags for the package.")
	panorama_createPackageCmd.MarkFlagRequired("package-name")
	panoramaCmd.AddCommand(panorama_createPackageCmd)
}
