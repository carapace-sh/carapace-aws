package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepoCmd = &cobra.Command{
	Use:   "serverlessrepo",
	Short: "The AWS Serverless Application Repository makes it easy for developers and enterprises to quickly find and deploy serverless applications in the AWS Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepoCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serverlessrepoCmd).Standalone()

	})
	rootCmd.AddCommand(serverlessrepoCmd)
}
