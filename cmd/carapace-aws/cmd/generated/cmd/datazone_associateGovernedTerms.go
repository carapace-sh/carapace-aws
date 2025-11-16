package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_associateGovernedTermsCmd = &cobra.Command{
	Use:   "associate-governed-terms",
	Short: "Associates governed terms with an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_associateGovernedTermsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_associateGovernedTermsCmd).Standalone()

		datazone_associateGovernedTermsCmd.Flags().String("domain-identifier", "", "The ID of the domain where governed terms are to be associated with an asset.")
		datazone_associateGovernedTermsCmd.Flags().String("entity-identifier", "", "The ID of the asset with which you want to associate a governed term.")
		datazone_associateGovernedTermsCmd.Flags().String("entity-type", "", "The type of the asset with which you want to associate a governed term.")
		datazone_associateGovernedTermsCmd.Flags().String("governed-glossary-terms", "", "The glossary terms in a restricted glossary.")
		datazone_associateGovernedTermsCmd.MarkFlagRequired("domain-identifier")
		datazone_associateGovernedTermsCmd.MarkFlagRequired("entity-identifier")
		datazone_associateGovernedTermsCmd.MarkFlagRequired("entity-type")
		datazone_associateGovernedTermsCmd.MarkFlagRequired("governed-glossary-terms")
	})
	datazoneCmd.AddCommand(datazone_associateGovernedTermsCmd)
}
