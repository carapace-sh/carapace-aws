package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createContentAssociationCmd = &cobra.Command{
	Use:   "create-content-association",
	Short: "Creates an association between a content resource in a knowledge base and [step-by-step guides](https://docs.aws.amazon.com/connect/latest/adminguide/step-by-step-guided-experiences.html). Step-by-step guides offer instructions to agents for resolving common customer issues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createContentAssociationCmd).Standalone()

	qconnect_createContentAssociationCmd.Flags().String("association", "", "The identifier of the associated resource.")
	qconnect_createContentAssociationCmd.Flags().String("association-type", "", "The type of association.")
	qconnect_createContentAssociationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_createContentAssociationCmd.Flags().String("content-id", "", "The identifier of the content.")
	qconnect_createContentAssociationCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_createContentAssociationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	qconnect_createContentAssociationCmd.MarkFlagRequired("association")
	qconnect_createContentAssociationCmd.MarkFlagRequired("association-type")
	qconnect_createContentAssociationCmd.MarkFlagRequired("content-id")
	qconnect_createContentAssociationCmd.MarkFlagRequired("knowledge-base-id")
	qconnectCmd.AddCommand(qconnect_createContentAssociationCmd)
}
