package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_createDataViewCmd = &cobra.Command{
	Use:   "create-data-view",
	Short: "Creates a Dataview for a Dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_createDataViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_createDataViewCmd).Standalone()

		finspaceData_createDataViewCmd.Flags().String("as-of-timestamp", "", "Beginning time to use for the Dataview.")
		finspaceData_createDataViewCmd.Flags().Bool("auto-update", false, "Flag to indicate Dataview should be updated automatically.")
		finspaceData_createDataViewCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_createDataViewCmd.Flags().String("dataset-id", "", "The unique Dataset identifier that is used to create a Dataview.")
		finspaceData_createDataViewCmd.Flags().String("destination-type-params", "", "Options that define the destination type for the Dataview.")
		finspaceData_createDataViewCmd.Flags().Bool("no-auto-update", false, "Flag to indicate Dataview should be updated automatically.")
		finspaceData_createDataViewCmd.Flags().String("partition-columns", "", "Ordered set of column names used to partition data.")
		finspaceData_createDataViewCmd.Flags().String("sort-columns", "", "Columns to be used for sorting the data.")
		finspaceData_createDataViewCmd.MarkFlagRequired("dataset-id")
		finspaceData_createDataViewCmd.MarkFlagRequired("destination-type-params")
		finspaceData_createDataViewCmd.Flag("no-auto-update").Hidden = true
	})
	finspaceDataCmd.AddCommand(finspaceData_createDataViewCmd)
}
