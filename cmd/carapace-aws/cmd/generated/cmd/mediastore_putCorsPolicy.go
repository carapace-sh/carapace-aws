package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_putCorsPolicyCmd = &cobra.Command{
	Use:   "put-cors-policy",
	Short: "Sets the cross-origin resource sharing (CORS) configuration on a container so that the container can service cross-origin requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_putCorsPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_putCorsPolicyCmd).Standalone()

		mediastore_putCorsPolicyCmd.Flags().String("container-name", "", "The name of the container that you want to assign the CORS policy to.")
		mediastore_putCorsPolicyCmd.Flags().String("cors-policy", "", "The CORS policy to apply to the container.")
		mediastore_putCorsPolicyCmd.MarkFlagRequired("container-name")
		mediastore_putCorsPolicyCmd.MarkFlagRequired("cors-policy")
	})
	mediastoreCmd.AddCommand(mediastore_putCorsPolicyCmd)
}
