package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_deleteContainerCmd = &cobra.Command{
	Use:   "delete-container",
	Short: "Deletes the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_deleteContainerCmd).Standalone()

	mediastore_deleteContainerCmd.Flags().String("container-name", "", "The name of the container to delete.")
	mediastore_deleteContainerCmd.MarkFlagRequired("container-name")
	mediastoreCmd.AddCommand(mediastore_deleteContainerCmd)
}
