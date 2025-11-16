package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_listEntitledApplicationsCmd = &cobra.Command{
	Use:   "list-entitled-applications",
	Short: "Retrieves a list of entitled applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_listEntitledApplicationsCmd).Standalone()

	appstream_listEntitledApplicationsCmd.Flags().String("entitlement-name", "", "The name of the entitlement.")
	appstream_listEntitledApplicationsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
	appstream_listEntitledApplicationsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	appstream_listEntitledApplicationsCmd.Flags().String("stack-name", "", "The name of the stack with which the entitlement is associated.")
	appstream_listEntitledApplicationsCmd.MarkFlagRequired("entitlement-name")
	appstream_listEntitledApplicationsCmd.MarkFlagRequired("stack-name")
	appstreamCmd.AddCommand(appstream_listEntitledApplicationsCmd)
}
