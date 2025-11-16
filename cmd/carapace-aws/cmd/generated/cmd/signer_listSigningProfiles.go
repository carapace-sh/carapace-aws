package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_listSigningProfilesCmd = &cobra.Command{
	Use:   "list-signing-profiles",
	Short: "Lists all available signing profiles in your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_listSigningProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_listSigningProfilesCmd).Standalone()

		signer_listSigningProfilesCmd.Flags().String("include-canceled", "", "Designates whether to include profiles with the status of `CANCELED`.")
		signer_listSigningProfilesCmd.Flags().String("max-results", "", "The maximum number of profiles to be returned.")
		signer_listSigningProfilesCmd.Flags().String("next-token", "", "Value for specifying the next set of paginated results to return.")
		signer_listSigningProfilesCmd.Flags().String("platform-id", "", "Filters results to return only signing jobs initiated for a specified signing platform.")
		signer_listSigningProfilesCmd.Flags().String("statuses", "", "Filters results to return only signing jobs with statuses in the specified list.")
	})
	signerCmd.AddCommand(signer_listSigningProfilesCmd)
}
