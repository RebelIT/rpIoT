package sys

type Sysinfo struct{
	Host 	HostInfo `json:"host"`
	Cpu 	CpuInfo `json:"cpu"`
	Mem 	MemInfo `json:"net"`
	Disk	DiskInfo `json:"disk"	`
}

type HostInfo struct{
	Uptime 		string `json:"uptime"`
	KernelVer	string `json:"kernel_version"`
	Platform 	string `json:"platform"`
	PlatformVer	string `json:"platform_ver"`
	Boottime	string `json:"boottime"`
	Hostname	string `json:"hostname"`
	HostId		string `json:"host_id"`
}

type CpuInfo struct{
	CpuStats 	[]CpuStat `json:"cpu_stats"`
}

type CpuStat struct{
	CpuNum			int `json:"cpu_num"`
	Mhz				float64 `json:"mhz"`
	ModelName		string `json:"model_name"`
	Cores			int32 `json:"cores"`
	UtilPercent 	float64 `json:"util_percent"`
}

type MemInfo struct{
	Total 	string `json:"total"`
	Free 	string `json:"free"`
}

type DiskInfo struct{
	Total	uint64 `json:"total"`
	Used 	uint64 `json:"used"`
	Free	uint64 `json:"free"`
}