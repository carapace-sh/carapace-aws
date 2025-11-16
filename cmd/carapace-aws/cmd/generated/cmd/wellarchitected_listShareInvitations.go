package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listShareInvitationsCmd = &cobra.Command{
	Use:   "list-share-invitations",
	Short: "List the share invitations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listShareInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_listShareInvitationsCmd).Standalone()

		wellarchitected_listShareInvitationsCmd.Flags().String("lens-name-prefix", "", "An optional string added to the beginning of each lens name returned in the results.")
		wellarchitected_listShareInvitationsCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
		wellarchitected_listShareInvitationsCmd.Flags().String("next-token", "", "")
		wellarchitected_listShareInvitationsCmd.Flags().String("profile-name-prefix", "", "An optional string added to the beginning of each profile name returned in the results.")
		wellarchitected_listShareInvitationsCmd.Flags().String("share-resource-type", "", "The type of share invitations to be returned.")
		wellarchitected_listShareInvitationsCmd.Flags().String("template-name-prefix", "", "An optional string added to the beginning of each review template name returned in the results.")
		wellarchitected_listShareInvitationsCmd.Flags().String("workload-name-prefix", "", "")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_listShareInvitationsCmd)
}
