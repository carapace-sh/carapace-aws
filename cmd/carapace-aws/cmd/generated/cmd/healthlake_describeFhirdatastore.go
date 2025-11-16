package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_describeFhirdatastoreCmd = &cobra.Command{
	Use:   "describe-fhirdatastore",
	Short: "Get properties for a FHIR-enabled data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_describeFhirdatastoreCmd).Standalone()

	healthlake_describeFhirdatastoreCmd.Flags().String("datastore-id", "", "The data store identifier.")
	healthlake_describeFhirdatastoreCmd.MarkFlagRequired("datastore-id")
	healthlakeCmd.AddCommand(healthlake_describeFhirdatastoreCmd)
}
