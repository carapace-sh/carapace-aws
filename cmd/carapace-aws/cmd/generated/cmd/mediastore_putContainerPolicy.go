package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_putContainerPolicyCmd = &cobra.Command{
	Use:   "put-container-policy",
	Short: "Creates an access policy for the specified container to restrict the users and clients that can access it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_putContainerPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_putContainerPolicyCmd).Standalone()

		mediastore_putContainerPolicyCmd.Flags().String("container-name", "", "The name of the container.")
		mediastore_putContainerPolicyCmd.Flags().String("policy", "", "The contents of the policy, which includes the following:")
		mediastore_putContainerPolicyCmd.MarkFlagRequired("container-name")
		mediastore_putContainerPolicyCmd.MarkFlagRequired("policy")
	})
	mediastoreCmd.AddCommand(mediastore_putContainerPolicyCmd)
}
