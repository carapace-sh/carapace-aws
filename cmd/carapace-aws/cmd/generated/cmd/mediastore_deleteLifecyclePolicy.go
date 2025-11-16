package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_deleteLifecyclePolicyCmd = &cobra.Command{
	Use:   "delete-lifecycle-policy",
	Short: "Removes an object lifecycle policy from a container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_deleteLifecyclePolicyCmd).Standalone()

	mediastore_deleteLifecyclePolicyCmd.Flags().String("container-name", "", "The name of the container that holds the object lifecycle policy.")
	mediastore_deleteLifecyclePolicyCmd.MarkFlagRequired("container-name")
	mediastoreCmd.AddCommand(mediastore_deleteLifecyclePolicyCmd)
}
