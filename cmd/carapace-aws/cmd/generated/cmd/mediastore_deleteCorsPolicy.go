package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_deleteCorsPolicyCmd = &cobra.Command{
	Use:   "delete-cors-policy",
	Short: "Deletes the cross-origin resource sharing (CORS) configuration information that is set for the container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_deleteCorsPolicyCmd).Standalone()

	mediastore_deleteCorsPolicyCmd.Flags().String("container-name", "", "The name of the container to remove the policy from.")
	mediastore_deleteCorsPolicyCmd.MarkFlagRequired("container-name")
	mediastoreCmd.AddCommand(mediastore_deleteCorsPolicyCmd)
}
