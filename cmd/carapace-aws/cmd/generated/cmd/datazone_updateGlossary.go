package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateGlossaryCmd = &cobra.Command{
	Use:   "update-glossary",
	Short: "Updates the business glossary in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateGlossaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateGlossaryCmd).Standalone()

		datazone_updateGlossaryCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_updateGlossaryCmd.Flags().String("description", "", "The description to be updated as part of the `UpdateGlossary` action.")
		datazone_updateGlossaryCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which a business glossary is to be updated.")
		datazone_updateGlossaryCmd.Flags().String("identifier", "", "The identifier of the business glossary to be updated.")
		datazone_updateGlossaryCmd.Flags().String("name", "", "The name to be updated as part of the `UpdateGlossary` action.")
		datazone_updateGlossaryCmd.Flags().String("status", "", "The status to be updated as part of the `UpdateGlossary` action.")
		datazone_updateGlossaryCmd.MarkFlagRequired("domain-identifier")
		datazone_updateGlossaryCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_updateGlossaryCmd)
}
