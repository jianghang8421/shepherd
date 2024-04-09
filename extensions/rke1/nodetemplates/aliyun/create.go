package nodetemplates

import (
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/extensions/rke1/nodetemplates"
	"github.com/rancher/shepherd/pkg/config"
)

const aliyunECSTemplateNameBase = "aliyunecsConfig"

// CreateAliyunECSNodeTemplate is a helper function that takes the rancher Client as a parameter and creates
// an Aliyun ECS node template and returns the NodeTemplate response
func CreateAliyunECSNodeTemplate(rancherClient *rancher.Client) (*nodetemplates.NodeTemplate, error) {
	var aliyunECSNodeNodeTemplateConfig nodetemplates.AliyunECSNodeTemplateConfig
	config.LoadConfig(nodetemplates.AliyunECSNodeTemplateConfigurationFileKey, &aliyunECSNodeNodeTemplateConfig)
	nodeTemplate := nodetemplates.NodeTemplate{
		EngineInstallURL:            "https://releases.rancher.com/install-docker/20.10.sh",
		Name:                        aliyunECSTemplateNameBase,
		AliyunECSNodeTemplateConfig: &aliyunECSNodeNodeTemplateConfig,
	}
	resp := &nodetemplates.NodeTemplate{}
	err := rancherClient.Management.APIBaseClient.Ops.DoCreate(management.NodeTemplateType, nodeTemplate, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
