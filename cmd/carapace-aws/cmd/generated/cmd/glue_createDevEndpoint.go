package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createDevEndpointCmd = &cobra.Command{
	Use:   "create-dev-endpoint",
	Short: "Creates a new development endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createDevEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createDevEndpointCmd).Standalone()

		glue_createDevEndpointCmd.Flags().String("arguments", "", "A map of arguments used to configure the `DevEndpoint`.")
		glue_createDevEndpointCmd.Flags().String("endpoint-name", "", "The name to be assigned to the new `DevEndpoint`.")
		glue_createDevEndpointCmd.Flags().String("extra-jars-s3-path", "", "The path to one or more Java `.jar` files in an S3 bucket that should be loaded in your `DevEndpoint`.")
		glue_createDevEndpointCmd.Flags().String("extra-python-libs-s3-path", "", "The paths to one or more Python libraries in an Amazon S3 bucket that should be loaded in your `DevEndpoint`.")
		glue_createDevEndpointCmd.Flags().String("glue-version", "", "Glue version determines the versions of Apache Spark and Python that Glue supports.")
		glue_createDevEndpointCmd.Flags().String("number-of-nodes", "", "The number of Glue Data Processing Units (DPUs) to allocate to this `DevEndpoint`.")
		glue_createDevEndpointCmd.Flags().String("number-of-workers", "", "The number of workers of a defined `workerType` that are allocated to the development endpoint.")
		glue_createDevEndpointCmd.Flags().String("public-key", "", "The public key to be used by this `DevEndpoint` for authentication.")
		glue_createDevEndpointCmd.Flags().String("public-keys", "", "A list of public keys to be used by the development endpoints for authentication.")
		glue_createDevEndpointCmd.Flags().String("role-arn", "", "The IAM role for the `DevEndpoint`.")
		glue_createDevEndpointCmd.Flags().String("security-configuration", "", "The name of the `SecurityConfiguration` structure to be used with this `DevEndpoint`.")
		glue_createDevEndpointCmd.Flags().String("security-group-ids", "", "Security group IDs for the security groups to be used by the new `DevEndpoint`.")
		glue_createDevEndpointCmd.Flags().String("subnet-id", "", "The subnet ID for the new `DevEndpoint` to use.")
		glue_createDevEndpointCmd.Flags().String("tags", "", "The tags to use with this DevEndpoint.")
		glue_createDevEndpointCmd.Flags().String("worker-type", "", "The type of predefined worker that is allocated to the development endpoint.")
		glue_createDevEndpointCmd.MarkFlagRequired("endpoint-name")
		glue_createDevEndpointCmd.MarkFlagRequired("role-arn")
	})
	glueCmd.AddCommand(glue_createDevEndpointCmd)
}
