package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_getCorsPolicyCmd = &cobra.Command{
	Use:   "get-cors-policy",
	Short: "Returns the cross-origin resource sharing (CORS) configuration information that is set for the container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_getCorsPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_getCorsPolicyCmd).Standalone()

		mediastore_getCorsPolicyCmd.Flags().String("container-name", "", "The name of the container that the policy is assigned to.")
		mediastore_getCorsPolicyCmd.MarkFlagRequired("container-name")
	})
	mediastoreCmd.AddCommand(mediastore_getCorsPolicyCmd)
}
