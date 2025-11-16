package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_updateProfileJobCmd = &cobra.Command{
	Use:   "update-profile-job",
	Short: "Modifies the definition of an existing profile job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_updateProfileJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_updateProfileJobCmd).Standalone()

		databrew_updateProfileJobCmd.Flags().String("configuration", "", "Configuration for profile jobs.")
		databrew_updateProfileJobCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) of an encryption key that is used to protect the job.")
		databrew_updateProfileJobCmd.Flags().String("encryption-mode", "", "The encryption mode for the job, which can be one of the following:")
		databrew_updateProfileJobCmd.Flags().String("job-sample", "", "Sample configuration for Profile Jobs only.")
		databrew_updateProfileJobCmd.Flags().String("log-subscription", "", "Enables or disables Amazon CloudWatch logging for the job.")
		databrew_updateProfileJobCmd.Flags().String("max-capacity", "", "The maximum number of compute nodes that DataBrew can use when the job processes data.")
		databrew_updateProfileJobCmd.Flags().String("max-retries", "", "The maximum number of times to retry the job after a job run fails.")
		databrew_updateProfileJobCmd.Flags().String("name", "", "The name of the job to be updated.")
		databrew_updateProfileJobCmd.Flags().String("output-location", "", "")
		databrew_updateProfileJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role to be assumed when DataBrew runs the job.")
		databrew_updateProfileJobCmd.Flags().String("timeout", "", "The job's timeout in minutes.")
		databrew_updateProfileJobCmd.Flags().String("validation-configurations", "", "List of validation configurations that are applied to the profile job.")
		databrew_updateProfileJobCmd.MarkFlagRequired("name")
		databrew_updateProfileJobCmd.MarkFlagRequired("output-location")
		databrew_updateProfileJobCmd.MarkFlagRequired("role-arn")
	})
	databrewCmd.AddCommand(databrew_updateProfileJobCmd)
}
