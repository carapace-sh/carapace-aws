package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilderCmd = &cobra.Command{
	Use:   "imagebuilder",
	Short: "EC2 Image Builder is a fully managed Amazon Web Services service that makes it easier to automate the creation, management, and deployment of customized, secure, and up-to-date \"golden\" server images that are pre-installed and pre-configured with software and settings to meet specific IT standards.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilderCmd).Standalone()

	})
	rootCmd.AddCommand(imagebuilderCmd)
}
