package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	customizations["ec2.modify-instance-attribute"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--enable-api-termination") // TODO should be this but seems awscli is missing the rename here
		cmd.AddFlag(command.Flag{
			Longhand:    "--no-disable-api-termination",
			Description: "Enable or disable termination protection for the instance.",
		})
		return nil
	}
	customizations["ec2.get-password-data"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "priv-launch-key",
			Description: "The file that contains the private key used to launch the instance",
			Value:       true,
		})
		return nil
	}

	customizations["ec2.run-instances"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--min-count=!")
		delete(cmd.Flags, "--max-count=!")

		cmd.AddFlag(command.Flag{
			Longhand:    "associate-public-ip-address",
			Description: "If specified a public IP address will be assigned to the new instance in a VPC",
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "no-associate-public-ip-address",
			Description: "If specified a public IP address will be assigned to the new instance in a VPC",
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "secondary-private-ip-address-count",
			Description: "The number of secondary IP addresses to assign to the network interface or instance",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "count",
			Description: "Number of instances to launch",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "secondary-private-ip-addresses",
			Description: "A secondary private IP address for the network interface or instance",
			Value:       true,
		})
		return nil
	}

	customizations["ec2.bundle-instances"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "prefix",
			Description: "The prefix for the image component names being stored in Amazon S3",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "owner-sak",
			Description: "The AWS secret access key for the owner of the Amazon S3 bucket ",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "owner-akid",
			Description: "The access key ID of the owner of the Amazon S3 bucket",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "bucket",
			Description: "The bucket in which to store the AMI",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "policy",
			Description: "An Amazon S3 upload policy that gives Amazon EC2 permission to upload items",
			Value:       true,
		})
		return nil
	}

	for _, operation := range []string{
		"ec2.authorize-security-group-ingress",
		"ec2.authorize-security-group-egress",
		"ec2.revoke-security-group-egress",
		"ec2.revoke-security-group-ingress",
	} {
		customizations[operation] = func(cmd *command.Command) error {
			cmd.AddFlag(command.Flag{
				Longhand:    "port",
				Description: "The range of ports to allow or ICMP code",
				Value:       true,
			})
			cmd.AddFlag(command.Flag{
				Longhand:    "cidr",
				Description: "The IPv4 address range, in CIDR format",
				Value:       true,
			})
			cmd.AddFlag(command.Flag{
				Longhand:    "protocol",
				Description: "The IP protoco",
				Value:       true,
			})
			cmd.AddFlag(command.Flag{
				Longhand:    "source-group",
				Description: "The name or ID of the source security group",
				Value:       true,
			})
			cmd.AddFlag(command.Flag{
				Longhand:    "group-owner",
				Description: "The AWS account ID that owns the source security group",
				Value:       true,
			})
			return nil
		}
	}
}
