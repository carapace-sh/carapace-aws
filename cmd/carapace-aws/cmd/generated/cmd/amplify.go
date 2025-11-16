package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyCmd = &cobra.Command{
	Use:   "amplify",
	Short: "Amplify enables developers to develop and deploy cloud-powered mobile and web apps.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyCmd).Standalone()

	rootCmd.AddCommand(amplifyCmd)
}
