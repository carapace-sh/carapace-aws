package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_getLifecyclePolicyCmd = &cobra.Command{
	Use:   "get-lifecycle-policy",
	Short: "Retrieves the object lifecycle policy that is assigned to a container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_getLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_getLifecyclePolicyCmd).Standalone()

		mediastore_getLifecyclePolicyCmd.Flags().String("container-name", "", "The name of the container that the object lifecycle policy is assigned to.")
		mediastore_getLifecyclePolicyCmd.MarkFlagRequired("container-name")
	})
	mediastoreCmd.AddCommand(mediastore_getLifecyclePolicyCmd)
}
