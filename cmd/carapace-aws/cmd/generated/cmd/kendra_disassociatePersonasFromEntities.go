package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_disassociatePersonasFromEntitiesCmd = &cobra.Command{
	Use:   "disassociate-personas-from-entities",
	Short: "Removes the specific permissions of users or groups in your IAM Identity Center identity source with access to your Amazon Kendra experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_disassociatePersonasFromEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_disassociatePersonasFromEntitiesCmd).Standalone()

		kendra_disassociatePersonasFromEntitiesCmd.Flags().String("entity-ids", "", "The identifiers of users or groups in your IAM Identity Center identity source.")
		kendra_disassociatePersonasFromEntitiesCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience.")
		kendra_disassociatePersonasFromEntitiesCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
		kendra_disassociatePersonasFromEntitiesCmd.MarkFlagRequired("entity-ids")
		kendra_disassociatePersonasFromEntitiesCmd.MarkFlagRequired("id")
		kendra_disassociatePersonasFromEntitiesCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_disassociatePersonasFromEntitiesCmd)
}
