package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_disassociateGovernedTermsCmd = &cobra.Command{
	Use:   "disassociate-governed-terms",
	Short: "Disassociates restricted terms from an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_disassociateGovernedTermsCmd).Standalone()

	datazone_disassociateGovernedTermsCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to disassociate restricted terms from an asset.")
	datazone_disassociateGovernedTermsCmd.Flags().String("entity-identifier", "", "The ID of an asset from which you want to disassociate restricted terms.")
	datazone_disassociateGovernedTermsCmd.Flags().String("entity-type", "", "The type of the asset from which you want to disassociate restricted terms.")
	datazone_disassociateGovernedTermsCmd.Flags().String("governed-glossary-terms", "", "The restricted glossary terms that you want to disassociate from an asset.")
	datazone_disassociateGovernedTermsCmd.MarkFlagRequired("domain-identifier")
	datazone_disassociateGovernedTermsCmd.MarkFlagRequired("entity-identifier")
	datazone_disassociateGovernedTermsCmd.MarkFlagRequired("entity-type")
	datazone_disassociateGovernedTermsCmd.MarkFlagRequired("governed-glossary-terms")
	datazoneCmd.AddCommand(datazone_disassociateGovernedTermsCmd)
}
