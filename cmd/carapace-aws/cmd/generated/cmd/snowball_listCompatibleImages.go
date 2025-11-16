package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_listCompatibleImagesCmd = &cobra.Command{
	Use:   "list-compatible-images",
	Short: "This action returns a list of the different Amazon EC2-compatible Amazon Machine Images (AMIs) that are owned by your Amazon Web Services accountthat would be supported for use on a Snow device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_listCompatibleImagesCmd).Standalone()

	snowball_listCompatibleImagesCmd.Flags().String("max-results", "", "The maximum number of results for the list of compatible images.")
	snowball_listCompatibleImagesCmd.Flags().String("next-token", "", "HTTP requests are stateless.")
	snowballCmd.AddCommand(snowball_listCompatibleImagesCmd)
}
