package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_listWhatsAppTemplateLibraryCmd = &cobra.Command{
	Use:   "list-whats-app-template-library",
	Short: "Lists templates available in Meta's template library for WhatsApp messaging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_listWhatsAppTemplateLibraryCmd).Standalone()

	socialmessaging_listWhatsAppTemplateLibraryCmd.Flags().String("filters", "", "Map of filters to apply (searchKey, topic, usecase, industry, language).")
	socialmessaging_listWhatsAppTemplateLibraryCmd.Flags().String("id", "", "The ID of the WhatsApp Business Account to list library templates for.")
	socialmessaging_listWhatsAppTemplateLibraryCmd.Flags().String("max-results", "", "The maximum number of results to return per page (1-100).")
	socialmessaging_listWhatsAppTemplateLibraryCmd.Flags().String("next-token", "", "The token for the next page of results.")
	socialmessaging_listWhatsAppTemplateLibraryCmd.MarkFlagRequired("id")
	socialmessagingCmd.AddCommand(socialmessaging_listWhatsAppTemplateLibraryCmd)
}
