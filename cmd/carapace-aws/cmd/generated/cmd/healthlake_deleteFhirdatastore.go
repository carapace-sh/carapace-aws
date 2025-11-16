package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_deleteFhirdatastoreCmd = &cobra.Command{
	Use:   "delete-fhirdatastore",
	Short: "Delete a FHIR-enabled data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_deleteFhirdatastoreCmd).Standalone()

	healthlake_deleteFhirdatastoreCmd.Flags().String("datastore-id", "", "The AWS-generated identifier for the data store to be deleted.")
	healthlake_deleteFhirdatastoreCmd.MarkFlagRequired("datastore-id")
	healthlakeCmd.AddCommand(healthlake_deleteFhirdatastoreCmd)
}
