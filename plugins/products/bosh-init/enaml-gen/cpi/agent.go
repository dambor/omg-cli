package cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Agent struct {

	/*Nats - Descr: Address of the nats server Default: <nil>
*/
	Nats *AgentNats `yaml:"nats,omitempty"`

	/*Mbus - Descr: Agent mbus Default: <nil>
*/
	Mbus interface{} `yaml:"mbus,omitempty"`

	/*Blobstore - Descr: AWS access_key_id for agent used by s3 blobstore plugin Default: <nil>
*/
	Blobstore *AgentBlobstore `yaml:"blobstore,omitempty"`

}