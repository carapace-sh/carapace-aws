package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_deleteContainerPolicyCmd = &cobra.Command{
	Use:   "delete-container-policy",
	Short: "Deletes the access policy that is associated with the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_deleteContainerPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_deleteContainerPolicyCmd).Standalone()

		mediastore_deleteContainerPolicyCmd.Flags().String("container-name", "", "The name of the container that holds the policy.")
		mediastore_deleteContainerPolicyCmd.MarkFlagRequired("container-name")
	})
	mediastoreCmd.AddCommand(mediastore_deleteContainerPolicyCmd)
}
