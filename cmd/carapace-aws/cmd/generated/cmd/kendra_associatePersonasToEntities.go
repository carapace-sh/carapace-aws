package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_associatePersonasToEntitiesCmd = &cobra.Command{
	Use:   "associate-personas-to-entities",
	Short: "Defines the specific permissions of users or groups in your IAM Identity Center identity source with access to your Amazon Kendra experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_associatePersonasToEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_associatePersonasToEntitiesCmd).Standalone()

		kendra_associatePersonasToEntitiesCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience.")
		kendra_associatePersonasToEntitiesCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
		kendra_associatePersonasToEntitiesCmd.Flags().String("personas", "", "The personas that define the specific permissions of users or groups in your IAM Identity Center identity source.")
		kendra_associatePersonasToEntitiesCmd.MarkFlagRequired("id")
		kendra_associatePersonasToEntitiesCmd.MarkFlagRequired("index-id")
		kendra_associatePersonasToEntitiesCmd.MarkFlagRequired("personas")
	})
	kendraCmd.AddCommand(kendra_associatePersonasToEntitiesCmd)
}
