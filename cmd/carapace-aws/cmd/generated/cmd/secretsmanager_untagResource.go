package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes specific tags from a secret.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_untagResourceCmd).Standalone()

		secretsmanager_untagResourceCmd.Flags().String("secret-id", "", "The ARN or name of the secret.")
		secretsmanager_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag key names to remove from the secret.")
		secretsmanager_untagResourceCmd.MarkFlagRequired("secret-id")
		secretsmanager_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_untagResourceCmd)
}
