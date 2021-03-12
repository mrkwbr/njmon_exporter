package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// AIX section
	aixVersion = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_aix_version",
			Help: "AIX version number",
		},
		[]string{"instance"},
	)
	aixTechLevel = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_aix_techlevel",
			Help: "AIX technology level",
		},
		[]string{"instance"},
	)
	aixServicePack = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_aix_servicepack",
			Help: "AIX service pack number",
		},
		[]string{"instance"},
	)
	// Timestamp
	clockDrift = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_clock_drift",
			Help: "Difference between remote UTC and local UTC in seconds",
		},
		[]string{"instance"},
	)
	hostUp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "up",
			Help: "Returns 0 if a known host has stopped submitting metrics",
		},
		[]string{"instance"},
	)
	// CPU Details
	cpuNumActive = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_details_active",
			Help: "Number of active CPUs",
		},
		[]string{"instance"},
	)
	cpuNumConf = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_details_configured",
			Help: "Number of configured CPUs",
		},
		[]string{"instance"},
	)
	cpuMHz = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_details_mhz",
			Help: "CPU speed in MHz",
		},
		[]string{"instance"},
	)
	// CPU Utilization
	cpuTotIdle = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_total_idle",
			Help: "Percentage of CPU cycles spent idle",
		},
		[]string{"instance"},
	)
	cpuTotKern = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_total_kern",
			Help: "Percentage of CPU cycles spent on kernel",
		},
		[]string{"instance"},
	)
	cpuTotUser = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_total_user",
			Help: "Percentage of CPU cycles spent on user",
		},
		[]string{"instance"},
	)
	cpuTotWait = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_total_wait",
			Help: "Percentage of CPU cycles spent waiting",
		},
		[]string{"instance"},
	)
	// CPU Logical
	cpuLogIdle = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_logical_idle",
			Help: "Percentage of logical CPU cycles spent idle",
		},
		[]string{"instance", "cpu"},
	)
	cpuLogSys = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_logical_sys",
			Help: "Percentage of logical CPU cycles spent on sys",
		},
		[]string{"instance", "cpu"},
	)
	cpuLogUser = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_logical_user",
			Help: "Percentage of logical CPU cycles spent on user",
		},
		[]string{"instance", "cpu"},
	)
	cpuLogWait = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_logical_wait",
			Help: "Percentage of logical CPU cycles spent waiting",
		},
		[]string{"instance", "cpu"},
	)
	// CPU Load Average
	cpuLoadAvg1 = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_load1",
			Help: "1 Minute Load Average",
		},
		[]string{"instance"},
	)
	cpuLoadAvg5 = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_load5",
			Help: "5 Minute Load Average",
		},
		[]string{"instance"},
	)
	cpuLoadAvg15 = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_cpu_load15",
			Help: "15 Minute Load Average",
		},
		[]string{"instance"},
	)
	// Filesystems
	filesystemSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_filesystem_size",
			Help: "Size of the filesystem in Bytes",
		},
		[]string{"instance", "device", "mountpoint"},
	)
	filesystemFree = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_filesystem_free",
			Help: "Available filesystem space in Bytes",
		},
		[]string{"instance", "device", "mountpoint"},
	)
	filesystemInode = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_filesystem_inode",
			Help: "filesystem Inode utilization",
		},
		[]string{"instance", "device", "mountpoint"},
	)
	diskReadIO = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_disk_read_bytes",
			Help: "disk read rate in bytes per second",
		},
		[]string{"instance", "device"},
	)
	diskWriteIO = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_disk_write_bytes",
			Help: "disk write rate in bytes per second",
		},
		[]string{"instance", "device"},
	)
	// Memory
	memOnline = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_online",
			Help: "Memory allocated to the system in Bytes",
		},
		[]string{"instance"},
	)
	memMax = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_max",
			Help: "Memory maximum that can be allocated in Bytes",
		},
		[]string{"instance"},
	)
	memPgspFree = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_pgsp_free",
			Help: "Memory pagespace free in Bytes",
		},
		[]string{"instance"},
	)
	memPgspRsvd = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_pgsp_rsvd",
			Help: "Memory pagespace reserved in Bytes",
		},
		[]string{"instance"},
	)
	memPgspTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_pgsp_total",
			Help: "Memory pagespace total in Bytes",
		},
		[]string{"instance"},
	)
	memRealFree = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_real_free",
			Help: "Memory real free in Bytes",
		},
		[]string{"instance"},
	)
	memRealInUse = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_real_inuse",
			Help: "Memory real in-use in Bytes",
		},
		[]string{"instance"},
	)
	memRealPinned = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_real_pinned",
			Help: "Memory real pinned in Bytes",
		},
		[]string{"instance"},
	)
	memRealProcess = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_real_process",
			Help: "Memory real total in Bytes",
		},
		[]string{"instance"},
	)
	memRealSystem = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_real_system",
			Help: "Memory real total in Bytes",
		},
		[]string{"instance"},
	)
	memRealTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_real_total",
			Help: "Memory real total in Bytes",
		},
		[]string{"instance"},
	)
	memRealUser = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_real_user",
			Help: "Memory real total in Bytes",
		},
		[]string{"instance"},
	)
	memPagingSpaceUsed = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_paging_space_used",
			Help: "Paging space used in Bytes",
		},
		[]string{"instance", "device"},
	)
	memPagingSpaceSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_mem_paging_space_size",
			Help: "Paging space total in Bytes",
		},
		[]string{"instance", "device"},
	)
	netPktRx = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_net_pkt_rx_total",
			Help: "Network packet receive rate",
		},
		[]string{"instance", "interface"},
	)
	netPktRxDrp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_net_pkt_rx_drop",
			Help: "Network total RX packets dropped",
		},
		[]string{"instance", "interface"},
	)
	netPktTx = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_net_pkt_tx_total",
			Help: "Network packet transmit rate",
		},
		[]string{"instance", "interface"},
	)
	netPktTxDrp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_net_pkt_tx_drop",
			Help: "Network total TX packets dropped",
		},
		[]string{"instance", "interface"},
	)
	netBpsRx = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_net_bps_rx",
			Help: "Network bytes/s receive rate",
		},
		[]string{"instance", "interface"},
	)
	netBpsTx = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_net_bps_tx",
			Help: "Network bytes/s transmit rate",
		},
		[]string{"instance", "interface"},
	)
	systemUptime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "njmon_system_uptime",
			Help: "System uptime in seconds",
		},
		[]string{"instance"},
	)
)

func init() {
	prometheus.MustRegister(aixVersion)
	prometheus.MustRegister(aixTechLevel)
	prometheus.MustRegister(aixServicePack)
	prometheus.MustRegister(clockDrift)
	prometheus.MustRegister(cpuNumActive)
	prometheus.MustRegister(cpuNumConf)
	prometheus.MustRegister(cpuMHz)
	prometheus.MustRegister(cpuTotIdle)
	prometheus.MustRegister(cpuTotKern)
	prometheus.MustRegister(cpuTotUser)
	prometheus.MustRegister(cpuTotWait)
	prometheus.MustRegister(cpuLogIdle)
	prometheus.MustRegister(cpuLogSys)
	prometheus.MustRegister(cpuLogUser)
	prometheus.MustRegister(cpuLogWait)
	prometheus.MustRegister(cpuLoadAvg1)
	prometheus.MustRegister(cpuLoadAvg5)
	prometheus.MustRegister(cpuLoadAvg15)
	prometheus.MustRegister(filesystemSize)
	prometheus.MustRegister(filesystemFree)
	prometheus.MustRegister(filesystemInode)
	prometheus.MustRegister(diskReadIO)
	prometheus.MustRegister(diskWriteIO)
	prometheus.MustRegister(hostUp)
	prometheus.MustRegister(memMax)
	prometheus.MustRegister(memOnline)
	prometheus.MustRegister(memPgspFree)
	prometheus.MustRegister(memPgspRsvd)
	prometheus.MustRegister(memPgspTotal)
	prometheus.MustRegister(memRealFree)
	prometheus.MustRegister(memRealInUse)
	prometheus.MustRegister(memRealPinned)
	prometheus.MustRegister(memRealProcess)
	prometheus.MustRegister(memRealSystem)
	prometheus.MustRegister(memRealTotal)
	prometheus.MustRegister(memRealUser)
	prometheus.MustRegister(memPagingSpaceUsed)
	prometheus.MustRegister(memPagingSpaceSize)
	prometheus.MustRegister(netBpsRx)
	prometheus.MustRegister(netBpsTx)
	prometheus.MustRegister(netPktRx)
	prometheus.MustRegister(netPktRxDrp)
	prometheus.MustRegister(netPktTx)
	prometheus.MustRegister(netPktTxDrp)
	prometheus.MustRegister(systemUptime)
}
