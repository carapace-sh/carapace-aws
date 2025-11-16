package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_addDraftAppVersionResourceMappingsCmd = &cobra.Command{
	Use:   "add-draft-app-version-resource-mappings",
	Short: "Adds the source of resource-maps to the draft version of an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_addDraftAppVersionResourceMappingsCmd).Standalone()

	resiliencehub_addDraftAppVersionResourceMappingsCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_addDraftAppVersionResourceMappingsCmd.Flags().String("resource-mappings", "", "Mappings used to map logical resources from the template to physical resources.")
	resiliencehub_addDraftAppVersionResourceMappingsCmd.MarkFlagRequired("app-arn")
	resiliencehub_addDraftAppVersionResourceMappingsCmd.MarkFlagRequired("resource-mappings")
	resiliencehubCmd.AddCommand(resiliencehub_addDraftAppVersionResourceMappingsCmd)
}
