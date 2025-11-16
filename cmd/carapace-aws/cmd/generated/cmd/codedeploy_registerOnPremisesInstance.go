package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_registerOnPremisesInstanceCmd = &cobra.Command{
	Use:   "register-on-premises-instance",
	Short: "Registers an on-premises instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_registerOnPremisesInstanceCmd).Standalone()

	codedeploy_registerOnPremisesInstanceCmd.Flags().String("iam-session-arn", "", "The ARN of the IAM session to associate with the on-premises instance.")
	codedeploy_registerOnPremisesInstanceCmd.Flags().String("iam-user-arn", "", "The ARN of the user to associate with the on-premises instance.")
	codedeploy_registerOnPremisesInstanceCmd.Flags().String("instance-name", "", "The name of the on-premises instance to register.")
	codedeploy_registerOnPremisesInstanceCmd.MarkFlagRequired("instance-name")
	codedeployCmd.AddCommand(codedeploy_registerOnPremisesInstanceCmd)
}
