package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describePackageCmd = &cobra.Command{
	Use:   "describe-package",
	Short: "Returns information about a package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_describePackageCmd).Standalone()

		panorama_describePackageCmd.Flags().String("package-id", "", "The package's ID.")
		panorama_describePackageCmd.MarkFlagRequired("package-id")
	})
	panoramaCmd.AddCommand(panorama_describePackageCmd)
}
