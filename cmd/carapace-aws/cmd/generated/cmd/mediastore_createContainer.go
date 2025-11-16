package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_createContainerCmd = &cobra.Command{
	Use:   "create-container",
	Short: "Creates a storage container to hold objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_createContainerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_createContainerCmd).Standalone()

		mediastore_createContainerCmd.Flags().String("container-name", "", "The name for the container.")
		mediastore_createContainerCmd.Flags().String("tags", "", "An array of key:value pairs that you define.")
		mediastore_createContainerCmd.MarkFlagRequired("container-name")
	})
	mediastoreCmd.AddCommand(mediastore_createContainerCmd)
}
