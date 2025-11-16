package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_deletePackageCmd = &cobra.Command{
	Use:   "delete-package",
	Short: "Deletes a package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_deletePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_deletePackageCmd).Standalone()

		panorama_deletePackageCmd.Flags().Bool("force-delete", false, "Delete the package even if it has artifacts stored in its access point.")
		panorama_deletePackageCmd.Flags().Bool("no-force-delete", false, "Delete the package even if it has artifacts stored in its access point.")
		panorama_deletePackageCmd.Flags().String("package-id", "", "The package's ID.")
		panorama_deletePackageCmd.Flag("no-force-delete").Hidden = true
		panorama_deletePackageCmd.MarkFlagRequired("package-id")
	})
	panoramaCmd.AddCommand(panorama_deletePackageCmd)
}
