package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createCollaborationCmd = &cobra.Command{
	Use:   "create-collaboration",
	Short: "Creates a new collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createCollaborationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_createCollaborationCmd).Standalone()

		cleanrooms_createCollaborationCmd.Flags().String("allowed-result-regions", "", "The Amazon Web Services Regions where collaboration query results can be stored.")
		cleanrooms_createCollaborationCmd.Flags().String("analytics-engine", "", "The analytics engine.")
		cleanrooms_createCollaborationCmd.Flags().String("auto-approved-change-request-types", "", "The types of change requests that are automatically approved for this collaboration.")
		cleanrooms_createCollaborationCmd.Flags().String("creator-display-name", "", "The display name of the collaboration creator.")
		cleanrooms_createCollaborationCmd.Flags().String("creator-member-abilities", "", "The abilities granted to the collaboration creator.")
		cleanrooms_createCollaborationCmd.Flags().String("creator-mlmember-abilities", "", "The ML abilities granted to the collaboration creator.")
		cleanrooms_createCollaborationCmd.Flags().String("creator-payment-configuration", "", "The collaboration creator's payment responsibilities set by the collaboration creator.")
		cleanrooms_createCollaborationCmd.Flags().String("data-encryption-metadata", "", "The settings for client-side encryption with Cryptographic Computing for Clean Rooms.")
		cleanrooms_createCollaborationCmd.Flags().String("description", "", "A description of the collaboration provided by the collaboration owner.")
		cleanrooms_createCollaborationCmd.Flags().String("job-log-status", "", "Specifies whether job logs are enabled for this collaboration.")
		cleanrooms_createCollaborationCmd.Flags().String("members", "", "A list of initial members, not including the creator.")
		cleanrooms_createCollaborationCmd.Flags().String("name", "", "The display name for a collaboration.")
		cleanrooms_createCollaborationCmd.Flags().String("query-log-status", "", "An indicator as to whether query logging has been enabled or disabled for the collaboration.")
		cleanrooms_createCollaborationCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
		cleanrooms_createCollaborationCmd.MarkFlagRequired("creator-display-name")
		cleanrooms_createCollaborationCmd.MarkFlagRequired("creator-member-abilities")
		cleanrooms_createCollaborationCmd.MarkFlagRequired("description")
		cleanrooms_createCollaborationCmd.MarkFlagRequired("members")
		cleanrooms_createCollaborationCmd.MarkFlagRequired("name")
		cleanrooms_createCollaborationCmd.MarkFlagRequired("query-log-status")
	})
	cleanroomsCmd.AddCommand(cleanrooms_createCollaborationCmd)
}
