package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_listManagedResourcesCmd = &cobra.Command{
	Use:   "list-managed-resources",
	Short: "Lists all the resources in your Amazon Web Services account in this Amazon Web Services Region that are managed for zonal shifts in Amazon Application Recovery Controller, and information about them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_listManagedResourcesCmd).Standalone()

	arcZonalShift_listManagedResourcesCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
	arcZonalShift_listManagedResourcesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	arcZonalShiftCmd.AddCommand(arcZonalShift_listManagedResourcesCmd)
}
