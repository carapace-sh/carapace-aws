package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listProfileSharesCmd = &cobra.Command{
	Use:   "list-profile-shares",
	Short: "List profile shares.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listProfileSharesCmd).Standalone()

	wellarchitected_listProfileSharesCmd.Flags().String("max-results", "", "The maximum number of results to return for this request.")
	wellarchitected_listProfileSharesCmd.Flags().String("next-token", "", "")
	wellarchitected_listProfileSharesCmd.Flags().String("profile-arn", "", "The profile ARN.")
	wellarchitected_listProfileSharesCmd.Flags().String("shared-with-prefix", "", "The Amazon Web Services account ID, organization ID, or organizational unit (OU) ID with which the profile is shared.")
	wellarchitected_listProfileSharesCmd.Flags().String("status", "", "")
	wellarchitected_listProfileSharesCmd.MarkFlagRequired("profile-arn")
	wellarchitectedCmd.AddCommand(wellarchitected_listProfileSharesCmd)
}
