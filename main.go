package mylogger

import "log"

type EloConfig struct {
	ServerMemoryLimit         string `hcl:"serverMemoryLimit" json:"serverMemoryLimit"`
	ServerRequestsLimit       string `hcl:"serverRequestsLimit" json:"serverRequestsLimit"`
  	ServerMemoryRequests      string `hcl:"serverMemoryRequests" json:"serverMemoryRequests"`
	IxLogLevel                string `hcl:"ixLogLevel" json:"ixLogLevel"`
	IxDisableDbChecks         string `hcl:"ixDisableDbChecks" json:"ixDisableDbChecks"`
	WcLogLevel                string `hcl:"wcLogLevel" json:"wcLogLevel"`
	AcLogLevel                string `hcl:"acLogLevel" json:"acLogLevel"`
	SearchAioMemoryLimit      string `hcl:"searchAioMemoryLimit" json:"searchAioMemoryLimit"`
	SearchAioMemoryRequests   string `hcl:"searchAioMemoryRequests" json:"searchAioMemoryRequests"`
	SearchAioMaxMemory        string `hcl:"searchAioMaxMemory" json:"searchAioMaxMemory"`
	SearchAioBlockstorageSize string `hcl:"searchAioBlockstorageSize" json:"searchAioBlockstorageSize"`
	FlowsWorkerMemoryLimit    string `hcl:"flowsWorkerMemoryLimit" json:"flowsWorkerMemoryLimit"`
	FlowsWorkerMemoryRequests string `hcl:"flowsWorkerMemoryRequests" json:"flowsWorkerMemoryRequests"`
	FlowsMemoryLimit          string `hcl:"flowsMemoryLimit" json:"flowsMemoryLimit"`
	FlowsMemoryRequests       string `hcl:"flowsMemoryRequests" json:"flowsMemoryRequests"`
	FlowsManagerLogLevel      string `hcl:"flowsManagerLogLevel" json:"flowsManagerLogLevel"`
	FlowsWorkerLogLevel       string `hcl:"flowsWorkerLogLevel" json:"flowsWorkerLogLevel"`
	FlowsWorkerReplicas       int		 `hcl:"flowsWorkerReplicas" json:"flowsWorkerReplicas"` 	
	FlowsRegistryLogLevel     string `hcl:"flowsRegistryLogLevel" json:"flowsRegistryLogLevel"`
	GlobalMaintenanceMode     bool   `hcl:"globalMaintenanceMode" json:"globalMaintenanceMode"`
	GlobalLogLevel            string `hcl:"globalLogLevel" json:"globalLogLevel"`
	GlobalLogAppender         string `hcl:"globalLogAppender" json:"globalLogAppender"`
	GlobalLogStdoutFormat     string `hcl:"globalLogStdoutFormat" json:"globalLogStdoutFormat"`
	NetworkPoliciesEnabled    bool   `hcl:"networkPoliciesEnabled" json:"networkPoliciesEnabled"`
}

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogWarning(message string) {
	log.Printf("WARN - %v", message)
}

func LogError(message string) {
	log.Printf("ERROR - %v", message)
}
