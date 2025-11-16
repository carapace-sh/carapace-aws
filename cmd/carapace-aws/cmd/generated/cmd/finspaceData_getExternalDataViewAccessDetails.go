package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_getExternalDataViewAccessDetailsCmd = &cobra.Command{
	Use:   "get-external-data-view-access-details",
	Short: "Returns the credentials to access the external Dataview from an S3 location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_getExternalDataViewAccessDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_getExternalDataViewAccessDetailsCmd).Standalone()

		finspaceData_getExternalDataViewAccessDetailsCmd.Flags().String("data-view-id", "", "The unique identifier for the Dataview that you want to access.")
		finspaceData_getExternalDataViewAccessDetailsCmd.Flags().String("dataset-id", "", "The unique identifier for the Dataset.")
		finspaceData_getExternalDataViewAccessDetailsCmd.MarkFlagRequired("data-view-id")
		finspaceData_getExternalDataViewAccessDetailsCmd.MarkFlagRequired("dataset-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_getExternalDataViewAccessDetailsCmd)
}
