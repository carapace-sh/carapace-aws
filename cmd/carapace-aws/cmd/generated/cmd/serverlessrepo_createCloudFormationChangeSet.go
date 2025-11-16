package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_createCloudFormationChangeSetCmd = &cobra.Command{
	Use:   "create-cloud-formation-change-set",
	Short: "Creates an AWS CloudFormation change set for the given application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_createCloudFormationChangeSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serverlessrepo_createCloudFormationChangeSetCmd).Standalone()

		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("capabilities", "", "A list of values that you must specify before you can deploy certain applications.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("change-set-name", "", "This property corresponds to the parameter of the same name for the *AWS CloudFormation [CreateChangeSet](https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)* API.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("client-token", "", "This property corresponds to the parameter of the same name for the *AWS CloudFormation [CreateChangeSet](https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)* API.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("description", "", "This property corresponds to the parameter of the same name for the *AWS CloudFormation [CreateChangeSet](https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)* API.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("notification-arns", "", "This property corresponds to the parameter of the same name for the *AWS CloudFormation [CreateChangeSet](https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)* API.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("parameter-overrides", "", "A list of parameter values for the parameters of the application.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("resource-types", "", "This property corresponds to the parameter of the same name for the *AWS CloudFormation [CreateChangeSet](https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)* API.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("rollback-configuration", "", "This property corresponds to the parameter of the same name for the *AWS CloudFormation [CreateChangeSet](https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)* API.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("semantic-version", "", "The semantic version of the application:")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("stack-name", "", "This property corresponds to the parameter of the same name for the *AWS CloudFormation [CreateChangeSet](https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)* API.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("tags", "", "This property corresponds to the parameter of the same name for the *AWS CloudFormation [CreateChangeSet](https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateChangeSet)* API.")
		serverlessrepo_createCloudFormationChangeSetCmd.Flags().String("template-id", "", "The UUID returned by CreateCloudFormationTemplate.")
		serverlessrepo_createCloudFormationChangeSetCmd.MarkFlagRequired("application-id")
		serverlessrepo_createCloudFormationChangeSetCmd.MarkFlagRequired("stack-name")
	})
	serverlessrepoCmd.AddCommand(serverlessrepo_createCloudFormationChangeSetCmd)
}
