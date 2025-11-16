package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_getManagedResourceCmd = &cobra.Command{
	Use:   "get-managed-resource",
	Short: "Get information about a resource that's been registered for zonal shifts with Amazon Application Recovery Controller in this Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_getManagedResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcZonalShift_getManagedResourceCmd).Standalone()

		arcZonalShift_getManagedResourceCmd.Flags().String("resource-identifier", "", "The identifier for the resource that Amazon Web Services shifts traffic for.")
		arcZonalShift_getManagedResourceCmd.MarkFlagRequired("resource-identifier")
	})
	arcZonalShiftCmd.AddCommand(arcZonalShift_getManagedResourceCmd)
}
