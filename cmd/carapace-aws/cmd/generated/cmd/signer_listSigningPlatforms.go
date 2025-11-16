package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_listSigningPlatformsCmd = &cobra.Command{
	Use:   "list-signing-platforms",
	Short: "Lists all signing platforms available in AWS Signer that match the request parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_listSigningPlatformsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(signer_listSigningPlatformsCmd).Standalone()

		signer_listSigningPlatformsCmd.Flags().String("category", "", "The category type of a signing platform.")
		signer_listSigningPlatformsCmd.Flags().String("max-results", "", "The maximum number of results to be returned by this operation.")
		signer_listSigningPlatformsCmd.Flags().String("next-token", "", "Value for specifying the next set of paginated results to return.")
		signer_listSigningPlatformsCmd.Flags().String("partner", "", "Any partner entities connected to a signing platform.")
		signer_listSigningPlatformsCmd.Flags().String("target", "", "The validation template that is used by the target signing platform.")
	})
	signerCmd.AddCommand(signer_listSigningPlatformsCmd)
}
