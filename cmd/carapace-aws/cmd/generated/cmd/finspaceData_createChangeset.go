package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_createChangesetCmd = &cobra.Command{
	Use:   "create-changeset",
	Short: "Creates a new Changeset in a FinSpace Dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_createChangesetCmd).Standalone()

	finspaceData_createChangesetCmd.Flags().String("change-type", "", "The option to indicate how a Changeset will be applied to a Dataset.")
	finspaceData_createChangesetCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspaceData_createChangesetCmd.Flags().String("dataset-id", "", "The unique identifier for the FinSpace Dataset where the Changeset will be created.")
	finspaceData_createChangesetCmd.Flags().String("format-params", "", "Options that define the structure of the source file(s) including the format type (`formatType`), header row (`withHeader`), data separation character (`separator`) and the type of compression (`compression`).")
	finspaceData_createChangesetCmd.Flags().String("source-params", "", "Options that define the location of the data being ingested (`s3SourcePath`) and the source of the changeset (`sourceType`).")
	finspaceData_createChangesetCmd.MarkFlagRequired("change-type")
	finspaceData_createChangesetCmd.MarkFlagRequired("dataset-id")
	finspaceData_createChangesetCmd.MarkFlagRequired("format-params")
	finspaceData_createChangesetCmd.MarkFlagRequired("source-params")
	finspaceDataCmd.AddCommand(finspaceData_createChangesetCmd)
}
