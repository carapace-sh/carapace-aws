package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_getRandomPasswordCmd = &cobra.Command{
	Use:   "get-random-password",
	Short: "Generates a random password.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_getRandomPasswordCmd).Standalone()

	secretsmanager_getRandomPasswordCmd.Flags().String("exclude-characters", "", "A string of the characters that you don't want in the password.")
	secretsmanager_getRandomPasswordCmd.Flags().String("exclude-lowercase", "", "Specifies whether to exclude lowercase letters from the password.")
	secretsmanager_getRandomPasswordCmd.Flags().String("exclude-numbers", "", "Specifies whether to exclude numbers from the password.")
	secretsmanager_getRandomPasswordCmd.Flags().String("exclude-punctuation", "", "Specifies whether to exclude the following punctuation characters from the password: ``!")
	secretsmanager_getRandomPasswordCmd.Flags().String("exclude-uppercase", "", "Specifies whether to exclude uppercase letters from the password.")
	secretsmanager_getRandomPasswordCmd.Flags().String("include-space", "", "Specifies whether to include the space character.")
	secretsmanager_getRandomPasswordCmd.Flags().String("password-length", "", "The length of the password.")
	secretsmanager_getRandomPasswordCmd.Flags().String("require-each-included-type", "", "Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation.")
	secretsmanagerCmd.AddCommand(secretsmanager_getRandomPasswordCmd)
}
