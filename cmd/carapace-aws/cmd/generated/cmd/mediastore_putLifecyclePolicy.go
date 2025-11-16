package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_putLifecyclePolicyCmd = &cobra.Command{
	Use:   "put-lifecycle-policy",
	Short: "Writes an object lifecycle policy to a container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_putLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_putLifecyclePolicyCmd).Standalone()

		mediastore_putLifecyclePolicyCmd.Flags().String("container-name", "", "The name of the container that you want to assign the object lifecycle policy to.")
		mediastore_putLifecyclePolicyCmd.Flags().String("lifecycle-policy", "", "The object lifecycle policy to apply to the container.")
		mediastore_putLifecyclePolicyCmd.MarkFlagRequired("container-name")
		mediastore_putLifecyclePolicyCmd.MarkFlagRequired("lifecycle-policy")
	})
	mediastoreCmd.AddCommand(mediastore_putLifecyclePolicyCmd)
}
