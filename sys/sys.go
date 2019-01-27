package sys

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

func GetHostStats()(info HostInfo){
	d := HostInfo{}

	h, err := host.Info()
	if err != nil{
		return d
	}

	d.Uptime = strconv.FormatUint(h.Uptime, 10)
	d.Platform = h.Platform
	d.KernelVer = h.KernelVersion
	d.Platform = h.Platform
	d.PlatformVer = h.PlatformVersion
	d.Boottime = strconv.FormatUint(h.BootTime, 10)
	d.Hostname = h.Hostname
	d.HostId = h.HostID

	return d
}

func GetCpuStats()(info CpuInfo){
	d := CpuInfo{}
	dd := []CpuStat{}

	c, err := cpu.Info()
	if err != nil{
		return d
	}

	for i, _ := range c{
		s := CpuStat{}

		s.CpuNum = i
		s.UtilPercent = c[i].Mhz
		s.ModelName = c[i].ModelName
		s.Cores = c[i].Cores
		s.Mhz = s.Mhz

		dd = append(dd, s)
	}

	d.CpuStats = dd

	return d
}

func GetMemStats()(info MemInfo){
	d := MemInfo{}

	m, err := mem.VirtualMemory()
	if err != nil{
		return d
	}

	d.Total = strconv.FormatUint(m.Total, 10)
	d.Free = strconv.FormatUint(m.Free, 10)

	return d
}

func GetDiskStats()(info DiskInfo){
	d := DiskInfo{}

	m, err := disk.Usage("/")
	if err != nil{
		return d
	}

	d.Total = m.Total
	d.Used = m.Used
	d.Free = m.Free

	return d
}