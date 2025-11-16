package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_getDataViewCmd = &cobra.Command{
	Use:   "get-data-view",
	Short: "Gets information about a Dataview.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_getDataViewCmd).Standalone()

	finspaceData_getDataViewCmd.Flags().String("data-view-id", "", "The unique identifier for the Dataview.")
	finspaceData_getDataViewCmd.Flags().String("dataset-id", "", "The unique identifier for the Dataset used in the Dataview.")
	finspaceData_getDataViewCmd.MarkFlagRequired("data-view-id")
	finspaceData_getDataViewCmd.MarkFlagRequired("dataset-id")
	finspaceDataCmd.AddCommand(finspaceData_getDataViewCmd)
}
