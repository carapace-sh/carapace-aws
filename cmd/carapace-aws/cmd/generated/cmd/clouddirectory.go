package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectoryCmd = &cobra.Command{
	Use:   "clouddirectory",
	Short: "Amazon Cloud Directory\n\nAmazon Cloud Directory is a component of the AWS Directory Service that simplifies the development and management of cloud-scale web, mobile, and IoT applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectoryCmd).Standalone()

	})
	rootCmd.AddCommand(clouddirectoryCmd)
}
