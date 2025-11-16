package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateGlossaryTermCmd = &cobra.Command{
	Use:   "update-glossary-term",
	Short: "Updates a business glossary term in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateGlossaryTermCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateGlossaryTermCmd).Standalone()

		datazone_updateGlossaryTermCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which a business glossary term is to be updated.")
		datazone_updateGlossaryTermCmd.Flags().String("glossary-identifier", "", "The identifier of the business glossary in which a term is to be updated.")
		datazone_updateGlossaryTermCmd.Flags().String("identifier", "", "The identifier of the business glossary term that is to be updated.")
		datazone_updateGlossaryTermCmd.Flags().String("long-description", "", "The long description to be updated as part of the `UpdateGlossaryTerm` action.")
		datazone_updateGlossaryTermCmd.Flags().String("name", "", "The name to be updated as part of the `UpdateGlossaryTerm` action.")
		datazone_updateGlossaryTermCmd.Flags().String("short-description", "", "The short description to be updated as part of the `UpdateGlossaryTerm` action.")
		datazone_updateGlossaryTermCmd.Flags().String("status", "", "The status to be updated as part of the `UpdateGlossaryTerm` action.")
		datazone_updateGlossaryTermCmd.Flags().String("term-relations", "", "The term relations to be updated as part of the `UpdateGlossaryTerm` action.")
		datazone_updateGlossaryTermCmd.MarkFlagRequired("domain-identifier")
		datazone_updateGlossaryTermCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_updateGlossaryTermCmd)
}
