package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_createProfileJobCmd = &cobra.Command{
	Use:   "create-profile-job",
	Short: "Creates a new job to analyze a dataset and create its data profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_createProfileJobCmd).Standalone()

	databrew_createProfileJobCmd.Flags().String("configuration", "", "Configuration for profile jobs.")
	databrew_createProfileJobCmd.Flags().String("dataset-name", "", "The name of the dataset that this job is to act upon.")
	databrew_createProfileJobCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) of an encryption key that is used to protect the job.")
	databrew_createProfileJobCmd.Flags().String("encryption-mode", "", "The encryption mode for the job, which can be one of the following:")
	databrew_createProfileJobCmd.Flags().String("job-sample", "", "Sample configuration for profile jobs only.")
	databrew_createProfileJobCmd.Flags().String("log-subscription", "", "Enables or disables Amazon CloudWatch logging for the job.")
	databrew_createProfileJobCmd.Flags().String("max-capacity", "", "The maximum number of nodes that DataBrew can use when the job processes data.")
	databrew_createProfileJobCmd.Flags().String("max-retries", "", "The maximum number of times to retry the job after a job run fails.")
	databrew_createProfileJobCmd.Flags().String("name", "", "The name of the job to be created.")
	databrew_createProfileJobCmd.Flags().String("output-location", "", "")
	databrew_createProfileJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role to be assumed when DataBrew runs the job.")
	databrew_createProfileJobCmd.Flags().String("tags", "", "Metadata tags to apply to this job.")
	databrew_createProfileJobCmd.Flags().String("timeout", "", "The job's timeout in minutes.")
	databrew_createProfileJobCmd.Flags().String("validation-configurations", "", "List of validation configurations that are applied to the profile job.")
	databrew_createProfileJobCmd.MarkFlagRequired("dataset-name")
	databrew_createProfileJobCmd.MarkFlagRequired("name")
	databrew_createProfileJobCmd.MarkFlagRequired("output-location")
	databrew_createProfileJobCmd.MarkFlagRequired("role-arn")
	databrewCmd.AddCommand(databrew_createProfileJobCmd)
}
