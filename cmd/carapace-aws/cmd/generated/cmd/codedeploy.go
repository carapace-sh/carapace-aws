package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeployCmd = &cobra.Command{
	Use:   "codedeploy",
	Short: "CodeDeploy is a deployment service that automates application deployments to Amazon EC2 instances, on-premises instances running in your own facility, serverless Lambda functions, or applications in an Amazon ECS service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeployCmd).Standalone()

	rootCmd.AddCommand(codedeployCmd)
}
