package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "Lambda\n\n**Overview**\n\nLambda is a compute service that lets you run code without provisioning or managing servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambdaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambdaCmd).Standalone()

	})
	rootCmd.AddCommand(lambdaCmd)
}
