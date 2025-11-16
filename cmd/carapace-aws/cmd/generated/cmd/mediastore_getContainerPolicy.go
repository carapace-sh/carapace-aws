package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_getContainerPolicyCmd = &cobra.Command{
	Use:   "get-container-policy",
	Short: "Retrieves the access policy for the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_getContainerPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_getContainerPolicyCmd).Standalone()

		mediastore_getContainerPolicyCmd.Flags().String("container-name", "", "The name of the container.")
		mediastore_getContainerPolicyCmd.MarkFlagRequired("container-name")
	})
	mediastoreCmd.AddCommand(mediastore_getContainerPolicyCmd)
}
