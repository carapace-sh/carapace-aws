package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createPersistentContactAssociationCmd = &cobra.Command{
	Use:   "create-persistent-contact-association",
	Short: "Enables rehydration of chats for the lifespan of a contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createPersistentContactAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createPersistentContactAssociationCmd).Standalone()

		connect_createPersistentContactAssociationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_createPersistentContactAssociationCmd.Flags().String("initial-contact-id", "", "This is the contactId of the current contact that the `CreatePersistentContactAssociation` API is being called from.")
		connect_createPersistentContactAssociationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createPersistentContactAssociationCmd.Flags().String("rehydration-type", "", "The contactId chosen for rehydration depends on the type chosen.")
		connect_createPersistentContactAssociationCmd.Flags().String("source-contact-id", "", "The contactId from which a persistent chat session must be started.")
		connect_createPersistentContactAssociationCmd.MarkFlagRequired("initial-contact-id")
		connect_createPersistentContactAssociationCmd.MarkFlagRequired("instance-id")
		connect_createPersistentContactAssociationCmd.MarkFlagRequired("rehydration-type")
		connect_createPersistentContactAssociationCmd.MarkFlagRequired("source-contact-id")
	})
	connectCmd.AddCommand(connect_createPersistentContactAssociationCmd)
}
