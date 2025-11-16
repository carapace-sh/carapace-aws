package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_updateChangesetCmd = &cobra.Command{
	Use:   "update-changeset",
	Short: "Updates a FinSpace Changeset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_updateChangesetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_updateChangesetCmd).Standalone()

		finspaceData_updateChangesetCmd.Flags().String("changeset-id", "", "The unique identifier for the Changeset to update.")
		finspaceData_updateChangesetCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_updateChangesetCmd.Flags().String("dataset-id", "", "The unique identifier for the FinSpace Dataset in which the Changeset is created.")
		finspaceData_updateChangesetCmd.Flags().String("format-params", "", "Options that define the structure of the source file(s) including the format type (`formatType`), header row (`withHeader`), data separation character (`separator`) and the type of compression (`compression`).")
		finspaceData_updateChangesetCmd.Flags().String("source-params", "", "Options that define the location of the data being ingested (`s3SourcePath`) and the source of the changeset (`sourceType`).")
		finspaceData_updateChangesetCmd.MarkFlagRequired("changeset-id")
		finspaceData_updateChangesetCmd.MarkFlagRequired("dataset-id")
		finspaceData_updateChangesetCmd.MarkFlagRequired("format-params")
		finspaceData_updateChangesetCmd.MarkFlagRequired("source-params")
	})
	finspaceDataCmd.AddCommand(finspaceData_updateChangesetCmd)
}
