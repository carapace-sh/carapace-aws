package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Attaches tags to a secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_tagResourceCmd).Standalone()

	secretsmanager_tagResourceCmd.Flags().String("secret-id", "", "The identifier for the secret to attach tags to.")
	secretsmanager_tagResourceCmd.Flags().String("tags", "", "The tags to attach to the secret as a JSON text string argument.")
	secretsmanager_tagResourceCmd.MarkFlagRequired("secret-id")
	secretsmanager_tagResourceCmd.MarkFlagRequired("tags")
	secretsmanagerCmd.AddCommand(secretsmanager_tagResourceCmd)
}
